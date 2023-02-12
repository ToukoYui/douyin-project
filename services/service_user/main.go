package main

import (
	"douyin-template/services/service_user/db"
	"douyin-template/utils"
	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()
	InitRouter(engine)
	Init()

	// 监听localhost:8080
	engine.Run()

}

// Init 初始化一些连接和配置
func Init() {
	utils.InitViper()                                           // 初始化viper
	db.InitDb(utils.Config.SshConfig, utils.Config.MysqlConfig) // mysql连接
	utils.InitRedis(utils.Config.RedisConfig)                   //Redis连接
	utils.NewSnowFlake()                                        //创建雪花算法初始配置，防止序号重复
}
