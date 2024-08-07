package service

import (
	"context"
	v1 "uniswap/api/swap/v1"
	"uniswap/internal/biz"
)

type PoolService struct {
	v1.UnimplementedPoolServer

	uc *biz.PoolUsecase
}

func NewPoolService(uc *biz.PoolUsecase) *PoolService {
	return &PoolService{uc: uc}
}

func (s *PoolService) CreatePool(ctx context.Context, req *v1.CreatePoolRequest) (*v1.CreatePoolReply, error) {
	pool, err := s.uc.CreatePool(req.Token0, req.Token1, req.Fee)
	if err != nil {
		return nil, err
	}
	return &v1.CreatePoolReply{}, nil
}
func (s *PoolService) UpdatePool(ctx context.Context, req *v1.UpdatePoolRequest) (*v1.UpdatePoolReply, error) {
	return &v1.UpdatePoolReply{}, nil
}
func (s *PoolService) DeletePool(ctx context.Context, req *v1.DeletePoolRequest) (*v1.DeletePoolReply, error) {
	return &v1.DeletePoolReply{}, nil
}
func (s *PoolService) GetPool(ctx context.Context, req *v1.GetPoolRequest) (*v1.GetPoolReply, error) {
	return &v1.GetPoolReply{}, nil
}
func (s *PoolService) ListPool(ctx context.Context, req *v1.ListPoolRequest) (*v1.ListPoolReply, error) {
	return &v1.ListPoolReply{}, nil
}
