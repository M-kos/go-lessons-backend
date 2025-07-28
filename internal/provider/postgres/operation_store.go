package postgres

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/shopspring/decimal"

	"github.com/meetmorrowsolonmars/education-pet-project/internal/domain"
)

type OperationStore struct {
	db *pgxpool.Pool
}

func NewOperationStore(db *pgxpool.Pool) *OperationStore {
	return &OperationStore{
		db: db,
	}
}

func (s *OperationStore) ListOperationsByUserID(ctx context.Context, userID int64) ([]domain.Operation, error) {
	// TODO: Implement me.
	return []domain.Operation{}, nil
}

func (s *OperationStore) GetUserBalance(ctx context.Context, userID int64) (decimal.Decimal, error) {
	// TODO: Implement me.
	return decimal.Zero, nil
}

func (s *OperationStore) CreateTransfer(
	ctx context.Context,
	sourceUserID int64,
	targetUserID int64,
	amount decimal.Decimal,
) (uuid.UUID, error) {
	// TODO: Implement me.
	return uuid.Nil, nil
}
