package controller

import (
	"context"
	"douyin-template/model"
	"douyin-template/services/service_comment/service"
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
		// 生成comment唯一id
		id := utils.NewSnowId()
		// 生成日期 MM-dd
		month := time.Now().Month().String()
		day := strconv.Itoa(time.Now().Day())
		creatData := strings.Join([]string{month, day}, "-")

		// todo redis持久化代替mysql，配置文件待设置redis数据库
		comment := model.Comment{
			Id:         id,
			VideoId:    request.GetVideoId(),
			Content:    request.GetCommentText(),
			CreateData: creatData,
		}
		//调用service层
		commentDto, _ := service.Send(comment, request.GetToken())

		//评论成功返回评论内容，不需要重新拉取整个列表
		return &model.DouyinCommentActionResponse{
			StatusCode: 0,
			StatusMsg:  "发布评论成功",
			Comment:    &commentDto,
		}, nil
	} else { // 2 删除评论
		err := service.DelComment(request.GetVideoId())
		if err != nil {
			fmt.Sprintf("删除评论失败：%v", err)
		}
		return &model.DouyinCommentActionResponse{
			StatusCode: 0,
			StatusMsg:  "删除评论成功",
			Comment:    nil,
		}, nil
	}
}

func (s *Server) CommentList(ctx context.Context, request *model.DouyinCommentListRequest) (*model.DouyinCommentListResponse, error) {
	// token验证
	if !utils.ValidateToken(request.GetToken()) {
		return &model.DouyinCommentListResponse{
			StatusCode:  1,
			StatusMsg:   "Token解析错误，验证失败",
			CommentList: nil,
		}, nil
	}
	//调用service层评论函数
	myClaim, _ := utils.ParseToken(request.GetToken())
	userId, err := strconv.ParseInt(myClaim.UserId, 10, 64)
	if err != nil {
		fmt.Sprintf("字符串id转int64失败：%v", err)
	}
	// 获取评论列表
	commentList, err := service.GetList(request.GetVideoId(), userId, request.GetToken())
	if err != nil {
		return &model.DouyinCommentListResponse{
			StatusCode:  1,
			StatusMsg:   "获取评论列表失败",
			CommentList: nil,
		}, nil
	}
	return &model.DouyinCommentListResponse{
		StatusCode:  0,
		StatusMsg:   "获取评论列表成功",
		CommentList: commentList,
	}, nil
}
