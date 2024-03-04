package handler

import (
	xhttp "github.com/zeromicro/x/http"
	"net/http"
	"wechat-ocr/api/internal/logic"
	"wechat-ocr/api/internal/svc"
)

func pingInfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewPingInfoLogic(r.Context(), svcCtx)
		resp, err := l.PingInfo()
		if err != nil {
			// code-data 响应格式
			xhttp.JsonBaseResponseCtx(r.Context(), w, err)
		} else {
			// code-data 响应格式
			xhttp.JsonBaseResponseCtx(r.Context(), w, resp)
		}
	}
}
