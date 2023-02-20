package main

import (
	"douyin-template/services/service_router/handlers"
	"douyin-template/services/service_router/rpc"
	"douyin-template/services/service_user/db"
	"douyin-template/utils"
	"github.com/gin-gonic/gin"
)

// Init 初始化一些连接和配置
func Init() {
	utils.InitViper("./services/service_user/application.yml")  // 初始化viper
	rpc.InitRPC()                                               //初始化RPC客户端
	db.InitDb(utils.Config.SshConfig, utils.Config.MysqlConfig) // mysql连接
	utils.InitRedis(utils.Config.RedisConfig)                   //Redis连接
	utils.NewSnowFlake()                                        //创建雪花算法初始配置，防止序号重复
}

func InitRouter(engine *gin.Engine) {
	engine.Static("/static", "./public")
	preGroup := engine.Group("/douyin")

	// user apis
	preGroup.GET("/user/", handlers.UserInfo)
	preGroup.POST("/user/register/", handlers.Register)
	preGroup.POST("/user/login/", handlers.Login)

	//video apis
	//preGroup.GET("/feed/", handlers.Feed)
	//preGroup.POST("/publish/action/", handlers.Publish)
	//preGroup.GET("/publish/list/", handlers.PublishList)
}

func main() {
	engine := gin.Default()
	InitRouter(engine)
	Init()

	// 监听localhost:8080
	engine.Run()

}
