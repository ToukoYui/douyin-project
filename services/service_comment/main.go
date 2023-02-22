package main

import (
	"douyin-template/model"
	"douyin-template/services/service_comment/controller"
	"douyin-template/services/service_comment/db"
	"douyin-template/services/service_comment/rpc"
	"douyin-template/utils"
	"google.golang.org/grpc"
	"log"
	"net"
)

/**
初始化comment服务的服务端
*/

// Init 初始化一些连接和配置
func Init() {
	rpc.InitCommentRpc()
	utils.InitViper("./services/service_comment/application.yml") // 初始化viper
	db.InitDb(utils.Config.SshConfig, utils.Config.MysqlConfig)   // mysql连接
	utils.InitRedis(utils.Config.RedisConfig)                     //Redis连接
	utils.NewSnowFlake()                                          //创建雪花算法初始配置，防止序号重复
}

func main() {
	Init()
	// 监听本地 5678 端口
	listen, err := net.Listen("tcp", ":5681")
	if err != nil {
		log.Fatal(err)
		return
	}

	// 创建 gRPC 服务器
	s := grpc.NewServer()
	// 将实现的接口注册进 gRPC 服务器
	model.RegisterCommentSrvServer(s, &controller.Server{})
	log.Println("comment server starts running...")
	// 启动 gRPC 服务器
	err = s.Serve(listen)
	if err != nil {
		log.Fatal(err)
		return
	}
}
