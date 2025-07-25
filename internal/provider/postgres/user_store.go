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
	//go:embed queries/update_user.sql
	updateUserQuery string
	//go:embed queries/get_user_by_email.sql
	getUserByEmail string
	//go:embed queries/list_users_by_email.sql
	listUsersByEmail string
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
	row := s.db.QueryRow(ctx, getUserByEmail, email)

	var user domain.User

	if err := row.Scan(&user.ID, &user.Email, &user.FullName, &user.CreateTime); err != nil {
		return domain.User{}, err
	}

	return user, nil
}

func (s *UserStore) ListUsersByEmail(ctx context.Context, emails []string) ([]domain.User, error) {
	rows, err := s.db.Query(ctx, listUsersByEmail, emails)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var users []domain.User

	for rows.Next() {
		var user domain.User
		if err := rows.Scan(&user.ID, &user.Email, &user.FullName, &user.CreateTime); err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}

func (s *UserStore) UpdateUser(ctx context.Context, user domain.User) error {
	ct, err := s.db.Exec(ctx, updateUserQuery, user.Email, user.FullName, user.ID)
	if err != nil {
		return err
	}

	if ct.RowsAffected() == 0 {
		return ErrUserDoesNotExist
	}

	return nil
}
