package handler

import (
	"net/http"

	"tian-kv/internal/logic"
	"tian-kv/internal/svc"
	"tian-kv/internal/types"

	"github.com/tal-tech/go-zero/rest/httpx"
)

func DeleteKeyHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DeleteKeyReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewDeleteKeyLogic(r.Context(), ctx)
		resp, err := l.DeleteKey(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
