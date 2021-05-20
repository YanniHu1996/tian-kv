package logic

import (
	"context"

	"tian-kv/internal/svc"
	"tian-kv/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type AddKeyLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddKeyLogic(ctx context.Context, svcCtx *svc.ServiceContext) AddKeyLogic {
	return AddKeyLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddKeyLogic) AddKey(req types.AddKeyReq) (*types.AddKeyResp, error) {
	// todo: add your logic here and delete this line

	return &types.AddKeyResp{}, nil
}
