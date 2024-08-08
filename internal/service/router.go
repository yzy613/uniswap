package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/shopspring/decimal"
	"uniswap/internal/biz"

	pb "uniswap/api/router/v1"
)

type RouterService struct {
	pb.UnimplementedRouterServer

	router *biz.RouterUsecase
}

func NewRouterService(router *biz.RouterUsecase) *RouterService {
	return &RouterService{router: router}
}

func (s *RouterService) ExactInputSingle(ctx context.Context, req *pb.ExactInputSingleRequest) (*pb.ExactInputSingleReply, error) {
	if gtime.New(req.Deadline).Before(gtime.Now()) {
		return nil, errors.BadRequest("router.exactinputsingle", "deadline exceeded")
	}

	amountIn, err := decimal.NewFromString(req.AmountIn)
	if err != nil {
		return nil, errors.BadRequest("router.exactinputsingle", "invalid amount in")
	}
	priceLimit, err := decimal.NewFromString(req.PriceLimit)
	if err != nil {
		return nil, errors.BadRequest("router.exactinputsingle", "invalid price limit")
	}

	amountOut, err := s.router.ExactInputInternal(
		amountIn,
		req.Recipient,
		priceLimit,
		biz.SwapCallbackData{
			Path: biz.Path{
				TokenA: req.TokenIn,
				TokenB: req.TokenOut,
				Fee:    req.Fee,
			},
			Payer: "msg.Sender",
		},
	)
	if err != nil {
		return nil, errors.InternalServer("router.exactinputsingle", err.Error())
	}

	reply := &pb.ExactInputSingleReply{
		AmountOut: amountOut.String(),
	}

	return reply, nil
}
func (s *RouterService) ExactOutputSingle(ctx context.Context, req *pb.ExactOutputSingleRequest) (*pb.ExactOutputSingleReply, error) {
	if gtime.New(req.Deadline).Before(gtime.Now()) {
		return nil, errors.BadRequest("router.exactoutputsingle", "deadline exceeded")
	}

	amountOut, err := decimal.NewFromString(req.AmountOut)
	if err != nil {
		return nil, errors.BadRequest("router.exactoutputsingle", "invalid amount out")
	}
	priceLimit, err := decimal.NewFromString(req.PriceLimit)
	if err != nil {
		return nil, errors.BadRequest("router.exactoutputsingle", "invalid price limit")
	}

	amountIn, err := s.router.ExactOutputInternal(
		amountOut,
		req.Recipient,
		priceLimit,
		biz.SwapCallbackData{
			Path: biz.Path{
				TokenA: req.TokenOut,
				TokenB: req.TokenIn,
				Fee:    req.Fee,
			},
			Payer: "msg.Sender",
		},
	)
	if err != nil {
		return nil, errors.InternalServer("router.exactoutputsingle", err.Error())
	}

	reply := &pb.ExactOutputSingleReply{
		AmountIn: amountIn.String(),
	}

	return reply, nil
}
