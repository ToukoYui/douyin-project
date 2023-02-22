package rpc

import (
	"douyin-template/model"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

var CommentClient model.CommentSrvClient

func initCommentRpc() {
	dial, err := grpc.Dial("127.0.0.1:5681", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
		return
	}
	CommentClient = model.NewCommentSrvClient(dial)
}
