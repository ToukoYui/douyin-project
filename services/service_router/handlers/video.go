package handlers

import (
	"context"
	"douyin-template/model"
	"douyin-template/services/service_router/rpc"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"strconv"
	"time"
)

func PublishList(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Query("user_id"), 10, 64)
	if err != nil {
		panic(any(fmt.Sprintf("字符串转int64错误:%v", err)))
	}
	request := model.DouyinPublishListRequest{
		UserId: id,
		Token:  ctx.Query("token"),
	}
	response, err := rpc.FeedRpcClient.PublishList(context.Background(), &request)
	fmt.Println("响应", *response)
	if err != nil {
		fmt.Println("出错", err)
	}
	ctx.JSON(200, *response)
}

func Publish(ctx *gin.Context) {
	fileHeader, _ := ctx.FormFile("data")
	open, err2 := fileHeader.Open()
	if err2 != nil {
		panic(any("打开视频流失败"))
	}

	//var dataArr []byte
	fileByte, err := ioutil.ReadAll(open)

	//open.Read(dataArr)

	request := model.DouyinPublishActionRequest{
		Token: ctx.PostForm("token"),
		Data:  fileByte,
		Title: ctx.PostForm("title"),
	}
	response, err := rpc.FeedRpcClient.PublishAction(context.Background(), &request)
	if err != nil {
		fmt.Println("出错", err)
	}
	ctx.JSON(200, *response)
}

func Feed(ctx *gin.Context) {
	var latestTime int64
	if timeStr := ctx.Query("latest_time"); timeStr == "" {
		latestTime = time.Now().Unix() //可选参数，为空字符串则用当前时间戳代替
	} else {
		// 前端传过来的时间戳是1675778933948（55073-04-27 20:19:08），Year会超出设定的范围
		// 报错为: “year is not in the range [1, 9999]: 55073”
		// 这里如果越界就用now代替了
		latestTime, _ = strconv.ParseInt(timeStr, 10, 64)
		if latestTime >= 253380831548 { // (9999-04-27 20:19:08)
			latestTime = time.Now().Unix()
		}
	}
	request := model.DouyinFeedRequest{
		LatestTime: latestTime,
		Token:      ctx.Query("token"),
	}
	response, err := rpc.FeedRpcClient.GetUserFeed(context.Background(), &request)
	if err != nil {
		fmt.Println("出错", err)
	}
	ctx.JSON(200, *response)
}
