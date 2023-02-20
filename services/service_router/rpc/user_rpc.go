package rpc

import (
	"douyin-template/model"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

/**
创建各个模块的RPC服务端
*/

var UserRpcClient model.UserSrvClient

// InitUserRpc 初始化User客户端
func InitUserRpc() {
	// 连接服务端，因为我们没有SSL证书，因此这里需要禁用安全传输
	dial, err := grpc.Dial("127.0.0.1:5678", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
		return
	}
	//defer dial.Close()
	UserRpcClient = model.NewUserSrvClient(dial)

}
