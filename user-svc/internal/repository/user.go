package repository

import (
	"context"
	"database/sql"
	"github.com/saleh-ghazimoradi/StandardMicroEcoBay/user-svc/internal/domain"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user *domain.User) error
	FindUserByEmail(ctx context.Context, email string) (*domain.User, error)
	SaveUser(ctx context.Context, user *domain.User) error
	FindUserByResetToken(ctx context.Context, token string) (*domain.User, error)
	FindUserById(ctx context.Context, id uint) (*domain.User, error)
	WithTx(ctx context.Context, tx *sql.Tx) UserRepository
}

type userRepository struct {
	dbWrite *sql.DB
	dbRead  *sql.DB
	tx      *sql.Tx
}

func (u *userRepository) CreateUser(ctx context.Context, user *domain.User) error {
	return nil
}

func (u *userRepository) FindUserByEmail(ctx context.Context, email string) (*domain.User, error) {
	return nil, nil
}

func (u *userRepository) FindUserByResetToken(ctx context.Context, token string) (*domain.User, error) {
	return nil, nil
}

func (u *userRepository) SaveUser(ctx context.Context, user *domain.User) error {
	return nil
}

func (u *userRepository) FindUserById(ctx context.Context, id uint) (*domain.User, error) {
	return nil, nil
}

func (u *userRepository) WithTx(ctx context.Context, tx *sql.Tx) UserRepository {
	return &userRepository{
		dbWrite: u.dbWrite,
		dbRead:  u.dbRead,
		tx:      u.tx,
	}
}

func NewUserRepository(dbWrite *sql.DB, dbRead *sql.DB) UserRepository {
	return &userRepository{
		dbWrite: dbWrite,
		dbRead:  dbRead,
	}
}
