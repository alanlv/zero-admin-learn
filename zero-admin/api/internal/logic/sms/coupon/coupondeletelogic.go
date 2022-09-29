package logic

import (
	"context"
	"zero-admin-learn/api/internal/common/errorx"
	"zero-admin-learn/rpc/sms/smsclient"

	"zero-admin-learn/api/internal/svc"
	"zero-admin-learn/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CouponDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCouponDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) CouponDeleteLogic {
	return CouponDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CouponDeleteLogic) CouponDelete(req types.DeleteCouponReq) (*types.DeleteCouponResp, error) {
	_, err := l.svcCtx.Sms.CouponDelete(l.ctx, &smsclient.CouponDeleteReq{
		Id: req.Id,
	})

	if err != nil {
		logx.WithContext(l.ctx).Errorf("根据Id: %d,删除优惠券异常:%s", req.Id, err.Error())
		return nil, errorx.NewDefaultError("删除优惠券失败")
	}
	return &types.DeleteCouponResp{
		Code:    "000000",
		Message: "",
	}, nil
}
