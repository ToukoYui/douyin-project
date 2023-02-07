package dao

import (
	"douyin-template/model/pb"
	"douyin-template/services/service_user/db"
	"fmt"
	"time"
)

// GetVideoInfoList 获取视频信息
func GetVideoInfoList(request *pb.DouyinFeedRequest) (*[]pb.VideoDto, int64) {
	// 根据LatestTime降序查询视频列表
	var videoList []pb.Video
	latestTime := time.Unix(request.GetLatestTime(), 0)
	row := db.Db.Where("created_time<?", latestTime).Find(&videoList).RowsAffected
	if row <= 0 {
		panic(any("批量查询失败！！！"))
	}
	resultList := make([]pb.VideoDto, len(videoList))
	for i, item := range videoList {
		//根据user_id查询User对象
		var user pb.User
		db.Db.First(&user, item.GetUserId())
		// 包装视频列表结果
		resultList[i] = pb.VideoDto{
			Id:            item.GetId(),
			Author:        user,
			PlayUrl:       item.GetPlayUrl(),
			CoverUrl:      item.GetCoverUrl(),
			FavoriteCount: item.GetFavoriteCount(),
			CommentCount:  item.GetCommentCount(),
			IsFavorite:    item.GetIsFavorite(),
			Title:         item.GetTitle(),
		}
	}
	fmt.Printf("结果列表:%#v\n", resultList[0])
	return &resultList, request.GetLatestTime()
}

// CreateVideoInfo 添加视频信息
func CreateVideoInfo(video *pb.Video) {
	row := db.Db.Create(video).RowsAffected
	if row != 1 {
		panic(any("添加视频信息失败！！！"))
	}
}
