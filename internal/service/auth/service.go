package auth

import (
	"github.com/kom1ssar/grpc-auth/internal/managers"
	"github.com/kom1ssar/grpc-auth/internal/repository"
	"github.com/kom1ssar/grpc-auth/internal/service"
)

var _ service.AuthService = (*authService)(nil)

type authService struct {
	userRepository  repository.UserRepository
	passwordManager managers.PasswordManager
	jwtManager      managers.JWTManager
}

func NewAuthService(userRepository repository.UserRepository, passwordManager managers.PasswordManager, jwtManager managers.JWTManager) service.AuthService {
	return &authService{
		userRepository:  userRepository,
		passwordManager: passwordManager,
		jwtManager:      jwtManager,
	}

}
