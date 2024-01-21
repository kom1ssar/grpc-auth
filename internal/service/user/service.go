package user

import (
	"github.com/kom1ssar/grpc-auth/internal/managers"
	"github.com/kom1ssar/grpc-auth/internal/repository"
	"github.com/kom1ssar/grpc-auth/internal/service"
)

var _ service.UserService = (*userService)(nil)

type userService struct {
	userRepo        repository.UserRepository
	passwordManager managers.PasswordManager
}

func NewUserService(userRepository repository.UserRepository, passwordManager managers.PasswordManager) service.UserService {
	return &userService{
		userRepo:        userRepository,
		passwordManager: passwordManager,
	}
}
