package dao

import (
	"douyin-template/model"
	"douyin-template/services/service_favorite/db"
	"douyin-template/utils"
	"fmt"
	"strconv"
)

// UpdateFavoriteType 更新点赞状态
func UpdateFavoriteType(request *model.DouyinFavoriteActionRequest) bool {
	// 解析token获取id
	myclaim, valid := utils.ParseToken(request.GetToken())
	if !valid {
		fmt.Sprintf("点赞模块，解析token失败")
	}
	userId, err := strconv.ParseInt(myclaim.UserId, 10, 64)
	if err != nil {
		fmt.Sprintf("字符串id转int64失败%v", err)
	}
	//favorite := model.Favorite{
	//	UserId:  userId,
	//	VideoId: request.GetVideoId(),
	//}
	// Model里用结构体提供查询条件的话 sql里会没有where语句，等待大伙解答
	rowsAffected := db.Db.Model(&model.Favorite{}).
		Where("user_id=? and video_id=? ", userId, request.GetVideoId()).
		Update("is_favorite", request.GetActionType()).RowsAffected
	//rowsAffected := db.Db.Model(&favorite).Update("is_favorite", request.GetActionType()).RowsAffected
	if rowsAffected <= 0 {
		fmt.Sprintf("点赞更新操作失败")
		return false
	}
	return true
}

// GetUserIdAndVideoId 查询被点赞过的视频id和用户id
func GetUserIdAndVideoId(request *model.DouyinFavoriteListRequest) []model.Favorite {
	var likedVideoIdSli []model.Favorite
	rowsAffected := db.Db.Select("user_id", "video_id").Where("user_id =? and is_favorite=?", request.GetUserId(), 1).Find(&likedVideoIdSli).RowsAffected
	if rowsAffected <= 0 {
		fmt.Println("查询点赞视频id错误")
	}
	return likedVideoIdSli
}
