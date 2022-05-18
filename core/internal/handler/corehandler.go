package handler

// func CoreHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		var req types.Request
// 		if err := httpx.Parse(r, &req); err != nil {
// 			httpx.Error(w, err)
// 			return
// 		}

// 		l := logic.NewCoreLogic(r.Context(), svcCtx)
// 		resp, err := l.Core(&req)
// 		if err != nil {
// 			httpx.Error(w, err)
// 		} else {
// 			httpx.OkJson(w, resp)
// 		}
// 	}
// }
