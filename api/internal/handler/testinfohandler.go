package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"wechat-ocr/api/internal/logic"
	"wechat-ocr/api/internal/svc"
)

func testInfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewTestInfoLogic(r.Context(), svcCtx)
		resp, err := l.TestInfo()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
