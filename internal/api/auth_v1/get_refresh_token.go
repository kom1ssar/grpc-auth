package auth_v1

import (
	"context"
	desc "github.com/kom1ssar/grpc-auth/pkg/auth_v1"
)

func (i *Implementation) GetRefreshToken(ctx context.Context, req *desc.GetRefreshTokenRequest) (*desc.GetRefreshTokenResponse, error) {
	return nil, nil
}
