// Package http provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen/v2 version v2.1.0 DO NOT EDIT.
package http

import (
	"time"
)

const (
	BearerAuthScopes = "bearerAuth.Scopes"
)

// Error Error
type Error struct {
	Message string `json:"message"`
	Slug    string `json:"slug"`
}

// SignInResponse defines model for SignInResponse.
type SignInResponse struct {
	Id         int64   `json:"id"`
	Name       string  `json:"name"`
	Patronymic *string `json:"patronymic,omitempty"`
	Surname    *string `json:"surname,omitempty"`
	Token      string  `json:"token"`
}

// Timestamp A timestamp representing a date and time in RFC3339 format
type Timestamp = time.Time

// User defines model for User.
type User struct {
	Age *int16 `json:"age,omitempty"`

	// CreatedAt A timestamp representing a date and time in RFC3339 format
	CreatedAt Timestamp `json:"created_at"`
	Email     *string   `json:"email,omitempty"`
	Gender    *int16    `json:"gender,omitempty"`
	Id        int64     `json:"id"`
	ImageId   *string   `json:"image_id,omitempty"`

	Name       string  `json:"name"`
	Patronymic *string `json:"patronymic,omitempty"`
	Phone      string  `json:"phone"`
	Surname    *string `json:"surname,omitempty"`
}

// SignInJSONBody defines parameters for SignIn.
type SignInJSONBody struct {
	Password string `json:"password"`
	Phone    string `json:"phone"`
}

// SignUpJSONBody defines parameters for SignUp.
type SignUpJSONBody struct {
	Email    *string `json:"email,omitempty"`
	Name     string  `json:"name"`
	Password string  `json:"password"`
	Phone    string  `json:"phone"`
	Surname  *string `json:"surname,omitempty"`
}

// UpdateUserJSONBody defines parameters for UpdateUser.
type UpdateUserJSONBody struct {
	Age        *int16  `json:"age,omitempty"`
	Email      *string `json:"email,omitempty"`
	Gender     *int16  `json:"gender,omitempty"`
	Name       string  `json:"name"`
	Patronymic *string `json:"patronymic,omitempty"`
	Phone      string  `json:"phone"`
	Surname    *string `json:"surname,omitempty"`
}

// SignInJSONRequestBody defines body for SignIn for application/json ContentType.
type SignInJSONRequestBody SignInJSONBody

// SignUpJSONRequestBody defines body for SignUp for application/json ContentType.
type SignUpJSONRequestBody SignUpJSONBody

// UpdateUserJSONRequestBody defines body for UpdateUser for application/json ContentType.
type UpdateUserJSONRequestBody UpdateUserJSONBody
