package utils

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v8"
	"testing"
	"time"
)

func TestRedis(t *testing.T) {
	client := initRedis()
	hasMap := make(map[string]string)
	hasMap["key5"] = "value5"
	hasMap["key6"] = "value6"
	marshal, err := json.Marshal(hasMap)
	fmt.Println(marshal)
	if err != nil {
		panic(any(err))
	}
	client.Set(context.Background(), "hash", marshal, time.Minute*10)

	//client.Set(context.Background(), "hash", "hasMap", time.Minute*10)
	result, _ := client.Get(context.Background(), "hash").Result()

	fmt.Println(result)
}

// 定义一个全局变量
var redisdb *redis.Client

func initRedis() *redis.Client {
	redisdb = redis.NewClient(&redis.Options{
		Addr:     "8.134.208.93:6379", // 指定
		Password: "",
		DB:       0, // redis一共16个库，指定其中一个库即可
	})

	return redisdb
}
