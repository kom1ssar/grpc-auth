package service

import (
	"context"
	"github.com/kom1ssar/grpc-auth/internal/model"
)

type UserService interface {
	GetById(ctx context.Context, id int64) (*model.User, error)
	Create(ctx context.Context, user *model.User, passwordConfirm string) (int64, error)
}
