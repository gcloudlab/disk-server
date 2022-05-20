package test

import (
	"context"
	"gcloud/core/define"
	"testing"
	"time"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()
var rdb = redis.NewClient(&redis.Options{
	Addr:     "localhost:6379",
	Password: define.RedisPassword, // no password set
	DB:       0,                    // use default DB
})

func TestSetValue(t *testing.T) {
	err := rdb.Set(ctx, "key", 1121, time.Second*10).Err()
	if err != nil {
		t.Error(err)
	}
}

func TestGetValue(t *testing.T) {
	val, _ := rdb.Get(ctx, "key").Result()
	// if err != nil {
	// 	t.Error(err)
	// }
	t.Log(val)
}
