package middleware

import (
	"fmt"
	"github.com/saleh-ghazimoradi/StandardMicroEcoBay/config"
	"github.com/saleh-ghazimoradi/StandardMicroEcoBay/internal/helper"
	"net/http"
)

type Middleware struct {
	config   *config.Config
	apiError *helper.APIError
}

func (m *Middleware) RecoverPanic(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				w.Header().Set("Connection", "close")
				m.apiError.ServerErrorResponse(w, r, fmt.Errorf("%v", err))
			}
		}()
		next.ServeHTTP(w, r)
	})
}

func NewMiddleware() *Middleware {
	return &Middleware{}
}
