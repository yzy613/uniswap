package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"uniswap/internal/biz"
	"uniswap/internal/dao"
)

var _ biz.UserRepo = (*userRepo)(nil)

type userRepo struct {
	data *Data
	log  *log.Helper
}

func NewUserRepo(data *Data, logger log.Logger) biz.UserRepo {
	return &userRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *userRepo) GetUserById(ctx context.Context, id int64) (*biz.User, error) {
	var user *biz.User

	err := dao.User.Ctx(ctx).
		Where(dao.User.Columns().UserId, id).
		Scan(&user)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errors.NotFound("USER_NOT_FOUND", "user not found")
	}
	return user, nil
}

func (r *userRepo) SaveUser(ctx context.Context, user *biz.User) (*biz.User, error) {
	result, err := dao.User.Ctx(ctx).
		Data(user).
		Save()
	if err != nil {
		return nil, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}
	user.UserId = id
	return user, nil
}
