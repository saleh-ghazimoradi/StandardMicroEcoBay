package handlers

import "net/http"

type UserHandler struct{}

func (u *UserHandler) Register(w http.ResponseWriter, r *http.Request) {}

func (u *UserHandler) Login(w http.ResponseWriter, r *http.Request) {}

func (u *UserHandler) ForgotPassword(w http.ResponseWriter, r *http.Request) {}

func (u *UserHandler) SetPassword(w http.ResponseWriter, r *http.Request) {}

func (u *UserHandler) CreateProfile(w http.ResponseWriter, r *http.Request) {}

func (u *UserHandler) GetProfile(w http.ResponseWriter, r *http.Request) {}

func (u *UserHandler) Authentication(w http.ResponseWriter, r *http.Request) {}

func NewUserHandler() *UserHandler {
	return &UserHandler{}
}
