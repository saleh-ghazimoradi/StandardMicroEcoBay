package service

import (
	"context"
	"github.com/saleh-ghazimoradi/StandardMicroEcoBay/internal/domain"
	"github.com/saleh-ghazimoradi/StandardMicroEcoBay/internal/dto"
	"github.com/saleh-ghazimoradi/StandardMicroEcoBay/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	Register(ctx context.Context, input *dto.UserSignup) error
	Login(ctx context.Context, input *dto.UserLogin) (*domain.User, error)
	ForgotPassword(ctx context.Context, input *dto.ForgotPassword) error
	SetPassword(ctx context.Context, input *dto.SetPassword) error
	CreateProfile(ctx context.Context, input *dto.UserProfile) error
	GetProfile(ctx context.Context, userId int64) (*domain.User, error)
}

type userService struct {
	userRepository repository.UserRepository
}

func (u *userService) Register(ctx context.Context, input *dto.UserSignup) error {
	hashedPassword, err := u.generatePasswordHash(input.Password)
	if err != nil {
		return err
	}

	user := &domain.User{
		Email:    input.Email,
		Password: hashedPassword,
		Phone:    input.Phone,
	}

	if err := u.userRepository.CreateUser(ctx, user); err != nil {
		return err
	}
	return nil
}

func (u *userService) Login(ctx context.Context, input *dto.UserLogin) (*domain.User, error) {
	user, err := u.getUserByEmail(ctx, input.Email)
	if err != nil {
		return nil, err
	}

	if err = u.verifyPassword(input.Password, user.Password); err != nil {
		return nil, err
	}

	return user, nil
}

func (u *userService) ForgotPassword(ctx context.Context, input *dto.ForgotPassword) error {
	user, err := u.getUserByEmail(ctx, input.Email)
	if err != nil {
		return err
	}

	resetToken, err := u.generatePasswordHash(user.Email)
	if err != nil {
		return err
	}

	user.ResetToken = resetToken

	if err := u.userRepository.SaveUser(ctx, user); err != nil {
		return err
	}

	return nil
}

func (u *userService) SetPassword(ctx context.Context, input *dto.SetPassword) error {
	user, err := u.findUserByResetToken(ctx, input.Token)
	if err != nil {
		return err
	}

	hashedPassword, err := u.generatePasswordHash(input.Password)
	if err != nil {
		return err
	}

	user.Password = hashedPassword
	user.ResetToken = ""

	if err := u.userRepository.SaveUser(ctx, user); err != nil {
		return err
	}
	return nil
}

func (u *userService) CreateProfile(ctx context.Context, input *dto.UserProfile) error {
	user, err := u.getUserById(ctx, input.UserId)
	if err != nil {
		return err
	}

	if input.FirstName != nil {
		user.FirstName = *input.FirstName
	}

	if input.LastName != nil {
		user.LastName = *input.LastName
	}

	if input.Email != nil {
		user.Email = *input.Email
	}

	if input.Phone != nil {
		user.Phone = *input.Phone
	}

	if input.Address.AddressLine1 != nil {
		user.Address.AddressLine1 = *input.Address.AddressLine1
	}

	if input.Address.AddressLine2 != nil {
		user.Address.AddressLine2 = *input.Address.AddressLine2
	}

	if input.Address.City != nil {
		user.Address.City = *input.Address.City
	}

	if input.Address.Country != nil {
		user.Address.Country = *input.Address.Country
	}

	if input.Address.PostCode != nil {
		user.Address.PostCode = *input.Address.PostCode
	}

	if err := u.userRepository.SaveUser(ctx, user); err != nil {
		return err
	}

	return nil
}

func (u *userService) GetProfile(ctx context.Context, userId int64) (*domain.User, error) {
	user, err := u.getUserById(ctx, userId)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *userService) getUserById(ctx context.Context, userId int64) (*domain.User, error) {
	return u.userRepository.FindUserById(ctx, userId)
}

func (u *userService) getUserByEmail(ctx context.Context, email string) (*domain.User, error) {
	return u.userRepository.FindUserByEmail(ctx, email)
}

func (u *userService) findUserByResetToken(ctx context.Context, token string) (*domain.User, error) {
	return u.userRepository.FindUserByResetToken(ctx, token)
}

func (u *userService) generatePasswordHash(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func (u *userService) verifyPassword(password, hashedPassword string) error {
	if err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password)); err != nil {
		return err
	}
	return nil
}

func NewUserService(userRepository repository.UserRepository) UserService {
	return &userService{
		userRepository: userRepository,
	}
}
