package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"uniswap/internal/biz"
	"uniswap/internal/model/entity"

	pb "uniswap/api/user/v1"
)

type UserService struct {
	pb.UnimplementedUserServer

	user *biz.UserUsecase
	log  *log.Helper
}

func NewUserService(user *biz.UserUsecase, logger log.Logger) *UserService {
	return &UserService{user: user, log: log.NewHelper(logger)}
}

func (s *UserService) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserReply, error) {
	u, err := s.user.SaveUser(ctx, &biz.User{
		User: entity.User{
			SettingJson: req.SettingJson,
		},
	})
	if err != nil {
		return nil, err
	}
	return &pb.CreateUserReply{
		Id: uint64(u.UserId),
	}, nil
}
func (s *UserService) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.UpdateUserReply, error) {
	return &pb.UpdateUserReply{}, nil
}
func (s *UserService) DeleteUser(ctx context.Context, req *pb.DeleteUserRequest) (*pb.DeleteUserReply, error) {
	return &pb.DeleteUserReply{}, nil
}
func (s *UserService) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserReply, error) {
	u, err := s.user.GetUserById(ctx, int64(req.Id))
	if err != nil {
		return nil, err
	}
	return &pb.GetUserReply{
		SettingJson: u.SettingJson,
		CreatedAt:   u.CreatedAt.String(),
		UpdatedAt:   u.UpdatedAt.String(),
	}, nil
}
func (s *UserService) ListUser(ctx context.Context, req *pb.ListUserRequest) (*pb.ListUserReply, error) {
	return &pb.ListUserReply{}, nil
}
