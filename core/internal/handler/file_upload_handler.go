package handler

import (
	"crypto/md5"
	"fmt"
	"gcloud/core/models"
	"net/http"
	"path"

	"gcloud/core/helper"
	"gcloud/core/internal/logic"
	"gcloud/core/internal/svc"
	"gcloud/core/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func FileUploadHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.FileUploadRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		file, fileHeader, err := r.FormFile("file")
		if err != nil {
			return
		}

		// 生成文件hash, 判断文件是否已存在
		bt := make([]byte, fileHeader.Size)
		_, err = file.Read(bt)
		if err != nil {
			return
		}
		hash := fmt.Sprintf("%x", md5.Sum(bt))

		rp := new(models.RepositoryPool)
		svcCtx.Engine.
			Where("hash = ?", hash).
			First(rp)

		// 文件已存在
		if rp.Id != 0 {
			httpx.OkJson(w, &types.FileUploadReply{
				Identity: rp.Identity,
				Ext:      rp.Ext,
				Name:     rp.Name,
				Msg:      "success",
			})
			return
		}

		// 文件不存在，上传文件到COS
		cosPath, err := helper.CosUpload(r)
		if err != nil {
			return
		}

		// to logic
		req.Name = fileHeader.Filename
		req.Hash = hash
		req.Path = cosPath
		req.Ext = path.Ext(fileHeader.Filename)
		req.Size = fileHeader.Size

		l := logic.NewFileUploadLogic(r.Context(), svcCtx)
		resp, err := l.FileUpload(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
