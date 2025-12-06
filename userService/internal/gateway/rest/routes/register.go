package routes

import (
	"github.com/julienschmidt/httprouter"
	"github.com/saleh-ghazimoradi/StandardMicroEcoBay/internal/gateway/rest/middlewares"
	"github.com/saleh-ghazimoradi/StandardMicroEcoBay/internal/helper"
	"net/http"
)

type Register struct {
	apiError    *helper.APIError
	middleware  *middlewares.Middlewares
	healthRoute *HealthRoutes
	userRoute   *UserRoutes
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

func WithUserRoute(userRoute *UserRoutes) Options {
	return func(r *Register) {
		r.userRoute = userRoute
	}
}

func WithMiddleware(middleware *middlewares.Middlewares) Options {
	return func(r *Register) {
		r.middleware = middleware
	}
}

func (r *Register) RegisterRoutes() http.Handler {
	router := httprouter.New()
	router.NotFound = http.HandlerFunc(r.apiError.NotFoundResponse)
	router.MethodNotAllowed = http.HandlerFunc(r.apiError.MethodNotAllowedResponse)

	r.healthRoute.HealthRoute(router)
	r.userRoute.UserRoute(router)

	return r.middleware.RecoverPanic(r.middleware.RateLimit(router))
}

func NewRegister(opts ...Options) *Register {
	register := &Register{}
	for _, opt := range opts {
		opt(register)
	}
	return register
}
