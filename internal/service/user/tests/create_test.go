package tests

//
//import (
//	"context"
//	"github.com/brianvoe/gofakeit/v6"
//	"github.com/gojuno/minimock/v3"
//	"github.com/kom1ssar/grpc-auth/internal/managers"
//	"github.com/kom1ssar/grpc-auth/internal/model"
//	"github.com/kom1ssar/grpc-auth/internal/repository"
//	repositoryMocks "github.com/kom1ssar/grpc-auth/internal/repository/mocks"
//	"github.com/kom1ssar/grpc-auth/internal/service/user"
//	"github.com/stretchr/testify/require"
//	"testing"
//)
//
//func TestCreate(t *testing.T) {
//	type userRepositoryMockFunc func(mc *minimock.Controller) repository.UserRepository
//	type passwordManagerMockFunc func(mc minimock.Controller) managers.PasswordManager
//
//	type args struct {
//		ctx             context.Context
//		req             *model.User
//		passwordConfirm string
//	}
//
//	var (
//		ctx = context.Background()
//		mc  = minimock.NewController(t)
//
//		id       = gofakeit.Int64()
//		name     = gofakeit.BeerName()
//		email    = gofakeit.Email()
//		password = gofakeit.Password(true, true, true, false, false, 5)
//		role     = int32(0)
//
//		req = &model.User{
//			Id:       id,
//			Name:     name,
//			Email:    email,
//			Role:     role,
//			Password: password,
//		}
//	)
//
//	mc.Finish()
//
//	tests := []struct {
//		name           string
//		args           args
//		want           int64
//		err            error
//		userRepository userRepositoryMockFunc
//		passwordManager passwordManagerMockFunc
//	}{
//		{
//			name: "success-case",
//			args: args{
//				ctx:             ctx,
//				req:             req,
//				passwordConfirm: password,
//			},
//			want: id,
//			err:  nil,
//			userRepository: func(mc *minimock.Controller) repository.UserRepository {
//				mock := repositoryMocks.NewUserRepositoryMock(mc)
//				mock.CreateMock.Expect(ctx, req).Return(id, nil)
//
//				return mock
//			},
//			passwordManager: func(mc minimock.Controller) managers.PasswordManager {
//				mock := psswor
//			},
//		},
//	}
//
//	for _, tt := range tests {
//		tt := tt
//
//		t.Run(tt.name, func(t *testing.T) {
//			userRepositoryMock := tt.userRepository(mc)
//			passwordManager := tt.passwordManager(mc)
//
//			api := user.NewUserService(userRepositoryMock)
//
//			resp, err := api.Create(tt.args.ctx, tt.args.req, tt.args.passwordConfirm)
//
//			require.Equal(t, tt.err, err)
//			require.Equal(t, tt.want, resp)
//
//		})
//	}
//} //todo fix test
