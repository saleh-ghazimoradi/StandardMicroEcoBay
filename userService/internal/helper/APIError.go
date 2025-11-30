package helper

import (
	"fmt"
	"log/slog"
	"net/http"
)

type APIError struct {
	logger *slog.Logger
}

func (a *APIError) LogError(r *http.Request, err error) {
	var (
		method = r.Method
		uri    = r.URL.RequestURI()
	)
	a.logger.Error(err.Error(), "method", method, "uri", uri)
}

func (a *APIError) ErrorResponse(w http.ResponseWriter, r *http.Request, status int, message any) {
	env := Envelope{"error": message}

	if err := WriteJSON(w, status, env, nil); err != nil {
		a.LogError(r, err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func (a *APIError) ServerErrorResponse(w http.ResponseWriter, r *http.Request, err error) {
	a.LogError(r, err)
	message := "the server encountered a problem and could not process your request"
	a.ErrorResponse(w, r, http.StatusInternalServerError, message)
}

func (a *APIError) NotFoundResponse(w http.ResponseWriter, r *http.Request) {
	message := "the requested resource could not be found"
	a.ErrorResponse(w, r, http.StatusNotFound, message)
}

func (a *APIError) MethodNotAllowedResponse(w http.ResponseWriter, r *http.Request) {
	message := fmt.Sprintf("the %s method is not supported for this resource", r.Method)
	a.ErrorResponse(w, r, http.StatusMethodNotAllowed, message)
}

func (a *APIError) BadRequestResponse(w http.ResponseWriter, r *http.Request, err error) {
	a.ErrorResponse(w, r, http.StatusBadRequest, err.Error())
}

func (a *APIError) FailedValidationResponse(w http.ResponseWriter, r *http.Request, errors map[string]string) {
	a.ErrorResponse(w, r, http.StatusUnprocessableEntity, errors)
}

func (a *APIError) EditConflictResponse(w http.ResponseWriter, r *http.Request) {
	message := "unable to update the record due to an edit conflict, please try again"
	a.ErrorResponse(w, r, http.StatusConflict, message)
}

func (a *APIError) RateLimitExceededResponse(w http.ResponseWriter, r *http.Request) {
	message := "rate limit exceeded, please try again"
	a.ErrorResponse(w, r, http.StatusTooManyRequests, message)
}

func NewAPIError(logger *slog.Logger) *APIError {
	return &APIError{
		logger: logger,
	}
}
