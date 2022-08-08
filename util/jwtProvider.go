package util

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
)

const SecretKey = "secret"

func GenerateJwt(issuer string) (string, error) {
	// JWT Provider or Make JWT
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		Issuer: issuer,
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
	})

	// Send token has been generated
	return claims.SignedString([]byte(SecretKey))
}

func ParseJwt(cookie string) (string, error) {
	// Parse Jwt from cookie
	token, err := jwt.ParseWithClaims(cookie, &jwt.RegisteredClaims{}, func (token *jwt.Token) (interface{}, error) {
		return 	[]byte(SecretKey), nil
	})
	if err != nil || !token.Valid {
		return "", err
	}
	claims := token.Claims.(*jwt.RegisteredClaims)
	
	// Send Issuer
	return claims.Issuer, nil
}