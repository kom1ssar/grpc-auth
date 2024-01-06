package user

import (
	"context"
	"github.com/kom1ssar/grpc-auth/internal/model"
)

func (s *userService) GetById(ctx context.Context, id int64) (*model.User, error) {

	user, err := s.userRepo.GetById(ctx, id)
	if err != nil {
		return nil, err
	}

	return user, nil
}
