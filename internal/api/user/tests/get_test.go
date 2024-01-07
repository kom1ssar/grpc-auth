package tests

import (
	"context"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/gojuno/minimock/v3"
	"github.com/kom1ssar/grpc-auth/internal/api/user"
	"github.com/kom1ssar/grpc-auth/internal/model"
	"github.com/kom1ssar/grpc-auth/internal/service"
	serviceMocks "github.com/kom1ssar/grpc-auth/internal/service/mocks"
	desc "github.com/kom1ssar/grpc-auth/pkg/user_v1"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/types/known/timestamppb"
	"testing"
	"time"
)

func TestGet(t *testing.T) {
	type userServiceMockFunc func(mc *minimock.Controller) service.UserService

	type args struct {
		ctx context.Context
		req *desc.GetUserRequest
	}

	var (
		ctx = context.Background()
		mc  = minimock.NewController(t)

		id        = gofakeit.Int64()
		name      = gofakeit.BeerName()
		email     = gofakeit.Email()
		password  = gofakeit.Password(true, true, true, false, false, 5)
		role      = int32(0)
		createdAt = time.Time{}
		updatedAt = time.Time{}

		req = &desc.GetUserRequest{Id: id}

		usr = &model.User{
			Id:        id,
			Name:      name,
			Email:     email,
			Role:      role,
			Password:  password,
			CreatedAt: createdAt,
			UpdatedAt: updatedAt,
		}

		res = &desc.GetUserResponse{
			Id:        id,
			Name:      name,
			Email:     email,
			Role:      desc.UserRole(role),
			CreatedAt: timestamppb.New(createdAt),
			UpdatedAt: timestamppb.New(updatedAt),
		}
	)

	defer t.Cleanup(mc.Finish)

	tests := []struct {
		name            string
		args            args
		want            *desc.GetUserResponse
		err             error
		userServiceMock userServiceMockFunc
	}{
		{
			name: "success-case",
			args: args{
				ctx: ctx,
				req: req,
			},
			want: res,
			err:  nil,
			userServiceMock: func(mc *minimock.Controller) service.UserService {
				mock := serviceMocks.NewUserServiceMock(mc)
				mock.GetByIdMock.Expect(ctx, id).Return(usr, nil)

				return mock
			},
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			userServiceMock := tt.userServiceMock(mc)

			api := user.NewImplementation(userServiceMock)

			resp, err := api.Get(tt.args.ctx, tt.args.req)

			require.Equal(t, tt.err, err)
			require.Equal(t, tt.want, resp)

		})
	}
}
