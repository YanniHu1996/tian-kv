package handler

import (
	"net/http"
	logic2 "tian-kv/internal/logic"
	svc2 "tian-kv/internal/svc"
	types2 "tian-kv/internal/types"

	"github.com/tal-tech/go-zero/rest/httpx"
)

func TiankvHandler(ctx *svc2.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types2.Request
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic2.NewTiankvLogic(r.Context(), ctx)
		resp, err := l.Tiankv(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
