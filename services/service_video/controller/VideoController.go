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

// GetUserFeed 获取视频流接口实现
func (s *Server) GetUserFeed(ctx context.Context, request *model.DouyinFeedRequest) (*model.DouyinFeedResponse, error) {
	//todoHandler:张庭杰 处理时间:2023年2月21日21:21:51
	//1. 组装redis中的key

	videoInfoList, nextTime := dao.GetVideoInfoList(request)
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
