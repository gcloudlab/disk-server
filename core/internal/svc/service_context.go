package svc

import (
	"gcloud/core/internal/config"
	"gcloud/core/internal/middleware"
	"gcloud/core/models"

	"github.com/go-redis/redis/v8"
	"github.com/zeromicro/go-zero/rest"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config config.Config   // 配置 (core-api.yaml)
	Engine *gorm.DB        // orm
	RDB    *redis.Client   // Redis
	Auth   rest.Middleware // auth
}

// 上下文
func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		Engine: models.Init(c.Mysql.DataSource),
		RDB:    models.InitRedis(c),
		Auth:   middleware.NewAuthMiddleware().Handle,
	}
}
