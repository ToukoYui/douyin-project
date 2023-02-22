package main

import (
	"douyin-template/model"
	"douyin-template/services/service_favorite/controller"
	"douyin-template/services/service_favorite/db"
	"douyin-template/services/service_favorite/rpc"
	"douyin-template/utils"
	"google.golang.org/grpc"
	"log"
	"net"
)

// Init 初始化一些连接和配置
func Init() {
	rpc.InitFavoriteRpc()
	utils.InitViper("./services/service_favorite/application.yml") // 初始化viper
	db.InitDb(utils.Config.SshConfig, utils.Config.MysqlConfig)    // mysql连接
	utils.InitRedis(utils.Config.RedisConfig)                      //Redis连接
	utils.NewSnowFlake()                                           //创建雪花算法初始配置，防止序号重复
}

func main() {
	Init()
	// 监听本地 5680 端口
	listen, err := net.Listen("tcp", ":5680")
	if err != nil {
		log.Fatal(err)
		return
	}

	// 创建 gRPC 服务器
	s := grpc.NewServer()
	// 将实现的接口注册进 gRPC 服务器
	model.RegisterFavoriteSrvServer(s, &controller.Server{})
	log.Println("favorite server starts running...")
	// 启动 gRPC 服务器
	err = s.Serve(listen)
	if err != nil {
		log.Fatal(err)
		return
	}
}
