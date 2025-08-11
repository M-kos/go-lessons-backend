package postgres

import (
	"context"
	"database/sql"
	_ "embed"
	"errors"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/meetmorrowsolonmars/education-pet-project/internal/domain"
)

var (
	//go:embed queries/create_user.sql
	createUserQuery string
	//go:embed queries/create_country.sql
	createCountryQuery string
	//go:embed queries/create_city.sql
	createCityQuery string
	//go:embed queries/create_address.sql
	createAddressQuery string
	//go:embed queries/update_user.sql
	updateUserQuery string
	//go:embed queries/get_user_by_email.sql
	getUserByEmail string
	//go:embed queries/get_country_by_name.sql
	getCountryByName string
	//go:embed queries/get_city_by_name.sql
	getCityByName string
	//go:embed queries/get_address_by_address.sql
	getAddressByAddress string
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
	tx, err := s.db.Begin(ctx)
	if err != nil {
		return domain.User{}, err
	}

	defer func() {
		if err != nil {
			_ = tx.Rollback(ctx)
			return
		}
		_ = tx.Commit(ctx)
	}()

	row := tx.QueryRow(ctx, getCountryByName, user.Country)
	var countryId int
	if err := row.Scan(&countryId); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			row = tx.QueryRow(ctx, createCountryQuery, user.Country)

			if err := row.Scan(&countryId); err != nil {
				return domain.User{}, err
			}
		} else {
			return domain.User{}, err
		}
	}

	row = tx.QueryRow(ctx, getCityByName, user.City, countryId)
	var cityId int
	if err := row.Scan(&cityId); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			row = tx.QueryRow(ctx, createCityQuery, user.City, countryId)

			if err := row.Scan(&cityId); err != nil {
				return domain.User{}, err
			}
		} else {
			return domain.User{}, err
		}
	}

	row = tx.QueryRow(ctx, getAddressByAddress, user.Street, user.Zip, cityId)
	var addressId int
	if err := row.Scan(&addressId); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			row = tx.QueryRow(ctx, createAddressQuery, user.Street, user.Zip, cityId)

			if err := row.Scan(&countryId); err != nil {
				return domain.User{}, err
			}
		} else {
			return domain.User{}, err
		}
	}

	row = tx.QueryRow(ctx, createUserQuery, user.Email, user.FullName, addressId)

	err = row.Scan(&user.ID, &user.CreateTime)
	if err != nil {
		return domain.User{}, err
	}

	return user, nil
}

func (s *UserStore) GetByEmail(ctx context.Context, email string) (domain.User, error) {
	row := s.db.QueryRow(ctx, getUserByEmail, email)

	var user domain.User

	if err := row.Scan(&user.ID, &user.Email, &user.FullName, &user.CreateTime, &user.Street, &user.Zip, &user.City, &user.Country); err != nil {
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
		if err := rows.Scan(&user.ID, &user.Email, &user.FullName, &user.CreateTime, &user.Street, &user.Zip, &user.City, &user.Country); err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}

func (s *UserStore) UpdateUser(ctx context.Context, user domain.User) error {
	tx, err := s.db.Begin(ctx)
	if err != nil {
		return err
	}

	defer func() {
		if err != nil {
			_ = tx.Rollback(ctx)
			return
		}
		_ = tx.Commit(ctx)
	}()

	var countryId int
	row := tx.QueryRow(ctx, getCountryByName, user.Country)
	if err := row.Scan(&countryId); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			row = tx.QueryRow(ctx, createCountryQuery, user.Country)

			if err := row.Scan(&countryId); err != nil {
				return err
			}
			return err
		} else {
		}
	}

	var cityId int
	row = tx.QueryRow(ctx, getCityByName, user.City, countryId)
	if err := row.Scan(&cityId); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			row = tx.QueryRow(ctx, createCityQuery, user.City, countryId)

			if err := row.Scan(&cityId); err != nil {
				return err
			}
		} else {
			return err
		}
	}

	row = tx.QueryRow(ctx, getAddressByAddress, user.Street, user.Zip, cityId)
	var addressId int
	if err := row.Scan(&addressId); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			row = tx.QueryRow(ctx, createAddressQuery, user.Street, user.Zip, cityId)

			if err := row.Scan(&countryId); err != nil {
				return err
			}
		} else {
			return err
		}
	}

	ct, err := s.db.Exec(ctx, updateUserQuery, user.Email, user.FullName, addressId, user.ID)
	if err != nil {
		return err
	}

	if ct.RowsAffected() == 0 {
		return ErrUserDoesNotExist
	}

	return nil
}
