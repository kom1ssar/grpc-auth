package user

import (
	"github.com/kom1ssar/grpc-auth/internal/model"
	desc "github.com/kom1ssar/grpc-auth/pkg/user_v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func ToUserFromDescRegister(user *desc.UserToCreate) *model.User {

	return &model.User{
		Id:       0,
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
	}
}

func ToUserFromService(user *model.User) *desc.User {

	return &desc.User{
		Name:      user.Name,
		Email:     user.Email,
		Role:      desc.UserRole(user.Role),
		CreatedAt: timestamppb.New(user.CreatedAt),
		UpdatedAt: timestamppb.New(user.UpdatedAt),
	}
}
