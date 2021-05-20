package logic

import (
	"context"

	"tian-kv/internal/svc"
	"tian-kv/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type DeleteKeyLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteKeyLogic(ctx context.Context, svcCtx *svc.ServiceContext) DeleteKeyLogic {
	return DeleteKeyLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteKeyLogic) DeleteKey(req types.DeleteKeyReq) (*types.DeleteKeyResp, error) {
	// todo: add your logic here and delete this line

	return &types.DeleteKeyResp{}, nil
}
