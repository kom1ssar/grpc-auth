package auth_v1

import (
	"github.com/kom1ssar/grpc-auth/internal/service"
	desc "github.com/kom1ssar/grpc-auth/pkg/auth_v1"
)

type Implementation struct {
	desc.UnimplementedAuthServiceServer
	authService service.AuthService
}

func NewImplementation(authService service.AuthService) *Implementation {
	return &Implementation{
		authService: authService,
	}
}
