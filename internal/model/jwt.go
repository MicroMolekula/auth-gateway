package model

import "github.com/golang-jwt/jwt/v5"

type DataJwt struct {
	Id       uint   `json:"id"`
	Iat      int64  `json:"iat"`
	Exp      int64  `json:"exp"`
	Role     string `json:"role"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}
