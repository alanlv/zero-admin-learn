package logic

import (
	"context"
	"zero-admin-learn/rpc/model/smsmodel"

	"zero-admin-learn/rpc/sms/internal/svc"
	"zero-admin-learn/rpc/sms/sms"

	"github.com/zeromicro/go-zero/core/logx"
)

type HomeNewProductUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewHomeNewProductUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HomeNewProductUpdateLogic {
	return &HomeNewProductUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *HomeNewProductUpdateLogic) HomeNewProductUpdate(in *sms.HomeNewProductUpdateReq) (*sms.HomeNewProductUpdateResp, error) {
	err := l.svcCtx.SmsHomeNewProductModel.Update(smsmodel.SmsHomeNewProduct{
		Id:              in.Id,
		ProductId:       in.ProductId,
		ProductName:     in.ProductName,
		RecommendStatus: in.RecommendStatus,
		Sort:            in.Sort,
	})
	if err != nil {
		return nil, err
	}

	return &sms.HomeNewProductUpdateResp{}, nil
}
