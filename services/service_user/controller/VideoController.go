package controller

import (
	"douyin-template/model/pb"
	"douyin-template/services/service_user/dao"
	"douyin-template/services/service_user/service"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

func Feed(ctx *gin.Context) {
	fmt.Println(ctx.Query("latest_time"))
	var latestTime int64
	if timeStr := ctx.Query("latest_time"); timeStr == "" {
		latestTime = time.Now().Unix() //可选参数，为空字符串则用当前时间戳代替
	} else {
		// 前端传过来的时间戳是1675778933948（55073-04-27 20:19:08），Year会超出gorm设定的范围
		// 报错为: “year is not in the range [1, 9999]: 55073”
		latestTime, _ = strconv.ParseInt(timeStr, 10, 64)
		if latestTime >= 253380831548 { // (9999-04-27 20:19:08)
			latestTime = time.Now().Unix() // 如果越界就用now代替了
		}
	}
	request := pb.DouyinFeedRequest{
		LatestTime: latestTime,
		Token:      ctx.Query("token"),
	}
	// todo
	videoInfoList, nextTime := dao.GetVideoInfoList(&request)
	data, err := json.Marshal(*videoInfoList)
	if err != nil {
		panic(any("转化为json字符串失败"))
	}
	fmt.Println(string(data))
	ctx.JSON(http.StatusOK, pb.DouyinFeedResponse{
		StatusCode: 0,
		StatusMsg:  "获取视频流成功",
		VideoList:  videoInfoList,
		NextTime:   nextTime,
	})
}

func Publish(ctx *gin.Context) {
	data, _ := ctx.FormFile("data")

	//dataArr := []byte{}
	//open.Read(dataArr)
	//defer open.Close()

	request := pb.DouyinPublishActionRequest{
		Token: ctx.PostForm("token"),
		Data:  data,
		Title: ctx.PostForm("title"),
	}
	err := service.UploadVideo(&request)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, pb.DouyinPublishActionResponse{
			StatusCode: 1,
			StatusMsg:  "发布视频失败",
		})
		return
	}
	ctx.JSON(http.StatusOK, pb.DouyinPublishActionResponse{
		StatusCode: 0,
		StatusMsg:  "发布视频成功",
	})
}
