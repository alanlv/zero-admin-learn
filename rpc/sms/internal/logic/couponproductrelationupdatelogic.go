package logic

import (
	"context"
	"zero-admin-learn/rpc/model/smsmodel"

	"zero-admin-learn/rpc/sms/internal/svc"
	"zero-admin-learn/rpc/sms/sms"

	"github.com/zeromicro/go-zero/core/logx"
)

type CouponProductRelationUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCouponProductRelationUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CouponProductRelationUpdateLogic {
	return &CouponProductRelationUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CouponProductRelationUpdateLogic) CouponProductRelationUpdate(in *sms.CouponProductRelationUpdateReq) (*sms.CouponProductRelationUpdateResp, error) {
	err := l.svcCtx.SmsCouponProductRelationModel.Update(smsmodel.SmsCouponProductRelation{
		Id:          in.Id,
		CouponId:    in.CouponId,
		ProductId:   in.ProductId,
		ProductName: in.ProductName,
		ProductSn:   in.ProductSn,
	})
	if err != nil {
		return nil, err
	}
	return &sms.CouponProductRelationUpdateResp{}, nil
}
