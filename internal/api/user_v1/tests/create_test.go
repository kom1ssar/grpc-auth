package tests

import (
	"context"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/gojuno/minimock/v3"
	user "github.com/kom1ssar/grpc-auth/internal/api/user_v1"
	"github.com/kom1ssar/grpc-auth/internal/model"
	"github.com/kom1ssar/grpc-auth/internal/service"
	serviceMocks "github.com/kom1ssar/grpc-auth/internal/service/mocks"
	desc "github.com/kom1ssar/grpc-auth/pkg/user_v1"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCreate(t *testing.T) {
	type userServiceMockFunc func(mc *minimock.Controller) service.UserService

	type args struct {
		ctx context.Context
		req *desc.CreateUserRequest
	}

	var (
		ctx = context.Background()
		mc  = minimock.NewController(t)

		id = gofakeit.Int64()

		name            = gofakeit.BeerName()
		email           = gofakeit.Email()
		password        = gofakeit.Password(true, true, true, false, false, 5)
		passwordConfirm = password
		//worngPasswordConfirm = gofakeit.Password(true, true, true, false, false, 6)
		//err = fmt.Errorf("error")

		req = &desc.CreateUserRequest{User: &desc.UserToCreate{
			Name:            name,
			Email:           email,
			Password:        password,
			PasswordConfirm: passwordConfirm,
		}}

		usr = &model.User{
			Name:     name,
			Email:    email,
			Password: password,
		}

		res = &desc.CreateUserResponse{Id: id}
	)

	mc.Finish()

	tests := []struct {
		name            string
		args            args
		want            *desc.CreateUserResponse
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
				mock.CreateMock.Expect(ctx, usr, passwordConfirm).Return(id, nil)

				return mock
			},
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			userServiceMock := tt.userServiceMock(mc)

			api := user.NewImplementation(userServiceMock)

			resp, err := api.Create(tt.args.ctx, tt.args.req)

			require.Equal(t, tt.err, err)
			require.Equal(t, tt.want, resp)

		})
	}

}
