package tests

//
//import (
//	"context"
//	"fmt"
//	"github.com/brianvoe/gofakeit/v6"
//	"github.com/gojuno/minimock/v3"
//	"github.com/kom1ssar/grpc-auth/internal/model"
//	"github.com/kom1ssar/grpc-auth/internal/repository"
//	repositoryMocks "github.com/kom1ssar/grpc-auth/internal/repository/mocks"
//	"github.com/kom1ssar/grpc-auth/internal/service/user"
//	"github.com/stretchr/testify/require"
//	"testing"
//)
//
//func TestGetById(t *testing.T) {
//	type userRepositoryMockFunc func(mc *minimock.Controller) repository.UserRepository
//
//	type args struct {
//		ctx             context.Context
//		req             int64
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
//		req = id
//
//		resp = &model.User{
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
//		want           *model.User
//		err            error
//		userRepository userRepositoryMockFunc
//	}{
//		{
//			name: "success-case",
//			args: args{
//				ctx:             ctx,
//				req:             req,
//				passwordConfirm: password,
//			},
//			want: resp,
//			err:  nil,
//			userRepository: func(mc *minimock.Controller) repository.UserRepository {
//				mock := repositoryMocks.NewUserRepositoryMock(mc)
//				mock.GetByIdMock.Expect(ctx, req).Return(resp, nil)
//
//				return mock
//			},
//		},
//	}
//
//	for _, tt := range tests {
//		tt := tt
//
//		t.Run(tt.name, func(t *testing.T) {
//			userRepositoryMock := tt.userRepository(mc)
//
//			api := user.NewUserService(userRepositoryMock)
//
//			res, err := api.GetById(tt.args.ctx, tt.args.req)
//
//			fmt.Printf("%+v", res)
//			require.Equal(t, tt.err, err)
//			require.Equal(t, tt.want, res)
//
//		})
//	}
//} // todo fix test
