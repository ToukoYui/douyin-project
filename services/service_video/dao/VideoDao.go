package dao

import (
	"context"
	"douyin-template/model"
	"douyin-template/model/pb"
	"douyin-template/services/service_user/db"
	"douyin-template/services/service_video/rpc"
	"fmt"
	"time"
)

// GetVideoInfoList 获取视频信息
func GetVideoInfoList(request *model.DouyinFeedRequest) ([]*model.VideoDto, int64) {
	// 根据LatestTime降序查询视频列表
	var videoList []model.Video
	latestTime := time.Unix(request.GetLatestTime(), 0)
	fmt.Println("请求时间为：", latestTime)
	row := db.Db.Where("created_time>=?", latestTime).Order("created_time desc").Find(&videoList).RowsAffected
	if row <= 0 {
		fmt.Sprintf("无查询结果,批量查询失败,查询时间参数：%v", latestTime)
	}

	fmt.Printf("videoList结果大小：%d,结果列表:%#v\n", len(videoList), videoList)

	resultList := make([]*model.VideoDto, len(videoList))
	for i, item := range videoList {
		//根据user_id查询User对象
		var user model.User
		db.Db.Select([]string{"id", "name", "follow_count", "follower_count"}).First(&user, item.GetUserId())
		// 查询favorite表获取is_favorite
		favorite := pb.Favorite{}
		db.Db.Select("is_favorite").Where("user_id=? and video_id=?", user.GetId(), item.GetId()).First(&favorite)
		// 包装视频列表结果
		resultList[i] = &model.VideoDto{
			Id:            item.GetId(),
			Author:        &user,
			PlayUrl:       item.GetPlayUrl(),
			CoverUrl:      item.GetCoverUrl(),
			FavoriteCount: item.GetFavoriteCount(),
			CommentCount:  item.GetCommentCount(),
			IsFavorite:    favorite.IsFavorite,
			Title:         item.GetTitle(),
		}
	}
	fmt.Printf("结果大小：%d,结果列表:%#v\n", len(resultList), resultList)
	// 获取视频列表中的最早发布的时间
	earlyTime := videoList[len(videoList)-1].CreatedAt
	return resultList, earlyTime.Unix()
}

// CreateVideoInfo 添加视频信息
func CreateVideoInfo(video *model.Video) {
	row := db.Db.Create(video).RowsAffected
	if row != 1 {
		panic(any("添加视频信息失败！！！"))
	}
}

// GetPublishVideoList 获取单个用户的发布视频列表
func GetPublishVideoList(request *model.DouyinPublishListRequest) []*model.VideoDto {
	// 查询该用户发布过的视频信息
	var videoList []model.Video
	db.Db.Where("user_id=?", request.UserId).Find(&videoList)
	// 遍历包装
	resultList := make([]*model.VideoDto, len(videoList))
	fmt.Println("调用User服务查询对象")
	for i, item := range videoList {
		//调用user服务查询对象，根据user_id查询User对象，此处可以通过批量处理优化
		userInfoResp, err := rpc.VideoToUserRpcClient.GetUserInfo(context.Background(), &model.DouyinUserRequest{
			UserId: item.GetUserId(),
			Token:  request.GetToken(),
		})
		if err != nil {
			fmt.Sprintf("调用User服务查询对象失败%v", err)
		}
		user := *userInfoResp.GetUser()

		// 查询favorite表获取is_favorite   todo 调用like服务查询is_favorite
		favorite := pb.Favorite{}
		db.Db.Select("is_favorite").Where("user_id=? and video_id=?", user.GetId(), item.GetId()).First(&favorite)
		// 包装视频列表结果
		resultList[i] = &model.VideoDto{
			Id:            item.GetId(),
			Author:        &user,
			PlayUrl:       item.GetPlayUrl(),
			CoverUrl:      item.GetCoverUrl(),
			FavoriteCount: item.GetFavoriteCount(),
			CommentCount:  item.GetCommentCount(),
			IsFavorite:    favorite.IsFavorite,
			Title:         item.GetTitle(),
		}
	}
	return resultList
}

// GetLikedVideo 根据userid和videoid查询video
func GetLikedVideo(request *model.DouyinUseridAndVideoid) *model.Video {
	video := model.Video{}
	db.Db.Where("user_id=? and id=?", request.GetUserId(), request.GetVideoId()).First(&video)
	return &video
}
