package handler

import (
	"net/http"

	"tian-kv/internal/logic"
	"tian-kv/internal/svc"
	"tian-kv/internal/types"

	"github.com/tal-tech/go-zero/rest/httpx"
)

func AddKeyHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AddKeyReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewAddKeyLogic(r.Context(), ctx)
		resp, err := l.AddKey(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
