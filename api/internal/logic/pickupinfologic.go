package logic

import (
	"context"
	"fmt"
	"github.com/zeromicro/x/errors"
	logic2 "wechat-ocr/api/common/logic"
	"wechat-ocr/api/common/model"
	"wechat-ocr/api/internal/svc"
	"wechat-ocr/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PickupInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPickupInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PickupInfoLogic {
	return &PickupInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PickupInfoLogic) PickupInfo(req *types.PickupInfoRequest) (resp *types.PickupInfoResponse, err error) {
	// todo: add your logic here and delete this line
	if req.Imagescode == "" {
		return nil, errors.New(1001, "请求参数不能为空")
	}
	fmt.Printf("参数:%v\n", req.Imagescode)
	var result logic2.PickupDemo

	text, err := logic2.ExtractText(req.Imagescode)
	if err != nil {
		return nil, errors.New(1002, err.Error())
	}
	result, err = logic2.SPickUpInfo(text)
	if err != nil {
		return nil, errors.New(1003, err.Error())
	}
	resp = &types.PickupInfoResponse{
		List: make([]*types.PickupInfoListResponse, 0),
	}
	for _, v := range result.Info {
		newInfo := model.PickupInfo{
			Text:       result.Text,
			NewText:    result.NewText,
			PickupCode: v.PickupCode,
			PickupName: v.PickupName,
		}
		id, err := newInfo.Create(l.svcCtx.UserDb, context.Background())
		if err != nil {
			return nil, errors.New(1004, err.Error())
		}

		resp.List = append(resp.List, &types.PickupInfoListResponse{
			Id:         id,
			PickupCode: v.PickupCode,
			PickupName: v.PickupName,
			Status:     0,
		})
	}

	return
}
