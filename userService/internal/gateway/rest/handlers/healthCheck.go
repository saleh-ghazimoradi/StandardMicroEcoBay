package handlers

import (
	"github.com/saleh-ghazimoradi/StandardMicroEcoBay/config"
	"github.com/saleh-ghazimoradi/StandardMicroEcoBay/internal/helper"
	"log/slog"
	"net/http"
)

type HealthHandler struct {
	logger   *slog.Logger
	config   *config.Config
	apiError *helper.APIError
}

func (h *HealthHandler) HealthCheck(w http.ResponseWriter, r *http.Request) {
	env := helper.Envelope{
		"status": "available",
		"system_info": map[string]string{
			"environment": h.config.Application.Environment,
			"version":     h.config.Application.Version,
		},
	}

	if err := helper.WriteJSON(w, http.StatusOK, env, nil); err != nil {
		h.apiError.ServerErrorResponse(w, r, err)
	}
}

func NewHealthHandler(logger *slog.Logger, config *config.Config, apiError *helper.APIError) *HealthHandler {
	return &HealthHandler{
		logger:   logger,
		config:   config,
		apiError: apiError,
	}
}
