package models

import "github.com/golang-jwt/jwt/v4"

type Claims struct {
	Role string `json:"role"`
	jwt.StandardClaims
}
