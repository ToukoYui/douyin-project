package service

import (
	"douyin-template/model"
	"douyin-template/model/pb"
	"douyin-template/services/consts"
	"douyin-template/services/service_video/dao"
	"douyin-template/utils"
)

// UploadVideo 进行上传到oss和添加视频信息到video表
func UploadVideo(request *model.DouyinPublishActionRequest) error {
	data := request.GetData()
	// todo 获取用户id
	//dataArr := request.GetData()
	//request.GetToken()

	// 上传到oss
	playUrl, coverUrl, upLoadErr := utils.UploadVideo(data)
	if upLoadErr != nil {
		return upLoadErr
	}
	// 添加视频信息,改动:这里修改为了自己实现的基于redis的全局唯一性算法
	video := pb.Video{
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
