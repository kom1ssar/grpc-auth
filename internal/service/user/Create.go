package user

import (
	"context"
	"errors"
	"github.com/kom1ssar/grpc-auth/internal/model"
)

func (s *userService) Create(ctx context.Context, user *model.User, passwordConfirm string) (int64, error) {
	if passwordConfirm != user.Password {
		return 0, errors.New("password and password_confirm do not match")
	}

	id, err := s.userRepo.Create(ctx, user)
	if err != nil {
		return 0, err
	}

	return id, nil

} // todo validate unique email or name
