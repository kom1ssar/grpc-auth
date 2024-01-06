package converter

import (
	"github.com/kom1ssar/grpc-auth/internal/model"
	desc "github.com/kom1ssar/grpc-auth/pkg/user_v1"
	"time"
)

func ToUserFromService(user *model.User) *desc.User {

	return &desc.User{
		Name:            user.Name,
		Email:           user.Email,
		Password:        user.Password,
		PasswordConfirm: user.Password,
		Role:            desc.UserRole(user.Role),
	}
}

func ToUserFromDesc(user *desc.User) *model.User {

	return &model.User{
		Id:        0,
		Name:      user.Name,
		Email:     user.Email,
		Role:      int32(user.Role.Number()),
		Password:  user.Password,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
	}
}
