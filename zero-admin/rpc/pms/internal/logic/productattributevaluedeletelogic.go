package logic

import (
	"context"

	"zero-admin-learn/rpc/pms/internal/svc"
	"zero-admin-learn/rpc/pms/pms"

	"github.com/zeromicro/go-zero/core/logx"
)

type ProductAttributeValueDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewProductAttributeValueDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProductAttributeValueDeleteLogic {
	return &ProductAttributeValueDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ProductAttributeValueDeleteLogic) ProductAttributeValueDelete(in *pms.ProductAttributeValueDeleteReq) (*pms.ProductAttributeValueDeleteResp, error) {
	err := l.svcCtx.PmsProductAttributeValueModel.Delete(in.Id)

	if err != nil {
		return nil, err
	}

	return &pms.ProductAttributeValueDeleteResp{}, nil
}
