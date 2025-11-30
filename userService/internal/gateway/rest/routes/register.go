package routes

import (
	"github.com/julienschmidt/httprouter"
	"github.com/saleh-ghazimoradi/StandardMicroEcoBay/internal/helper"
	"net/http"
)

type Register struct {
	apiError    *helper.APIError
	healthRoute *HealthRoutes
}

type Options func(*Register)

func WithAPIError(apiError *helper.APIError) Options {
	return func(r *Register) {
		r.apiError = apiError
	}
}

func WithHealthRoute(healthRoute *HealthRoutes) Options {
	return func(r *Register) {
		r.healthRoute = healthRoute
	}
}

func (r *Register) RegisterRoutes() http.Handler {
	router := httprouter.New()
	router.NotFound = http.HandlerFunc(r.apiError.NotFoundResponse)
	router.MethodNotAllowed = http.HandlerFunc(r.apiError.MethodNotAllowedResponse)
	r.healthRoute.HealthRoute(router)
	return router
}

func NewRegister(opts ...Options) *Register {
	register := &Register{}
	for _, opt := range opts {
		opt(register)
	}
	return register
}
