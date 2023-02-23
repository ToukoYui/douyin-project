package rpc

import (
	"douyin-template/model"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

// CommentToUserRpcClient 用于comment服务向user服务发送请求
var CommentToUserRpcClient model.UserSrvClient

var CommentTOVideoRpcClient model.FeedSrvClient

// InitCommentRpc 初始化User客户端
func InitCommentRpc() {
	// 连接服务端，因为我们没有SSL证书，因此这里需要禁用安全传输
	dial, err := grpc.Dial("127.0.0.1:5678", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
		return
	}
	//defer dial.Close()
	CommentToUserRpcClient = model.NewUserSrvClient(dial)
	CommentTOVideoRpcClient = model.NewFeedSrvClient(dial)
}
