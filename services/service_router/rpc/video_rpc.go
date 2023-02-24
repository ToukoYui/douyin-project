package rpc

import (
	"douyin-template/model"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

var FeedRpcClient model.FeedSrvClient

// initFeedRpc 初始化Video客户端
func initFeedRpc() {
	size := 1024 * 1024 * 20

	//dial, err := grpc.Dial("127.0.0.1:5679", grpc.WithTransportCredentials(insecure.NewCredentials()))
	dial, err := grpc.Dial("127.0.0.1:5679", grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(size)))
	if err != nil {
		log.Fatal(err)
		return
	}
	FeedRpcClient = model.NewFeedSrvClient(dial)
}
