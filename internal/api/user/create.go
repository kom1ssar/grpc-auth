package user

import (
	"context"
	desc "github.com/kom1ssar/grpc-auth/pkg/user_v1"
)

func (i *Implementation) Create(ctx context.Context, req *desc.CreateUserRequest) (*desc.CreateUserResponse, error) {
	//s.userService.Create(ctx,converter.ToUserFromDesc(req.User))
	//user := model.User{
	//	Name:     req.User.Name,
	//	Email:    req.User.Email,
	//	Password: req.User.Password,
	//}
	//
	//fmt.Printf("%+v", user)
	//
	//id, err := s.userRepo.Create(ctx, &user)
	//if err != nil {
	//	fmt.Println("asdasd")
	//	fmt.Println(err)
	//	return nil, errors.New("pdior")
	//}
	//
	//return &desc.CreateUserResponse{Id: int64(id)}, nil

	return nil, nil
}
