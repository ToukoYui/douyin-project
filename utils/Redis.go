package utils

import (
	"fmt"
	"github.com/go-redis/redis/v8"
	"strconv"
)

var RedisDb *redis.Client

func InitRedis(config Redis) {
	fmt.Println("Redis已经连接...")
	db, _ := strconv.Atoi(config.Db)
	//fmt.Println("redis", Config.RedisConfig.Addr)
	RedisDb = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", config.Addr, config.Port), // 指定
		Password: "",
		DB:       db, // redis一共16个库，指定其中一个库即可
	})
}
