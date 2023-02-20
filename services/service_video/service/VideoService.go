package service

import (
	"douyin-template/model/pb"
	"douyin-template/services/service_video/dao"
	"douyin-template/utils"
)

// UploadVideo 进行上传到oss和添加视频信息到video表
func UploadVideo(request *pb.DouyinPublishActionRequest) error {
	file := request.GetData()
	request.GetToken()
	// 上传到oss
	playUrl, coverUrl, upLoadErr := utils.UploadVideo(file)
	if upLoadErr != nil {
		return upLoadErr
	}
	// 添加视频信息
	video := pb.Video{
		Id:            utils.NewSnowId(),
		UserId:        446553213450061057,
		PlayUrl:       playUrl,
		CoverUrl:      coverUrl,
		FavoriteCount: 0,
		CommentCount:  0,
		//IsFavorite:    nil,
		Title: request.GetTitle(),
	}
	dao.CreateVideoInfo(&video)
	// 保存视频信息
	return nil
}
