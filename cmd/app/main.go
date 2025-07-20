package main

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"

	v1 "github.com/meetmorrowsolonmars/education-pet-project/internal/controller/v1"
	"github.com/meetmorrowsolonmars/education-pet-project/internal/provider/postgres"
)

func main() {
	// Create a database connection.
	db, err := openDatabaseConnection(context.Background())
	if err != nil {
		slog.Error("Open database connection", slog.String("error", err.Error()))
		os.Exit(1)
	}

	defer db.Close()

	slog.Info("Connected to database")

	userStore := postgres.NewUserStore(db)
	userController := v1.NewUserController(userStore)

	// Start HTTP server.
	mux := http.NewServeMux()

	userController.Register(mux)

	server := &http.Server{
		Addr:    os.Getenv("SERVER_ADDRESS"),
		Handler: mux,
	}

	go func() {
		slog.Info("Starting server", slog.String("address", server.Addr))

		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			slog.Error("Server error", slog.String("error", err.Error()))
		}
	}()

	// Graceful shutdown.
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	<-ctx.Done()

	slog.Info("Stopping server", slog.String("address", server.Addr))

	ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_ = server.Shutdown(ctx)

	slog.Info("Server stopped", slog.String("address", server.Addr))
}

func openDatabaseConnection(ctx context.Context) (*pgxpool.Pool, error) {
	connectionURL := os.Getenv("DATABASE_URL")

	if connectionURL == "" {
		return nil, errors.New("connection-url parameter is required")
	}

	config, err := pgxpool.ParseConfig(connectionURL)
	if err != nil {
		return nil, fmt.Errorf("error parsing connection url: %w", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	db, err := pgxpool.NewWithConfig(ctx, config)
	if err != nil {
		return nil, fmt.Errorf("error creating connection pool: %w", err)
	}

	return db, nil
}
