package models

import (
	"github.com/dgrijalva/jwt-go"
)

type JwtClaims struct {
	ID int `json:"id"`
	jwt.StandardClaims
}

type Login struct {
	Email    string `validate:"required,email" json:"email"`
	Password string `validate:"required" json:"password"`
}
