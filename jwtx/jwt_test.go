package jwtx

import (
	"testing"

	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	UserId string `json:"userId"`
	jwt.RegisteredClaims
}

func TestJwt(t *testing.T) {
	// var claims = Claims{
	// 	UserId: "123",
	// 	RegisteredClaims: jwt.RegisteredClaims{
	// 		ExpiresAt: jwt.NewNumericDate(time.Now().Add(1 * time.Hour)),
	// 		Issuer:    "test",
	// 	},
	// }

	claims := map[string]interface{}{
		"userId": "123",
	}

	jwt, err := CreateMap("123", 3600, claims)
	if err != nil {
		t.Fatalf(err.Error())
	}
	// eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VySWQiOiIxMjMiLCJpc3MiOiJ0ZXN0IiwiZXhwIjoxNjkzMTUwNTEyfQ.SZpw4jy9oR1l9lJ_J6FN3jrxDbpNqB3wP45ShdcTTb8
	t.Logf(jwt)
}

func TestTokenParse(t *testing.T) {
	var token = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2OTYyMzUxODEsInVzZXJJZCI6IjEyMyJ9.QYtOQ19igOKC1iNaAQzoX2H6GIH-lUtMfyerz51-ATE"
	val, err := ParseMap("123", token)
	if err != nil {
		t.Fatalf(err.Error())
	}
	t.Logf("%+v", val)
}
