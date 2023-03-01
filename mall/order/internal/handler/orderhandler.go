package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"zero-start/mall/order/internal/logic"
	"zero-start/mall/order/internal/svc"
	"zero-start/mall/order/internal/types"
)

func OrderHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Request
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewOrderLogic(r.Context(), svcCtx)
		resp, err := l.Order(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
