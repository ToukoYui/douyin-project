package service

import (
	"bytes"
	"douyin-template/model"
	"douyin-template/services/consts"
	"douyin-template/services/service_video/dao"
	"douyin-template/utils"
)

// UploadVideo 进行上传到oss和添加视频信息到video表
func UploadVideo(request *model.DouyinPublishActionRequest) error {
	dataArr := request.GetData()
	reader := bytes.NewReader(dataArr)

	// 上传到oss
	playUrl, coverUrl, upLoadErr := utils.UploadVideo(reader, request.GetTitle())
	if upLoadErr != nil {
		return upLoadErr
	}
	// 添加视频信息,改动:这里修改为了自己实现的基于redis的全局唯一性算法
	video := model.Video{
		Id:            utils.NextId(consts.VIDEO_GET_KEY),
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
