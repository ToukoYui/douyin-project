package handlers

import (
	"context"
	"douyin-template/model"
	"douyin-template/services/service_router/rpc"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

func FavoriteAction(ctx *gin.Context) {
	videoId, err := strconv.ParseInt(ctx.Query("video_id"), 10, 64)
	if err != nil {
		fmt.Sprintf("字符串id转int64失败：%v", err)
	}
	actionType, err := strconv.ParseInt(ctx.Query("action_type"), 10, 32)
	if err != nil {
		fmt.Sprintf("actionType转int64失败：%v", err)
	}
	request := model.DouyinFavoriteActionRequest{
		Token:      ctx.Query("token"),
		VideoId:    videoId,
		ActionType: int32(actionType),
	}
	response, err := rpc.FavoriteRpcClient.FavoriteAction(context.Background(), &request)
	if err != nil {
		fmt.Println("出错", err)
	}
	ctx.JSON(200, *response)
}

func FavoriteList(ctx *gin.Context) {
	userId, err := strconv.ParseInt(ctx.Query("user_id"), 10, 64)
	if err != nil {
		fmt.Sprintf("字符串id转int64失败：%v", err)
	}
	request := model.DouyinFavoriteListRequest{
		UserId: userId,
		Token:  ctx.Query("token"),
	}
	response, err := rpc.FavoriteRpcClient.FavoriteList(context.Background(), &request)
	if err != nil {
		fmt.Println("出错", err)
	}
	ctx.JSON(200, *response)
}
