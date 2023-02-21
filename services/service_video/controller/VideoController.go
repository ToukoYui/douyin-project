package controller

import (
	"context"
	"douyin-template/model"
	"douyin-template/services/service_video/dao"
	"douyin-template/services/service_video/service"
	"fmt"
	"github.com/goccy/go-json"
)

// Feed 视频流接口

//func (s *Server) GetUserInfo(ctx context.Context, request *model.DouyinUserRequest) (*model.DouyinUserResponse, error) {
//
//}

type Server struct {
	// 继承 protoc-gen-go-grpc 生成的服务端代码
	model.UnimplementedFeedSrvServer
}

//func Feed(ctx *gin.Context) {
//	var latestTime int64
//	if timeStr := ctx.Query("latest_time"); timeStr == "" {
//		latestTime = time.Now().Unix() //可选参数，为空字符串则用当前时间戳代替
//	} else {
//		// 前端传过来的时间戳是1675778933948（55073-04-27 20:19:08），Year会超出设定的范围
//		// 报错为: “year is not in the range [1, 9999]: 55073”
//		// 这里如果越界就用now代替了
//		latestTime, _ = strconv.ParseInt(timeStr, 10, 64)
//		if latestTime >= 253380831548 { // (9999-04-27 20:19:08)
//			latestTime = time.Now().Unix()
//		}
//	}
//	request := pb.DouyinFeedRequest{
//		LatestTime: latestTime,
//		Token:      ctx.Query("token"),
//	}
//	videoInfoList, nextTime := dao.GetVideoInfoList(&request)
//	// todo redis缓存
//
//	// 控制台打印结果，可以不要
//	data, err := json.Marshal(*videoInfoList)
//	if err != nil {
//		panic(any("转化为json字符串失败"))
//	}
//	fmt.Println(string(data))
//
//	ctx.JSON(http.StatusOK, pb.DouyinFeedResponse{
//		StatusCode: 0,
//		StatusMsg:  "获取视频流成功",
//		VideoList:  videoInfoList,
//		NextTime:   nextTime,
//	})
//}

// Publish 视频投稿接口
//func Publish(ctx *gin.Context) {
//	data, _ := ctx.FormFile("data")
//
//	request := model.DouyinPublishActionRequest{
//		Token: ctx.PostForm("token"),
//		Data:  data,
//		Title: ctx.PostForm("title"),
//	}
//	err := service.UploadVideo(&request)
//	if err != nil {
//		ctx.JSON(http.StatusInternalServerError, pb.DouyinPublishActionResponse{
//			StatusCode: 1,
//			StatusMsg:  "发布视频失败",
//		})
//		return
//	}
//	ctx.JSON(http.StatusOK, pb.DouyinPublishActionResponse{
//		StatusCode: 0,
//		StatusMsg:  "发布视频成功",
//	})
//}

// PublishList 发布列表接口实现
func (s *Server) PublishList(ctx context.Context, request *model.DouyinPublishListRequest) (*model.DouyinPublishListResponse, error) {
	resultList := dao.GetPublishVideoList(request)
	return &model.DouyinPublishListResponse{
		StatusCode: 0,
		StatusMsg:  "获取用户发布视频列表成功",
		VideoList:  resultList,
	}, nil
}

// PublishAction 发布视频接口实现
func (s *Server) PublishAction(ctx context.Context, request *model.DouyinPublishActionRequest) (*model.DouyinPublishActionResponse, error) {

	err := service.UploadVideo(request)
	if err != nil {
		return &model.DouyinPublishActionResponse{
			StatusCode: 1,
			StatusMsg:  "发布视频失败",
		}, nil

	}
	return &model.DouyinPublishActionResponse{
		StatusCode: 0,
		StatusMsg:  "发布视频成功",
	}, nil
}

func (s *Server) GetUserFeed(ctx context.Context, request *model.DouyinFeedRequest) (*model.DouyinFeedResponse, error) {
	videoInfoList, nextTime := dao.GetVideoInfoList(request)
	// todo redis缓存

	// 控制台打印结果，可以不要
	data, err := json.Marshal(videoInfoList)
	if err != nil {
		panic(any("转化为json字符串失败"))
	}
	fmt.Println(string(data))

	return &model.DouyinFeedResponse{
		StatusCode: 0,
		StatusMsg:  "获取视频流成功",
		VideoList:  videoInfoList,
		NextTime:   nextTime,
	}, nil
}
