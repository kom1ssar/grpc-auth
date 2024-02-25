package utils

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/kom1ssar/grpc-auth/internal/model"
	"github.com/pkg/errors"
)

func GenerateToken(claims *model.UserJwtClaims, secretKey []byte) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(secretKey)
}

func VerifyToken(tokenStr string, secretKey []byte) (*model.UserJwtClaims, error) {

	token, err := jwt.ParseWithClaims(
		tokenStr,
		&model.UserJwtClaims{},
		func(token *jwt.Token) (interface{}, error) {
			_, ok := token.Method.(*jwt.SigningMethodHMAC)
			if !ok {
				return nil, errors.Errorf("unexpected token signing method")
			}
			return secretKey, nil

		},
	)
	if err != nil {
		return nil, errors.Errorf("invalid token: %s", err.Error())

	}

	claims, ok := token.Claims.(*model.UserJwtClaims)
	if !ok {
		return nil, errors.Errorf("invalid token claims")
	}

	return claims, nil

}
