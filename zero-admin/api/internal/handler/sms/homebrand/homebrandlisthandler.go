package handler

import (
	"net/http"

	"zero-admin-learn/api/internal/logic/sms/homebrand"
	"zero-admin-learn/api/internal/svc"
	"zero-admin-learn/api/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func HomeBrandListHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ListHomeBrandReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewHomeBrandListLogic(r.Context(), ctx)
		resp, err := l.HomeBrandList(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
