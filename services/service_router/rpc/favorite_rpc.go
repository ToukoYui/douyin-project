package rpc

import (
	"douyin-template/model"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

var FavoriteRpcClient model.FavoriteSrvClient

// func initFavoriteRpc 初始化Favorite客户端
func initFavoriteRpc() {
	dial, err := grpc.Dial("127.0.0.1:5680", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
		return
	}
	FavoriteRpcClient = model.NewFavoriteSrvClient(dial)
}
