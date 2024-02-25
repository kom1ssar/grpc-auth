package managers

import (
	"context"
	"github.com/kom1ssar/grpc-auth/internal/model"
)

type JWTManager interface {
	GenerateAccess(userId int64, email string, role int32) (string, error)
	GenerateRefresh(useId int64) (string, error)
	VerifyAccess(token string) (*model.UserJwtClaims, error)
	VerifyRefresh(token string) (*model.UserJwtClaims, error)
}

type PasswordManager interface {
	Hash(ctx context.Context, password string) (string, error)
	Compare(ctx context.Context, hashPassword string, password string) bool
}
