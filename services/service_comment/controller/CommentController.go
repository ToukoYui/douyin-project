package controller

import (
	"context"
	"douyin-template/model"
	"douyin-template/services/service_comment/db"
	"douyin-template/services/service_comment/rpc"
	"douyin-template/utils"
	"fmt"
	"strconv"
	"strings"
	"time"
)

type Server struct {
	// 继承 protoc-gen-go-grpc 生成的服务端代码
	model.UnimplementedCommentSrvServer
}

func (s *Server) CommentAction(ctx context.Context, request *model.DouyinCommentActionRequest) (*model.DouyinCommentActionResponse, error) {
	// token验证
	if !utils.ValidateToken(request.GetToken()) {
		return &model.DouyinCommentActionResponse{
			StatusCode: 1,
			StatusMsg:  "Token解析错误，验证失败",
		}, nil
	}
	// 添加数据
	if action := request.GetActionType(); action == 1 { //1 发布评论
		// 调用user服务获取user对象
		myClaim, _ := utils.ParseToken(request.GetToken())
		userId, err := strconv.ParseInt(myClaim.UserId, 10, 64)
		if err != nil {
			fmt.Sprintf("字符串id转int64失败：%v", err)
		}
		response, err := rpc.CommentToUserRpcClient.GetUserInfo(context.Background(), &model.DouyinUserRequest{
			UserId: userId,
			Token:  request.GetToken(),
		})
		if err != nil {
			fmt.Sprintf("调用user服务失败：%v", err)
		}
		user := response.GetUser()

		// 生成comment唯一id
		id := utils.NewSnowId()
		// 生成日期 MM-dd
		month := time.Now().Month().String()
		day := strconv.Itoa(time.Now().Day())
		creatData := strings.Join([]string{month, day}, "-")

		// todo redis持久化代替mysql，配置文件待设置redis数据库
		db.Db.Create(&model.Comment{
			Id:         id,
			Content:    request.GetCommentText(),
			CreateDate: creatData,
		})

		//评论成功返回评论内容，不需要重新拉取整个列表
		comment := model.Comment{
			Id:         id,
			User:       user,
			Content:    request.GetCommentText(),
			CreateDate: creatData,
		}
		return &model.DouyinCommentActionResponse{
			StatusCode: 0,
			StatusMsg:  "发布评论成功",
			Comment:    &comment,
		}, nil
	} else { // 2 删除评论

		return &model.DouyinCommentActionResponse{
			StatusCode: 0,
			StatusMsg:  "删除评论成功",
			Comment:    nil,
		}, nil
	}

}

func (s *Server) FavoriteList(ctx context.Context, request *model.DouyinCommentListRequest) (*model.DouyinCommentListResponse, error) {
	// token验证
	if !utils.ValidateToken(request.GetToken()) {
		return &model.DouyinCommentListResponse{
			StatusCode:  1,
			StatusMsg:   "Token解析错误，验证失败",
			CommentList: nil,
		}, nil
	}

	// todo 1.根据video_id获取所有该视频下的所有评论对象

	// todo 2. 根据每个评论对象的user_id调用user服务获取user对象

	return &model.DouyinCommentListResponse{
		StatusCode:  1,
		StatusMsg:   "Token解析错误，验证失败",
		CommentList: nil,
	}, nil
}
