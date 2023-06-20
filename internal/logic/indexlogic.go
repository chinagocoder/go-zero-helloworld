package logic

import (
	"context"

	"go-zero-helloworld/internal/svc"
	"go-zero-helloworld/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type IndexLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewIndexLogic(ctx context.Context, svcCtx *svc.ServiceContext) *IndexLogic {
	return &IndexLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *IndexLogic) Index(req *types.IndexReq) (resp *types.IndexResp, err error) {

	return &types.IndexResp{
		Data: "hello world!",
	}, nil
}
