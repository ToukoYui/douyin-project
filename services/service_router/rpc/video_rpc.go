package rpc

import (
	"douyin-template/model"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

// initFeedRpc 初始化Video客户端
func initFeedRpc() {
	dial, err := grpc.Dial("127.0.0.1:5679", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
		return
	}
	FeedRpcClient = model.NewFeedSrvClient(dial)
}
