package jwtx

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func Create(key string, claims jwt.Claims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(key))
	return tokenString, err
}

func Parse(key string, tokenString string, c jwt.Claims) error {
	token, err := jwt.ParseWithClaims(tokenString, c, func(token *jwt.Token) (interface{}, error) {
		return []byte(key), nil
	})
	if err != nil {
		return err
	}
	if token.Valid {
		return nil
	}
	return jwt.ErrSignatureInvalid
}

func CreateMap(key string, exp int64, claims map[string]interface{}) (string, error) {
	claims["exp"] = time.Now().Add(time.Duration(exp) * time.Second).Unix()
	return Create(key, jwt.MapClaims(claims))
}

func ParseMap(key string, tokenString string) (map[string]interface{}, error) {
	var claims jwt.MapClaims
	err := Parse(key, tokenString, &claims)
	if err != nil {
		return nil, err
	}
	return claims, nil
}
