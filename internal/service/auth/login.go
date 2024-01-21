package auth

import (
	"context"
	"errors"
	"github.com/kom1ssar/grpc-auth/internal/model"
)

func (a *authService) Login(ctx context.Context, email string, password string) (*model.JwtTokens, error) {

	user, err := a.userRepository.GetByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	if !a.passwordManager.Compare(ctx, user.Password, password) {
		return nil, errors.New("password not match")
	}

	return &model.JwtTokens{
		RefreshToken: "token",
		AccessToken:  "token",
	}, nil
}
