package converter

import (
	"github.com/kom1ssar/grpc-auth/internal/model"
	modelRepo "github.com/kom1ssar/grpc-auth/internal/repository/user/model"
)

func ToUserFromRepo(user *modelRepo.User) *model.User {

	return &model.User{
		Id:        user.Id,
		Name:      user.Name,
		Email:     user.Email,
		Role:      user.Role,
		Password:  user.Password,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}
