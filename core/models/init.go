package models

import (
	"gcloud/core/internal/config"
	"log"

	"github.com/go-redis/redis/v8"
	_ "github.com/go-sql-driver/mysql"
	"xorm.io/xorm"
)

var Engine = Init("root:root@tcp(127.0.0.1:3306)/gcloud?charset=utf8mb4&parseTime=True&loc=Local")

/*
	初始化数据库
*/
func Init(dataSource string) *xorm.Engine {
	engine, err := xorm.NewEngine("mysql", dataSource)
	if err != nil {
		log.Printf("Xorm New Engine Error:%v", err)
		return nil
	}
	return engine
}

/*
	初始化redis
*/
func InitRedis(c config.Config) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     c.Redis.Addr,
		Password: "", // no password set
		DB:       0,  // use default DB
	})
}
