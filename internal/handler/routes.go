// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"
	svc2 "tian-kv/internal/svc"

	"github.com/tal-tech/go-zero/rest"
)

func RegisterHandlers(engine *rest.Server, serverCtx *svc2.ServiceContext) {
	engine.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/from/:name",
				Handler: TiankvHandler(serverCtx),
			},
		},
	)
}
