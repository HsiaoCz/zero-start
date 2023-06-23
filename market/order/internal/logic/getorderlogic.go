package logic

import (
	"context"

	"zero-start/mall/user/types/user"
	"zero-start/market/order/internal/svc"
	"zero-start/market/order/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetOrderLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetOrderLogic {
	return &GetOrderLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetOrderLogic) GetOrder(req *types.GetRequest) (resp *types.GetResponse, err error) {
	// todo: add your logic here and delete this line
	userID := l.getOrderByID(req.ID)
	// 根据用户ID去User服务获取用户信息
	userResponse, err := l.svcCtx.UserRpc.GetUser(context.Background(), &user.IDRequest{
		Id: userID,
	})
	if err != nil {
		return nil, err
	}
	return &types.GetResponse{
		ID:       req.ID,
		Name:     "Hello Order name",
		Username: userResponse.Name,
	}, nil
}

func (l *GetOrderLogic) getOrderByID(id string) string {
	return "1"
}
