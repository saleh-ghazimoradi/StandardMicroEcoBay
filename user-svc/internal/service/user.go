package service

import (
	"context"
	"github.com/saleh-ghazimoradi/StandardMicroEcoBay/user-svc/internal/domain"
	"github.com/saleh-ghazimoradi/StandardMicroEcoBay/user-svc/internal/dto"
)

type UserService interface {
	Register(ctx context.Context, input *dto.UserSignup) error
	Login(ctx context.Context, input *dto.UserLogin) (*domain.User, error)
	ForgotPassword(ctx context.Context, input *dto.ForgotPassword) error
	SetPassword(ctx context.Context, input *dto.SetPassword) error
	CreateProfile(ctx context.Context, profile *dto.UserProfile) error
	GetProfile(ctx context.Context, id uint) (*domain.User, error)
}

type userService struct{}

func (u *userService) Register(ctx context.Context, input *dto.UserSignup) error {
	return nil
}

func (u *userService) Login(ctx context.Context, input *dto.UserLogin) (*domain.User, error) {
	return nil, nil
}

func (u *userService) ForgotPassword(ctx context.Context, input *dto.ForgotPassword) error {
	return nil
}

func (u *userService) SetPassword(ctx context.Context, input *dto.SetPassword) error {
	return nil
}

func (u *userService) CreateProfile(ctx context.Context, profile *dto.UserProfile) error {
	return nil
}

func (u *userService) GetProfile(ctx context.Context, id uint) (*domain.User, error) {
	return nil, nil
}

func NewUserService() UserService {
	return &userService{}
}
