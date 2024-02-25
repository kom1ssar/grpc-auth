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

	accessToken, err := a.jwtManager.GenerateAccess(user.Id, user.Email, user.Role)
	if err != nil {
		return nil, err
	}

	refreshToken, err := a.jwtManager.GenerateRefresh(user.Id)
	if err != nil {
		return nil, err
	}

	return &model.JwtTokens{
		RefreshToken: accessToken,
		AccessToken:  refreshToken,
	}, nil
}
