package handlers

import (
	"context"
	"douyin-template/model"
	"douyin-template/services/service_router/rpc"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

func CommentAction(ctx *gin.Context) {
	videoId, err := strconv.ParseInt(ctx.Query("video_id"), 10, 64)
	if err != nil {
		fmt.Sprintf("字符串id转int64失败：%v", err)
	}
	commentId, err := strconv.ParseInt(ctx.Query("comment_id"), 10, 64)
	if err != nil {
		fmt.Sprintf("字符串id转int64失败：%v", err)
	}
	actionType, err := strconv.ParseInt(ctx.Query("action_type"), 10, 32)
	if err != nil {
		fmt.Sprintf("actionType转int64失败：%v", err)
	}
	request := model.DouyinCommentActionRequest{
		Token:       ctx.Query("token"),
		VideoId:     videoId,
		ActionType:  int32(actionType),
		CommentText: ctx.Query("comment_text"),
		CommentId:   commentId,
	}
	response, err := rpc.CommentClient.CommentAction(context.Background(), &request)
	if err != nil {
		fmt.Println("出错", err)
	}
	ctx.JSON(200, *response)

	//listReq := model.DouyinCommentListRequest{
	//
	//}
	//commentList, err := rpc.CommentClient.CommentList(context.Background(), &listReq)
	//if err != nil {
	//	fmt.Println("出错", err)
	//}
}

func CommentList(ctx *gin.Context) {
	videoId, err := strconv.ParseInt(ctx.Query("video_id"), 10, 64)
	if err != nil {
		fmt.Sprintf("字符串id转int64失败：%v", err)
	}
	request := model.DouyinCommentListRequest{
		Token:   ctx.Query("token"),
		VideoId: videoId,
	}
	response, err := rpc.CommentClient.CommentList(context.Background(), &request)
	if err != nil {
		fmt.Println("出错", err)
	}
	ctx.JSON(200, *response)
}
