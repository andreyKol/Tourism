package auth

import (
	"context"
	"tourism/internal/common/errors"
	"tourism/internal/domain"
)

type authContextKey string

var authContextKeyValue authContextKey = "auth"

func SetAuthContext(ctx context.Context, data *domain.AuthContext) context.Context {
	return context.WithValue(ctx, authContextKeyValue, data)
}

func GetAuthContextFromContext(ctx context.Context) (*domain.AuthContext, error) {
	val, ok := ctx.Value(authContextKeyValue).(*domain.AuthContext)
	if !ok || val == nil {
		return nil, errors.NewAuthError("auth context is not presented", "auth")
	}
	return val, nil
}
