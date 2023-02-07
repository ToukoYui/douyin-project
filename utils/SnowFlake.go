package utils

import "github.com/sony/sonyflake"

var SnowFlake *sonyflake.Sonyflake

func NewSnowFlake() {
	SnowFlake = sonyflake.NewSonyflake(sonyflake.Settings{})
}

// NewSnowId 通过雪花算法生成唯一id
func NewSnowId() int64 {
	id, err := SnowFlake.NextID()
	if err != nil {
		panic(any("生成唯一id失败！！！"))
	}
	return int64(id)
}
