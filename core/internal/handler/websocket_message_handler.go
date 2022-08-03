package handler

import (
	"net/http"

	"gcloud/core/internal/logic"
	"gcloud/core/internal/svc"
	"gcloud/core/internal/types"
	"gcloud/core/internal/ws"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func WebsocketMessageHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.WebsocketMessageRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		hub := ws.NewHub()
		go hub.Run()
		ws.ServeWs(hub, w, r)

		l := logic.NewWebsocketMessageLogic(r.Context(), svcCtx)
		resp, err := l.WebsocketMessage(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
