package handlers

import (
	"errors"
	"fmt"
	"github.com/saleh-ghazimoradi/StandardMicroEcoBay/internal/dto"
	"github.com/saleh-ghazimoradi/StandardMicroEcoBay/internal/gateway/rest/auth"
	"github.com/saleh-ghazimoradi/StandardMicroEcoBay/internal/helper"
	"github.com/saleh-ghazimoradi/StandardMicroEcoBay/internal/repository"
	"github.com/saleh-ghazimoradi/StandardMicroEcoBay/internal/service"
	"net/http"
)

type UserHandler struct {
	apiError      *helper.APIError
	authenticator auth.Authenticator
	userService   service.UserService
}

func (u *UserHandler) Register(w http.ResponseWriter, r *http.Request) {
	var payload *dto.UserSignup

	if err := helper.ReadJSON(w, r, &payload); err != nil {
		u.apiError.BadRequestResponse(w, r, err)
		return
	}

	v := helper.NewValidator()
	dto.ValidateUserSignup(v, payload)
	if !v.Valid() {
		u.apiError.FailedValidationResponse(w, r, v.Errors)
		return
	}

	if err := u.userService.Register(r.Context(), payload); err != nil {
		switch {
		case errors.Is(err, repository.ErrDuplicateEmail):
			v.AddError("email", "a user with this email address already exists")
			u.apiError.FailedValidationResponse(w, r, v.Errors)
		default:
			u.apiError.ServerErrorResponse(w, r, err)
		}
		return
	}

	if err := helper.WriteJSON(w, http.StatusCreated, helper.Envelope{"message": "user successfully created"}, nil); err != nil {
		u.apiError.ServerErrorResponse(w, r, err)
	}

}

func (u *UserHandler) Login(w http.ResponseWriter, r *http.Request) {
	var payload *dto.UserLogin

	if err := helper.ReadJSON(w, r, &payload); err != nil {
		u.apiError.BadRequestResponse(w, r, err)
		return
	}

	v := helper.NewValidator()
	dto.ValidateUserLogin(v, payload)
	if !v.Valid() {
		u.apiError.FailedValidationResponse(w, r, v.Errors)
		return
	}

	user, err := u.userService.Login(r.Context(), payload)
	if err != nil {
		//TODO: Make sure about the returned error
		u.apiError.InvalidCredentials(w, r, err)
		return
	}

	token, err := u.authenticator.GenerateToken(user.Id, user.Email)
	if err != nil {
		u.apiError.ServerErrorResponse(w, r, fmt.Errorf("failed to generate token: %w", err))
	}

	if err := helper.WriteJSON(w, http.StatusOK, helper.Envelope{"token": token}, nil); err != nil {
		u.apiError.ServerErrorResponse(w, r, err)
	}

}

func (u *UserHandler) ForgotPassword(w http.ResponseWriter, r *http.Request) {
	var payload *dto.ForgotPassword

	if err := helper.ReadJSON(w, r, &payload); err != nil {
		u.apiError.BadRequestResponse(w, r, err)
		return
	}

	v := helper.NewValidator()
	dto.ValidateForgotPassword(v, payload)
	if !v.Valid() {
		u.apiError.FailedValidationResponse(w, r, v.Errors)
		return
	}

	if err := u.userService.ForgotPassword(r.Context(), payload); err != nil {
		u.apiError.ServerErrorResponse(w, r, err)
		return
	}

	if err := helper.WriteJSON(w, http.StatusOK, helper.Envelope{"message": "password reset link sent"}, nil); err != nil {
		u.apiError.ServerErrorResponse(w, r, err)
	}
}

func (u *UserHandler) SetPassword(w http.ResponseWriter, r *http.Request) {
	var payload *dto.SetPassword

	if err := helper.ReadJSON(w, r, &payload); err != nil {
		u.apiError.BadRequestResponse(w, r, err)
		return
	}

	v := helper.NewValidator()
	dto.ValidateSetPassword(v, payload)
	if !v.Valid() {
		u.apiError.FailedValidationResponse(w, r, v.Errors)
		return
	}

	if err := u.userService.SetPassword(r.Context(), payload); err != nil {
		u.apiError.ServerErrorResponse(w, r, err)
		return
	}

	if err := helper.WriteJSON(w, http.StatusOK, helper.Envelope{"message": "password successfully set"}, nil); err != nil {
		u.apiError.ServerErrorResponse(w, r, err)
	}
}

func (u *UserHandler) CreateProfile(w http.ResponseWriter, r *http.Request) {}

func (u *UserHandler) GetProfile(w http.ResponseWriter, r *http.Request) {}

func (u *UserHandler) Auth(w http.ResponseWriter, r *http.Request) {}

func (u *UserHandler) Me(w http.ResponseWriter, r *http.Request) {}

func NewUserHandler(apiError *helper.APIError, authenticator auth.Authenticator, userService service.UserService) *UserHandler {
	return &UserHandler{
		apiError:      apiError,
		authenticator: authenticator,
		userService:   userService,
	}
}
