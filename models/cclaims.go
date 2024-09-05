package models

import "github.com/golang-jwt/jwt/v4"

type claims struct {
	Role string `json:"role"`
	jwt.StandardClaims
}
