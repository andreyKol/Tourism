package http

import (
	"Tourism/internal/domain"
	"Tourism/internal/handlers/httphelp"
	"context"
	"errors"
	"github.com/go-chi/render"
	"io"
	"net/http"
)

type UserUseCase interface {
	GetUser(ctx context.Context, id int64) (*domain.User, error)
	GetUsersByRole(ctx context.Context, role int16) ([]*domain.User, error)
	UpdateUser(ctx context.Context, user *domain.User) error
	SetUserImage(ctx context.Context, id int64, img []byte) error
}

func FromDomainUserToUser(user *domain.User) *User {
	return &User{
		Age:        user.Age,
		CreatedAt:  user.CreatedAt,
		Email:      user.Email,
		Gender:     user.Gender,
		Id:         user.ID,
		ImageId:    user.ImageID,
		LastOnline: user.LastOnline,
		Name:       user.Name,
		Patronymic: user.Patronymic,
		Phone:      user.Phone,
		Surname:    user.Surname,
	}
}

// @Summary      GetUser
// @Description  Returns information about user
// @Tags         user
// @Accept       json
// @Produce      json
// @Param		 user_id path int true "User ID"
// @Success      200
// @Failure      404  {object}  Error
// @Failure      500  {object}  Error
// @Router       /users/{user_id} [get]
func (h HttpHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	userID := httphelp.ParseParamInt64("user_id", r)
	user, err := h.userUseCase.GetUser(r.Context(), userID)
	if err != nil {
		ErrorResponse(w, r, err)
		return
	}

	render.JSON(w, r, FromDomainUserToUser(user))
}

// @Summary      GetUsersByRole
// @Description  Returns information about users with the specified role
// @Tags         user
// @Accept       json
// @Produce      json
// @Param		 role query int true "User Role"
// @Success      200
// @Failure      400  {object}  Error
// @Failure      404  {object}  Error
// @Failure      500  {object}  Error
// @Router       /users/{role} [get]
func (h HttpHandler) GetUsersByRole(w http.ResponseWriter, r *http.Request) {
	role := httphelp.ParseParamInt64("role", r)
	users, err := h.userUseCase.GetUsersByRole(r.Context(), int16(role))
	if err != nil {
		ErrorResponse(w, r, err)
		return
	}

	render.JSON(w, r, users)
}

// @Summary      UpdateUser
// @Description  Returns information about updated user
// @Tags         user
// @Accept       json
// @Produce      json
// @Param		 request body UpdateUserJSONRequestBody true "Update User Request Body"
// @Success      200
// @Failure      404  {object}  Error
// @Failure      500  {object}  Error
// @Router       /users/update/{user_id} [patch]
func (h HttpHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	userID := httphelp.ParseParamInt64("user_id", r)
	var req UpdateUserJSONRequestBody
	err := render.DecodeJSON(r.Body, &req)
	if err != nil {
		ErrorResponse(w, r, err)
		return
	}

	err = h.userUseCase.UpdateUser(r.Context(), &domain.User{
		ID:         userID,
		Name:       req.Name,
		Surname:    req.Surname,
		Patronymic: req.Patronymic,
		Age:        req.Age,
		Gender:     req.Gender,
		Phone:      req.Phone,
		Email:      req.Email,
	})
	if err != nil {
		ErrorResponse(w, r, err)
		return
	}

	render.Status(r, http.StatusOK)
}

// @Summary      SetUserImage
// @Description  Add image
// @Tags         user
// @Accept       json
// @Produce      json
// @Param		 user_id path int true "User ID"
// @Success      200
// @Failure      404  {object}  Error
// @Failure      500  {object}  Error
// @Router       /users/image/{user_id} [post]
func (h HttpHandler) SetUserImage(w http.ResponseWriter, r *http.Request) {
	userID := httphelp.ParseParamInt64("user_id", r)
	file, _, err := r.FormFile("file")
	if err != nil {
		if errors.Is(err, http.ErrMissingFile) {
			ErrorResponse(w, r, err)
			return
		}
		ErrorResponse(w, r, err)
		return
	}
	defer file.Close()

	content, err := io.ReadAll(file)
	if err != nil {
		ErrorResponse(w, r, err)
		return
	}

	err = h.userUseCase.SetUserImage(r.Context(), userID, content)
	if err != nil {
		ErrorResponse(w, r, err)
		return
	}

	w.WriteHeader(http.StatusOK)
}
