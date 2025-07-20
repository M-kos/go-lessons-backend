package v1

import (
	"context"
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/meetmorrowsolonmars/education-pet-project/internal/domain"
)

type UserStore interface {
	Create(ctx context.Context, user domain.User) (domain.User, error)
}

type UserController struct {
	userStore UserStore
}

func NewUserController(userStore UserStore) *UserController {
	return &UserController{
		userStore: userStore,
	}
}

type CreateUserRequest struct {
	Email    string `json:"email"`
	FullName string `json:"full_name"`
}

type CreateUserResponse struct {
	ID int64 `json:"id"`
}

func (c *UserController) CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	req, err := DecodeRequest[CreateUserRequest](r)
	if err != nil {
		slog.Error("Decode create user request", slog.String("error", err.Error()))

		err = WriteErrorResponse(w, http.StatusBadRequest, "Invalid request body: %s", err)
		if err != nil {
			slog.Error("Write error response", slog.String("error", err.Error()))
		}

		return
	}

	user, err := c.userStore.Create(r.Context(), domain.User{
		Email:    req.Email,
		FullName: req.FullName,
	})
	if err != nil {
		slog.Error("Create user", slog.String("error", err.Error()))

		err = WriteErrorResponse(w, http.StatusInternalServerError, "Create user error")
		if err != nil {
			slog.Error("Write error response", slog.String("error", err.Error()))
		}

		return
	}

	w.WriteHeader(http.StatusCreated)

	err = json.NewEncoder(w).Encode(CreateUserResponse{
		ID: user.ID,
	})
	if err != nil {
		slog.Error("Write create user response", slog.String("error", err.Error()))
	}
}

func (c *UserController) Register(mux *http.ServeMux) {
	mux.HandleFunc("POST /v1/users", c.CreateUser)
}
