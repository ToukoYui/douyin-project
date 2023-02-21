package utils

import (
	"context"
	"fmt"
	"time"
)

const LAYOUT string = "2006-01-02"
const SeqBits int64 = 32

var ctx = context.Background()

/**
* 全局唯一性ID生成器
*
* @author: 张庭杰
* @date: 2023年02月21日 19:59
 */

func NextId(serviceKey string) int64 {
	//1. 获取当前时间戳
	timeStamp := time.Now().Unix()
	//2.获取当前的日期
	date := time.Now().Format(LAYOUT)
	//3. 从redis中取序号
	//3.1 组装key
	key := "icr:" + serviceKey + date
	//3.2 连接redis并且取数据
	seq, err := RedisDb.Get(ctx, key).Int()
	//3.3 在取完数据之后,自增,避免重复ID
	RedisDb.Incr(ctx, key)
	if err != nil {
		fmt.Printf("获取redis数据失败!---全局唯一性ID,%s", err)
		return 0
	}
	fmt.Printf("从redis中获取到的key为:%d", seq)
	if err != nil {
		return 0
	}
	//4. 拼接ID并且返回
	return (timeStamp << SeqBits) | int64(seq)
}
