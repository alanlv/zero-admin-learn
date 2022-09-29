package logic

import (
	"context"
	"zero-admin-learn/api/internal/common/errorx"
	"zero-admin-learn/rpc/ums/umsclient"

	"zero-admin-learn/api/internal/svc"
	"zero-admin-learn/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type MemberRuleSettingDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMemberRuleSettingDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) MemberRuleSettingDeleteLogic {
	return MemberRuleSettingDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MemberRuleSettingDeleteLogic) MemberRuleSettingDelete(req types.DeleteMemberRuleSettingReq) (*types.DeleteMemberRuleSettingResp, error) {
	_, err := l.svcCtx.Ums.MemberRuleSettingDelete(l.ctx, &umsclient.MemberRuleSettingDeleteReq{
		Id: req.Id,
	})

	if err != nil {
		logx.WithContext(l.ctx).Errorf("根据Id: %d,删除积分规则异常:%s", req.Id, err.Error())
		return nil, errorx.NewDefaultError("删除积分规则失败")
	}
	return &types.DeleteMemberRuleSettingResp{
		Code:    "000000",
		Message: "",
	}, nil
}
