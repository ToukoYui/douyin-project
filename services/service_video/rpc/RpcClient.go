package rpc

import (
	"douyin-template/model"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

// VideoToUserRpcClient 用于video服务向user服务发送请求
var VideoToUserRpcClient model.UserSrvClient

var VideoToFavoriteRpcClient model.FavoriteSrvClient

// InitUserRpc 初始化User客户端
func InitUserRpc() {
	// 连接服务端，因为我们没有SSL证书，因此这里需要禁用安全传输
	size := 1024 * 1024 * 20
	dial, err := grpc.Dial("127.0.0.1:5678", grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(size)))
	//dial, err := grpc.Dial("127.0.0.1:5678", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
		return
	}

	dial2, err := grpc.Dial("127.0.0.1:5680", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
		return
	}
	VideoToUserRpcClient = model.NewUserSrvClient(dial)
	VideoToFavoriteRpcClient = model.NewFavoriteSrvClient(dial2)
}
