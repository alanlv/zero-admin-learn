package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"zero-admin-learn/rpc/model/paymodel"
	"zero-admin-learn/rpc/pay/internal/config"
)

type ServiceContext struct {
	c                config.Config
	WxRecordModel    paymodel.PayWxRecordModel
	WxMerchantsModel paymodel.PayWxMerchantsModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	sqlConn := sqlx.NewMysql(c.Mysql.Datasource)
	return &ServiceContext{
		c:                c,
		WxRecordModel:    paymodel.NewPayWxRecordModel(sqlConn),
		WxMerchantsModel: paymodel.NewPayWxMerchantsModel(sqlConn),
	}
}
