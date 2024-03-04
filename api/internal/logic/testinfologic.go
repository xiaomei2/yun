package logic

import (
	"context"

	"wechat-ocr/api/internal/svc"
	"wechat-ocr/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type TestInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewTestInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TestInfoLogic {
	return &TestInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *TestInfoLogic) TestInfo() (resp *types.TestInfoResponse, err error) {
	// todo: add your logic here and delete this line

	return &types.TestInfoResponse{
		Text: "hello world",
	}, nil
}
