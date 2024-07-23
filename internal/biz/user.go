package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"uniswap/internal/model/entity"
)

type User struct {
	entity.User
}

type UserRepo interface {
	GetUserById(ctx context.Context, id int64) (*User, error)
	SaveUser(ctx context.Context, user *User) (*User, error)
}

type UserUsecase struct {
	repo UserRepo
	log  *log.Helper
}

func NewUserUsecase(repo UserRepo, logger log.Logger) *UserUsecase {
	return &UserUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *UserUsecase) GetUserById(ctx context.Context, id int64) (*User, error) {
	return uc.repo.GetUserById(ctx, id)
}

func (uc *UserUsecase) SaveUser(ctx context.Context, user *User) (*User, error) {
	return uc.repo.SaveUser(ctx, user)
}
