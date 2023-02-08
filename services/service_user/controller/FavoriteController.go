package controller

import (
	"douyin-template/model/pb"
	"github.com/gin-gonic/gin"
	common "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

// FavoriteListResponse 点赞列表返回值
type FavoriteListResponse struct {
	common.Response
	VideoList []pb.Video `json:"video_list"`
}

// FavoriteActionRequest 点赞与取消请求
type FavoriteActionRequest struct {
	Token      string `from:"token" validate:"required,jwt"`
	VideoID    int64  `from:"video_id" validate:"required,numeric,min=1"`
	ActionType int32  `from:"action_type" validate:"required,numeric,one of=1 2"`
}

// FavoriteListRequest 点赞列表请求
type FavoriteListRequest struct {
	UserId int64  `form:"user_id" validate:"required,numeric,min=1"`
	Token  string `form:"token"   validate:"required,jwt"`
}

// FavoriteAction 点赞操作
func FavoriteAction(ctx *gin.Context) {

}

// FavoriteList 点赞列表
func FavoriteList(ctx *gin.Context) {

}
