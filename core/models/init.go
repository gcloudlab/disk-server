package models

import (
	"gcloud/core/define"
	"log"

	"github.com/go-redis/redis/v8"
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Engine = Init("root:root@tcp(127.0.0.1:3306)/gcloud?charset=utf8mb4&parseTime=True&loc=Local")
var RDB = InitRedis()

/*
	初始化数据库
*/
func Init(dataSource string) *gorm.DB {
	// engine, err := xorm.NewEngine("mysql", dataSource)
	engine, err := gorm.Open(mysql.Open(dataSource), &gorm.Config{})
	if err != nil {
		log.Printf("Xorm New Engine Error:%v", err)
		return nil
	}

	return engine
}

/*
	初始化redis
*/
func InitRedis() *redis.Client {
	return redis.NewClient(&redis.Options{
		// Addr:     c.Redis.Addr,
		Addr:     "localhost:6379",
		Password: define.RedisPassword, // no password set
		DB:       0,                    // use default DB
	})
}
