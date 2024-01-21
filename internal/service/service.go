package service

import (
	"context"
	"github.com/kom1ssar/grpc-auth/internal/model"
)

type UserService interface {
	GetById(ctx context.Context, id int64) (*model.User, error)
	Create(ctx context.Context, user *model.User, passwordConfirm string) (int64, error)
}

type AuthService interface {
	Login(ctx context.Context, email string, password string) (*model.JwtTokens, error)
	GetRefreshToken(ctx context.Context, refreshToken string) (string, error)
	GetAccessToken(ctx context.Context, refreshToken string) (string, error)
}
