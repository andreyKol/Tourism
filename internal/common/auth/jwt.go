package auth

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

var (
	//secret = os.Getenv("JWT_SECRET")
	secret = []byte("secret")
)

type Claims struct {
	ID   int64 `json:"id"`
	Type int16 `json:"type"`
}

func GenerateJWT(claims Claims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  claims.ID,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})

	return token.SignedString(secret)
}

func ValidateJWT(tokenString string) (Claims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		id := claims["id"].(int64)
		typ := claims["type"].(int16)
		return Claims{
			ID:   id,
			Type: typ,
		}, nil
	} else {
		return Claims{}, err
	}
}
