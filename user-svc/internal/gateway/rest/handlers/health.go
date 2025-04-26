package handlers

import "net/http"

type HealthCheckHandler struct {
}

func (h *HealthCheckHandler) HealthCheck(w http.ResponseWriter, r *http.Request) {
	
}

func NewHealthHandler() *HealthCheckHandler {
	return &HealthCheckHandler{}
}
