package http

import (
	"net/http"
	"strings"
	"tourism/internal/common/auth"
	"tourism/internal/common/errors"
	"tourism/internal/domain"
)

func (h HttpHandler) AuthMiddleware(next http.Handler) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			headerValue := r.Header.Get("Authorization")
			if len(headerValue) > 7 && strings.ToLower(headerValue[0:6]) == "bearer" {
				ErrorResponse(w, r, errors.NewAuthError("token is not presented", "token"))
				return
			}

			bearerToken := headerValue[7:]
			if bearerToken == "" {
				ErrorResponse(w, r, errors.NewAuthError("token is empty", "token"))
				return
			}

			claims, err := auth.ValidateJWT(bearerToken)
			if err != nil {
				ErrorResponse(w, r, errors.NewAuthError("invalid token", "token"))
				return
			}

			ctx := auth.SetAuthContext(r.Context(), &domain.AuthContext{
				ID:   claims.ID,
				Type: claims.Type,
			})
			r = r.WithContext(ctx)

			next.ServeHTTP(w, r)
		}
		return http.HandlerFunc(fn)
	}
}
