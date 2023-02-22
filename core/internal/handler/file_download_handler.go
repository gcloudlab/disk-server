package handler

import (
	"net/http"

	"gcloud/core/helper"
	"gcloud/core/internal/svc"
	"gcloud/core/internal/types"
	"gcloud/core/models"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func FileDownloadHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.FileDownloadRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		rp := new(models.RepositoryPool)
		// https://gcloud-1303456836.cos.ap-chengdu.myqcloud.com/
		svcCtx.Engine.
			Where("path = ?", req.Path).
			First(rp)
		if rp.Id == 0 {
			httpx.OkJson(w, &types.FileDownloadReply{
				Msg: "file not exits",
			})
			return
		}

		fb, err := helper.CosDownload(r, req.Path, req.Name)
		if err != nil {
			httpx.Error(w, err)
		}

		// l := logic.NewFileDownloadLogic(r.Context(), svcCtx)
		// resp, err := l.FileDownload(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, &types.FileDownloadReply{
				Data: fb,
				Msg:  "success",
			})
		}
	}
}
