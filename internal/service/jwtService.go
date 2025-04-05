package service

import (
	"errors"
	"github.com/MicroMolekula/auth-gateway/internal/config"
	"github.com/MicroMolekula/auth-gateway/internal/model"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

var (
	ErrorInvalidJWTToken = errors.New("invalid JWT token")
)

func VerifyToken(jwt string) (int, error) {
	jwtToken, err := ParseToken(jwt)
	if err != nil {
		return 0, err
	}
	if time.Now().Unix() > jwtToken.Exp {
		return 0, errors.New("jwt token expired")
	}
	return int(jwtToken.Id), nil
}

func ParseToken(tokenString string) (*model.DataJwt, error) {
	token, err := jwt.ParseWithClaims(tokenString, &model.DataJwt{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.Cfg.JWT.Secret), nil
	})

	if err != nil {
		return nil, err
	}

	result, ok := token.Claims.(*model.DataJwt)
	if !ok {
		return nil, ErrorInvalidJWTToken
	}

	return result, nil
}
