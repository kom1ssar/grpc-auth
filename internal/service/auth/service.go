package auth

import (
	"github.com/kom1ssar/grpc-auth/internal/config"
	"github.com/kom1ssar/grpc-auth/internal/managers"
	"github.com/kom1ssar/grpc-auth/internal/repository"
	"github.com/kom1ssar/grpc-auth/internal/service"
)

var _ service.AuthService = (*authService)(nil)

type authService struct {
	userRepository  repository.UserRepository
	passwordManager managers.PasswordManager
	jwtConfig       config.JWTConfig
}

func NewAuthService(userRepository repository.UserRepository, passwordManager managers.PasswordManager, jwtConfig config.JWTConfig) service.AuthService {
	return &authService{
		userRepository:  userRepository,
		passwordManager: passwordManager,
		jwtConfig:       jwtConfig,
	}

}
