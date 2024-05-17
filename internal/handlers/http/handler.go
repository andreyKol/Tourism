package http

import (
	"Tourism/internal/domain/ws"
	stderrors "errors"
	"net/http"

	"github.com/go-chi/render"

	"Tourism/internal/common/errors"
)

type HttpHandler struct {
	authUseCase AuthUseCase
	userUseCase UserUseCase
	wsUseCase   WsUseCase
	hub         *ws.Hub
}

func NewHandler(
	authUseCase AuthUseCase,
	userUseCase UserUseCase,
	wsUseCase WsUseCase,
	hub *ws.Hub,
) *HttpHandler {
	return &HttpHandler{
		authUseCase: authUseCase,
		userUseCase: userUseCase,
		wsUseCase:   wsUseCase,
		hub:         hub,
	}
}

func ErrorResponse(w http.ResponseWriter, r *http.Request, err error) {
	var (
		domainError   errors.Error
		responseError Error
		statusCode    = http.StatusInternalServerError
	)

	if stderrors.As(err, &domainError) {
		responseError.Message = domainError.Error()
		responseError.Slug = domainError.Slug()

		switch domainError.Type() {
		case errors.ErrorTypeAuth:
			statusCode = http.StatusUnauthorized
		case errors.ErrorTypeNotFound:
			statusCode = http.StatusNotFound
		case errors.ErrorTypeInvalidInput:
			statusCode = http.StatusBadRequest
		default:
			statusCode = http.StatusInternalServerError
		}
	} else {
		responseError.Message = err.Error()
	}

	render.Status(r, statusCode)
	render.JSON(w, r, responseError)
}
