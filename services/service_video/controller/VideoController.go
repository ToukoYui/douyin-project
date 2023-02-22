package controller

import (
	"context"
	"douyin-template/model"
	"douyin-template/services/service_video/dao"
	"douyin-template/services/service_video/rpc"
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

// GetLikedVideo 获取用户点赞过的视频列表，提供给favorite服务调用
func (s *Server) GetLikedVideo(ctx context.Context, request *model.DouyinUseridAndVideoid) (*model.VideoDto, error) {
	fmt.Println()
	// 根据userid和videoid查询video
	video := dao.GetLikedVideo(request)
	userReq := model.DouyinUserRequest{
		UserId: video.GetUserId(),
	}
	fmt.Println("Video调用User。。。")
	response, err := rpc.VideoToUserRpcClient.GetUserInfo(context.Background(), &userReq)
	if err != nil {
		fmt.Sprintf("出错:%v", err)
	}

	return &model.VideoDto{
		Id:            video.GetId(),
		Author:        response.GetUser(),
		PlayUrl:       video.GetPlayUrl(),
		CoverUrl:      video.GetCoverUrl(),
		FavoriteCount: video.GetFavoriteCount(),
		CommentCount:  video.GetCommentCount(),
		IsFavorite:    video.GetIsFavorite(),
		Title:         video.GetTitle(),
	}, nil

}
