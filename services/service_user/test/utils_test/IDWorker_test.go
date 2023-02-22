package utils

import (
	"douyin-template/services/consts"
	"douyin-template/utils"
	"fmt"
	"testing"
)

/**
* 测试全局性唯一ID
*
* @author: 张庭杰
* @date: 2023年02月21日 20:19
 */

func TestNextId(t *testing.T) {
	config := utils.Config.RedisConfig
	config.Addr = "8.134.208.93"
	config.Port = "6379"
	config.Db = "1"
	utils.InitRedis(config)
	fmt.Println(utils.NextId(consts.VIDEO_GET_KEY))
	fmt.Println(utils.NextId(consts.VIDEO_GET_KEY))
	fmt.Println(utils.NextId(consts.VIDEO_GET_KEY))
	fmt.Println(utils.NextId(consts.VIDEO_GET_KEY))
	fmt.Println(utils.NextId(consts.VIDEO_GET_KEY))
}
