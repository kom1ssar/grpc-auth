package auth_v1

import (
	"context"
	desc "github.com/kom1ssar/grpc-auth/pkg/auth_v1"
)

func (i *Implementation) GetAccessToken(ctx context.Context, req *desc.GetAccessTokenRequest) (*desc.GetAccessTokenResponse, error) {
	return nil, nil

}
