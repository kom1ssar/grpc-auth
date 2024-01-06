package user

import (
	"context"
	"fmt"
	"github.com/kom1ssar/grpc-auth/internal/converter"
	desc "github.com/kom1ssar/grpc-auth/pkg/user_v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (i *Implementation) Get(ctx context.Context, req *desc.GetUserRequest) (*desc.GetUserResponse, error) {
	user, err := i.userService.GetById(ctx, req.Id)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	userConverted := converter.ToUserFromService(user)
	return &desc.GetUserResponse{
		Id:        user.Id,
		Name:      userConverted.Name,
		Email:     userConverted.Email,
		Role:      userConverted.Role,
		CreatedAt: timestamppb.New(user.CreatedAt),
		UpdatedAt: timestamppb.New(user.CreatedAt),
	}, nil
}
