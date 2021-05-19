package logic

import (
	"context"
	svc2 "tian-kv/internal/svc"
	types2 "tian-kv/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type TiankvLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc2.ServiceContext
}

func NewTiankvLogic(ctx context.Context, svcCtx *svc2.ServiceContext) TiankvLogic {
	return TiankvLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *TiankvLogic) Tiankv(req types2.Request) (*types2.Response, error) {
	// todo: add your logic here and delete this line

	return &types2.Response{}, nil
}
