package model

import "github.com/dgrijalva/jwt-go"

type UserJwtClaims struct {
	jwt.StandardClaims
	UserId int64
	Email  string
	Role   int32
}

type JwtTokens struct {
	RefreshToken string
	AccessToken  string
}
