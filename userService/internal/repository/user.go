package repository

import (
	"context"
	"database/sql"
	"errors"
	"github.com/saleh-ghazimoradi/StandardMicroEcoBay/internal/domain"
	"time"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user *domain.User) error
	FindUserByEmail(ctx context.Context, email string) (*domain.User, error)
	FindUserByResetToken(ctx context.Context, token string) (*domain.User, error)
	FindUserById(ctx context.Context, id int64) (*domain.User, error)
	SaveUser(ctx context.Context, user *domain.User) error
}

type userRepository struct {
	dbWrite *sql.DB
	dbRead  *sql.DB
}

func (u *userRepository) CreateUser(ctx context.Context, user *domain.User) error {
	query := `INSERT INTO users (email, password, phone) VALUES ($1, $2, $3) RETURNING id, created_at, updated_at, version`
	args := []interface{}{user.Email, user.Password, user.Phone}
	if err := u.dbWrite.QueryRowContext(ctx, query, args...).Scan(&user.Id, &user.CreatedAt, &user.UpdatedAt, &user.Version); err != nil {
		switch {
		case err.Error() == `pq: duplicate key value violates unique constraint "users_email_key"`:
			return ErrDuplicateEmail
		default:
			return err

		}
	}
	return nil
}

func (u *userRepository) FindUserByEmail(ctx context.Context, email string) (*domain.User, error) {
	var user domain.User

	query := `SELECT id, email, password, first_name, last_name, phone, reset_token, created_at, updated_at, version FROM users WHERE email = $1`

	args := []any{email}
	if err := u.dbRead.QueryRowContext(ctx, query, args).Scan(
		&user.Id,
		&user.Email,
		&user.Password,
		&user.FirstName,
		&user.LastName,
		&user.Phone,
		&user.ResetToken,
		&user.CreatedAt,
		&user.UpdatedAt,
		&user.Version,
	); err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrRecordNotFound
		default:
			return nil, err
		}
	}

	return &user, nil
}

func (u *userRepository) FindUserByResetToken(ctx context.Context, token string) (*domain.User, error) {
	var user domain.User

	query := `SELECT id, email, password, first_name, last_name, phone, reset_token, created_at, updated_at, version FROM users WHERE reset_token = $1`

	args := []any{token}

	if err := u.dbRead.QueryRowContext(ctx, query, args).Scan(
		&user.Id,
		&user.Email,
		&user.Password,
		&user.FirstName,
		&user.LastName,
		&user.Phone,
		&user.ResetToken,
		&user.CreatedAt,
		&user.UpdatedAt,
		&user.Version,
	); err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrInvalidResetToken
		default:
			return nil, err
		}
	}

	return &user, nil
}

func (u *userRepository) FindUserById(ctx context.Context, id int64) (*domain.User, error) {
	var user domain.User

	query := `SELECT id, email, password, first_name, last_name, phone, reset_token, created_at, updated_at, version FROM users WHERE id = $1`

	args := []any{id}

	if err := u.dbRead.QueryRowContext(ctx, query, args).Scan(
		&user.Id,
		&user.Email,
		&user.Password,
		&user.FirstName,
		&user.LastName,
		&user.Phone,
		&user.ResetToken,
		&user.CreatedAt,
		&user.UpdatedAt,
		&user.Version,
	); err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrRecordNotFound
		}
		return nil, err
	}

	addrQuery := `SELECT id, address_line1, address_line2, city, post_code, country, user_id, created_at, updated_at, version FROM addresses WHERE user_id = $1`
	var addr domain.Address
	if err := u.dbRead.QueryRowContext(ctx, addrQuery, id).Scan(); err == nil {
		user.Address = addr
	} else if !errors.Is(err, sql.ErrNoRows) {
		return nil, err
	}

	return &user, nil
}

func (u *userRepository) SaveUser(ctx context.Context, user *domain.User) error {
	query := `UPDATE users SET first_name = $1, last_name = $2, email = $3, phone = $4, reset_token = $5, updated_at = $6, version = version + 1
	          WHERE id = $7 RETURNING updated_at, version`
	now := time.Now().UTC()
	args := []any{user.FirstName, user.LastName, user.Email, user.Phone, user.ResetToken, now, user.Id}

	if err := u.dbWrite.QueryRowContext(ctx, query, args...).Scan(&user.UpdatedAt, &user.Version); err != nil {
		return err
	}

	if user.Address.AddressLine1 != "" || user.Address.AddressLine2 != "" || user.Address.City != "" || user.Address.Country != "" || user.Address.PostCode != "" {
		if user.Address.Id == 0 {
			addrQuery := `INSERT INTO addresses (address_line1, address_line2, city, post_code, country, user_id)
			              VALUES ($1, $2, $3, $4, $5, $6) RETURNING id, created_at, updated_at, version`
			addrArgs := []any{user.Address.AddressLine1, user.Address.AddressLine2, user.Address.City, user.Address.PostCode, user.Address.Country, user.Id}
			if err := u.dbWrite.QueryRowContext(ctx, addrQuery, addrArgs...).Scan(&user.Address.Id, &user.Address.CreatedAt, &user.Address.UpdatedAt, &user.Address.Version); err != nil {
				return err
			}
		} else {
			addrQuery := `UPDATE addresses SET address_line1 = $1, address_line2 = $2, city = $3, post_code = $4, country = $5, updated_at = $6, version = version + 1
			             WHERE id = $7 RETURNING updated_at, version`
			addrArgs := []any{user.Address.AddressLine1, user.Address.AddressLine2, user.Address.City, user.Address.PostCode, user.Address.Country, now, user.Address.Id}
			if err := u.dbWrite.QueryRowContext(ctx, addrQuery, addrArgs...).Scan(&user.Address.UpdatedAt, &user.Address.Version); err != nil {
				return err
			}
		}
	}
	return nil
}

func NewUserRepository(dbWrite, dbRead *sql.DB) UserRepository {
	return &userRepository{
		dbWrite: dbWrite,
		dbRead:  dbRead,
	}
}
