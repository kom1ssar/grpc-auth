package repository

import (
	"context"
	"github.com/kom1ssar/grpc-auth/internal/model"
)

type UserRepository interface {
	GetById(ctx context.Context, id int64) (*model.User, error)
	Create(ctx context.Context, user *model.User) (int64, error)
}
