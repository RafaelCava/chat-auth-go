package models

import (
	"github.com/dgrijalva/jwt-go"
)

type TokenPair struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type Claims struct {
	UserId string `json:"user_id"`
	jwt.StandardClaims
}
