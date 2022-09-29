package handler

import (
	"net/http"

	"zero-admin-learn/api/internal/logic/member/statistics"
	"zero-admin-learn/api/internal/svc"
	"zero-admin-learn/api/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func MemberStatisticsInfoListHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ListMemberStatisticsInfoReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewMemberStatisticsInfoListLogic(r.Context(), ctx)
		resp, err := l.MemberStatisticsInfoList(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
