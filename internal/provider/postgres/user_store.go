package postgres

import (
	"context"
	_ "embed"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/meetmorrowsolonmars/education-pet-project/internal/domain"
)

var (
	//go:embed queries/create_user.sql
	createUserQuery string
)

type UserStore struct {
	db *pgxpool.Pool
}

func NewUserStore(db *pgxpool.Pool) *UserStore {
	return &UserStore{
		db: db,
	}
}

func (s *UserStore) Create(ctx context.Context, user domain.User) (domain.User, error) {
	row := s.db.QueryRow(ctx, createUserQuery, user.Email, user.FullName)

	err := row.Scan(&user.ID, &user.CreateTime)
	if err != nil {
		return domain.User{}, err
	}

	return user, nil
}

func (s *UserStore) GetByEmail(ctx context.Context, email string) (domain.User, error) {
	// TODO: Implement it.
	return domain.User{}, nil
}
