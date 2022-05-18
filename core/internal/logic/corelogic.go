package logic

import (
	"context"

	"gcloud/core/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type CoreLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCoreLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CoreLogic {
	return &CoreLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// func (l *CoreLogic) Core(req *types.Request) (resp *types.Response, err error) {
// 	// 获取用户列表
// 	data := make([]*models.UserBasic, 0)
// 	err = models.Engine.Find(&data)
// 	if err != nil {
// 		log.Println("Get User Basic Error:", err)
// 	}

// 	b, err := json.Marshal(data)
// 	if err != nil {
// 		log.Println("Marshal Error:", err)
// 	}

// 	dst := new(bytes.Buffer)
// 	err = json.Indent(dst, b, "", "  ")
// 	if err != nil {
// 		log.Println("Indent Error:", err)
// 	}
// 	// fmt.Println(dst.String())

// 	sql := "select * from user_basic"
// 	results, _ := models.Engine.QueryString(sql)
// 	log.Println("数据", results)

// 	resp = new(types.Response)
// 	resp.Message = dst.String()
// 	return
// }
