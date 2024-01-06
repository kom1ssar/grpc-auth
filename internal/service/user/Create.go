package user

import (
	"context"
	"github.com/kom1ssar/grpc-auth/internal/model"
)

func (s *userService) Create(ctx context.Context, user *model.User) (int64, error) {

	id, err := s.userRepo.Create(ctx, user)
	if err != nil {
		return 0, err
	}

	return id, nil

}
