package rpc

import (
	"douyin-template/model"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

// FavoriteToVideoRpcClient 用于Favorite服务向video服务发送请求
var FavoriteToVideoRpcClient model.FeedSrvClient

//var FeedRpcClient model.FeedSrvClient

// InitFavoriteRpc 初始化Favorite客户端
func InitFavoriteRpc() {
	// 连接服务端，因为我们没有SSL证书，因此这里需要禁用安全传输
	dial, err := grpc.Dial("127.0.0.1:5679", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
		return
	}
	//defer dial.Close()
	FavoriteToVideoRpcClient = model.NewFeedSrvClient(dial)
}
