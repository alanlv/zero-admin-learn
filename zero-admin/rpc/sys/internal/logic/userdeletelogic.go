package logic

import (
	"context"
	"zero-admin-learn/rpc/sys/internal/svc"
	"zero-admin-learn/rpc/sys/sys"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserDeleteLogic {
	return &UserDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserDeleteLogic) UserDelete(in *sys.UserDeleteReq) (*sys.UserDeleteResp, error) {
	_ = l.svcCtx.UserModel.Delete(in.Id)

	return &sys.UserDeleteResp{}, nil
}
