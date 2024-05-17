package http

import (
	"Tourism/internal/domain"
	"context"
	"github.com/go-chi/render"
	"net/http"
	"time"
)

type AuthUseCase interface {
	SignIn(ctx context.Context, req *domain.SignInRequest) (*domain.SignInResponse, error)
	SignUp(ctx context.Context, req *domain.SignUpRequest) error
	GetAuthContextByUserID(ctx context.Context, id int64) (*domain.AuthContext, error)
}

// @Summary      SignIn
// @Description  Starts user session. On success returns token and sets HTTP cookie.
// @Description  If authorization will fail, `401 Unauthorized` status code will be returned without any additional data.
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        request body domain.SignInRequest true "Request body."
// @Success      200  {object}	domain.SignInResponse
// @Failure      401  {object}  Error
// @Failure      500  {object}  Error
// @Router       /auth/sign-in [post]
func (h HttpHandler) SignIn(w http.ResponseWriter, r *http.Request) {
	var req SignInJSONRequestBody
	if err := render.DecodeJSON(r.Body, &req); err != nil {
		ErrorResponse(w, r, err)
		return
	}

	response, err := h.authUseCase.SignIn(r.Context(), &domain.SignInRequest{
		Phone:    req.Phone,
		Password: req.Password,
	})
	if err != nil {
		ErrorResponse(w, r, err)
		return
	}
	render.Status(r, http.StatusOK)

	atCookie := http.Cookie{
		Name:    "jwt",
		Value:   response.Token,
		Expires: time.Now().Add(time.Minute * 15),
	}

	http.SetCookie(w, &atCookie)
	render.JSON(w, r, SignInResponse{
		Id:         response.ID,
		Name:       response.Name,
		Patronymic: response.Patronymic,
		Surname:    response.Surname,
		Token:      response.Token,
	})
}

// @Summary      SignUp
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        request body domain.SignUpRequest true "Request body."
// @Success      200  {object}	domain.SignUpResponse
// @Failure      401  {object}  Error
// @Failure      500  {object}  Error
// @Router       /auth/sign-up [post]
func (h HttpHandler) SignUp(w http.ResponseWriter, r *http.Request) {
	var req SignUpJSONRequestBody
	if err := render.DecodeJSON(r.Body, &req); err != nil {
		ErrorResponse(w, r, err)
		return
	}

	err := h.authUseCase.SignUp(r.Context(), &domain.SignUpRequest{
		Name:     req.Name,
		Phone:    req.Phone,
		Password: req.Password,
		Surname:  req.Surname,
		Email:    req.Email,
	})
	if err != nil {
		ErrorResponse(w, r, err)
		return
	}

	w.WriteHeader(http.StatusOK)
}
