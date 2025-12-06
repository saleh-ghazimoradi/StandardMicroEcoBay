package routes

import (
	"github.com/julienschmidt/httprouter"
	"github.com/saleh-ghazimoradi/StandardMicroEcoBay/internal/gateway/rest/auth"
	"github.com/saleh-ghazimoradi/StandardMicroEcoBay/internal/gateway/rest/handlers"
	"net/http"
)

type UserRoutes struct {
	userHandler   *handlers.UserHandler
	authenticator auth.Authenticator
}

func (u *UserRoutes) UserRoute(router *httprouter.Router) {
	router.HandlerFunc(http.MethodPost, "/v1/register", u.userHandler.Register)
	router.HandlerFunc(http.MethodPost, "/v1/login", u.userHandler.Login)
	router.HandlerFunc(http.MethodPost, "/v1/forgot-password", u.userHandler.ForgotPassword)
	router.HandlerFunc(http.MethodPost, "/v1/set-password", u.userHandler.SetPassword)
	//TODO: Set Authentication Middleware to the endpoints below
	router.HandlerFunc(http.MethodPost, "/v1/profile", u.userHandler.CreateProfile)
	router.HandlerFunc(http.MethodGet, "/v1/profile", u.userHandler.GetProfile)
	router.HandlerFunc(http.MethodGet, "/v1/auth", u.userHandler.Auth)
	router.HandlerFunc(http.MethodGet, "/v1/me", u.userHandler.Me)
}

func NewUserRoutes(userHandler *handlers.UserHandler, authenticator auth.Authenticator) *UserRoutes {
	return &UserRoutes{
		userHandler:   userHandler,
		authenticator: authenticator,
	}
}
