package controller

import (
	"context"
	"douyin-template/model"
	"douyin-template/services/service_favorite/dao"
	"douyin-template/services/service_favorite/rpc"
	"douyin-template/utils"
	"fmt"
)

type Server struct {
	model.UnimplementedFavoriteSrvServer
}

// FavoriteAction 点赞操作
func (i *Server) FavoriteAction(ctx context.Context, request *model.DouyinFavoriteActionRequest) (*model.DouyinFavoriteActionResponse, error) {
	if isUpdate := dao.UpdateFavoriteType(request); !isUpdate {
		return &model.DouyinFavoriteActionResponse{
			StatusCode: 1,
			StatusMsg:  "点赞操作失败",
		}, nil
	}
	return &model.DouyinFavoriteActionResponse{
		StatusCode: 0,
		StatusMsg:  "点赞操作成功",
	}, nil

}

// FavoriteList 点赞列表
func (i *Server) FavoriteList(ctx context.Context, request *model.DouyinFavoriteListRequest) (*model.DouyinFavoriteListResponse, error) {
	// token验证
	if !utils.ValidateToken(request.GetToken()) {
		return &model.DouyinFavoriteListResponse{
			StatusCode: 1,
			StatusMsg:  "Token解析错误，验证失败",
		}, nil
	}
	// 获取点赞视频id和用户id
	userIdAndVideoIdSli := dao.GetUserIdAndVideoId(request)

	// 调用video服务，获取video列表
	videoDatoList := make([]*model.VideoDto, len(userIdAndVideoIdSli))
	for i, ids := range userIdAndVideoIdSli {
		item := model.DouyinUseridAndVideoid{
			UserId:  ids.GetUserId(),
			VideoId: ids.GetVideoId(),
		}
		fmt.Println(item)
		// todo 出错
		videoDto, err := rpc.FavoriteToVideoRpcClient.GetLikedVideo(context.Background(), &item)
		if err != nil {
			fmt.Sprintf("favorite服务调用Video服务失败:%v", err)
		}
		//fmt.Println("dawafwafawfawfa", res)
		videoDatoList[i] = videoDto
	}
	return &model.DouyinFavoriteListResponse{
		StatusCode: 0,
		StatusMsg:  "获取点赞视频列表成功",
		VideoList:  videoDatoList,
	}, nil

}
