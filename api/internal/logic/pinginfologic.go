package logic

import (
	"context"
	"github.com/zeromicro/x/errors"

	"wechat-ocr/api/internal/svc"
	"wechat-ocr/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PingInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPingInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PingInfoLogic {
	return &PingInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PingInfoLogic) PingInfo() (resp *types.PingInfoResponse, err error) {
	// todo: add your logic here and delete this line

	return nil, errors.New(1001, "用户名或密码出错")
}
