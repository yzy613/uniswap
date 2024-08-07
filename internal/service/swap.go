package service

import (
	"context"
	"github.com/shopspring/decimal"
	v1 "uniswap/api/swap/v1"
	"uniswap/internal/biz"
)

type SwapService struct {
	v1.UnimplementedSwapServer

	uc *biz.PoolUsecase
}

func NewSwapService(uc *biz.PoolUsecase) *SwapService {
	return &SwapService{uc: uc}
}

func (s *SwapService) CreateSwap(ctx context.Context, req *v1.CreateSwapRequest) (*v1.CreateSwapReply, error) {
	pool, err := s.uc.GetPool(req.Token0, req.Token1, req.Fee)
	if err != nil {
		return nil, err
	}
	amountSpecified, err := decimal.NewFromString(req.AmountSpecified)
	if err != nil {
		return nil, err
	}
	priceLimit, err := decimal.NewFromString(req.PriceLimit)
	if err != nil {
		return nil, err
	}
	amount0, amount1, err := s.uc.Swap(*pool, req.Recipient, req.ZeroForOne, amountSpecified, priceLimit)
	if err != nil {
		return nil, err
	}
	return &v1.CreateSwapReply{}, nil
}
func (s *SwapService) UpdateSwap(ctx context.Context, req *v1.UpdateSwapRequest) (*v1.UpdateSwapReply, error) {
	return &v1.UpdateSwapReply{}, nil
}
func (s *SwapService) DeleteSwap(ctx context.Context, req *v1.DeleteSwapRequest) (*v1.DeleteSwapReply, error) {
	return &v1.DeleteSwapReply{}, nil
}
func (s *SwapService) GetSwap(ctx context.Context, req *v1.GetSwapRequest) (*v1.GetSwapReply, error) {
	return &v1.GetSwapReply{}, nil
}
func (s *SwapService) ListSwap(ctx context.Context, req *v1.ListSwapRequest) (*v1.ListSwapReply, error) {
	return &v1.ListSwapReply{}, nil
}
