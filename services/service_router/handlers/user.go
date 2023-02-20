package handlers

import (
	"context"
	"douyin-template/model"
	"douyin-template/services/service_router/rpc"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

/**
接收并包装前端参数，通过rpc客户端发送给controller的rpc服务端
*/

func Login(ctx *gin.Context) {
	request := model.DouyinUserLoginRequest{
		Username: ctx.Query("username"),
		Password: ctx.Query("password"),
	}
	// 向user服务端发送请求并返回响应
	response, err := rpc.UserRpcClient.Login(context.Background(), &request)
	if err != nil {
		fmt.Println("出错", err)
	}
	//ctx.JSON(http.StatusOK, *response)
	ctx.JSON(200, *response)
}

func Register(ctx *gin.Context) {
	request := model.DouyinUserRegisterRequest{
		Username: ctx.Query("username"),
		Password: ctx.Query("password"),
	}
	response, err := rpc.UserRpcClient.Register(context.Background(), &request)
	if err != nil {
		fmt.Println("出错", err)
	}
	ctx.JSON(200, *response)
}

func UserInfo(ctx *gin.Context) {
	// 包装request
	id, err := strconv.ParseInt(ctx.Query("user_id"), 10, 64)
	if err != nil {
		fmt.Println(err)
		panic(any("id字符串转化失败！！！"))
	}
	request := model.DouyinUserRequest{
		UserId: id,
		Token:  ctx.Query("token"),
	}
	response, err := rpc.UserRpcClient.GetUserInfo(context.Background(), &request)
	if err != nil {
		fmt.Println("出错", err)
	}
	ctx.JSON(200, *response)
}
