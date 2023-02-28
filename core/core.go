package main

import (
	"flag"
	"fmt"
	"net/http"

	"gcloud/core/internal/config"
	"gcloud/core/internal/handler"
	"gcloud/core/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/core-api.yaml", "the config file")

func main() {
	flag.Parse()
	logx.Disable()

	// api config
	var c config.Config
	conf.MustLoad(*configFile, &c)

	ctx := svc.NewServiceContext(c) // 注入全局上下文
	server := rest.MustNewServer(c.RestConf, rest.WithCustomCors(nil, notAllowedFn))
	defer server.Stop()

	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}

func notAllowedFn(w http.ResponseWriter) {
	w.Header().Add("Access-Control-Allow-Methods", "GET, POST, DELETE, OPTIONS")
	w.Header().Add("Access-Control-Allow-Headers", "DNT,X-Mx-ReqToken,Keep-Alive,User-Agent,X-Requested-With,If-Modified-Since,Cache-Control,Content-Type,Authorization,token, access-control-allow-origin-type")
}
