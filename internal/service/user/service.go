package user

import (
	"github.com/kom1ssar/grpc-auth/internal/repository"
	"github.com/kom1ssar/grpc-auth/internal/service"
)

var _ service.UserService = (*userService)(nil)

type userService struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) service.UserService {
	return &userService{userRepo: userRepository}
}
