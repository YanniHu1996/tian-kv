package handler

import (
	"net/http"

	"tian-kv/internal/logic"
	"tian-kv/internal/svc"
	"tian-kv/internal/types"

	"github.com/tal-tech/go-zero/rest/httpx"
)

func GetKeyHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetKeyReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewGetKeyLogic(r.Context(), ctx)
		resp, err := l.GetKey(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
