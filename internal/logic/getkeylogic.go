package logic

import (
	"context"

	"tian-kv/internal/svc"
	"tian-kv/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type GetKeyLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetKeyLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetKeyLogic {
	return GetKeyLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetKeyLogic) GetKey(req types.GetKeyReq) (*types.GetKeyResp, error) {
	// todo: add your logic here and delete this line

	return &types.GetKeyResp{}, nil
}
