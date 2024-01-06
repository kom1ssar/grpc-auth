package user

import (
	"context"
	"github.com/kom1ssar/grpc-auth/internal/converter"
	desc "github.com/kom1ssar/grpc-auth/pkg/user_v1"
)

func (i *Implementation) Create(ctx context.Context, req *desc.CreateUserRequest) (*desc.CreateUserResponse, error) {
	id, err := i.userService.Create(ctx, converter.ToUserFromDesc(req.User), req.User.PasswordConfirm)
	if err != nil {
		return nil, err
	}

	return &desc.CreateUserResponse{Id: id}, nil
}
