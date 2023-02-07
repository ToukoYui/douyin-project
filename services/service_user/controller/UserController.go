package controller

import (
	"douyin-template/model/pb"
	"douyin-template/services/service_user/dao"
	"douyin-template/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func Register(ctx *gin.Context) {
	request := pb.DouyinUserRegisterRequest{
		Username: ctx.Query("username"),
		Password: ctx.Query("password"),
	}
	if dao.IsUserExist(request.GetUsername()) {
		ctx.JSON(http.StatusOK, pb.DouyinUserRegisterResponse{
			StatusCode: 1,
			StatusMsg:  "注册失败，该用户名已存在",
		})
		return
	}

	// 将user添加进表中
	id := dao.InsertUser(&request)
	token := utils.CreateToken("my_key") // 生成token
	// TODO 需要在此访问线程中保存token和user_id

	ctx.JSON(http.StatusOK, pb.DouyinUserRegisterResponse{
		StatusCode: 0,
		StatusMsg:  "注册成功",
		UserId:     id,
		Token:      token,
	})
}

func Login(ctx *gin.Context) {
	request := pb.DouyinUserLoginRequest{
		Username: ctx.Query("username"),
		Password: ctx.Query("password"),
	}
	if !dao.IsUserExist(request.GetUsername()) {
		ctx.JSON(http.StatusBadRequest, pb.DouyinUserLoginResponse{
			StatusCode: 1,
			StatusMsg:  "该用户不存在",
		})
		return
	}
	if id := dao.VerifyUser(&request); id != 0 {
		//生成Token
		token := utils.CreateToken("my_key")
		// TODO 需要在此访问线程中保存token和user_id

		ctx.JSON(http.StatusOK, pb.DouyinUserLoginResponse{
			StatusCode: 0,
			StatusMsg:  "登录成功",
			UserId:     id,
			Token:      token,
		})
		return
	}
	ctx.JSON(http.StatusBadRequest, pb.DouyinUserLoginResponse{
		StatusCode: 1,
		StatusMsg:  "密码错误",
	})
}

func UserInfo(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Query("user_id"), 10, 64)
	if err != nil {
		fmt.Println(err)
		panic(any("id字符串转化失败！！！"))
	}
	request := pb.DouyinUserRequest{
		UserId: id,
		Token:  ctx.Query("token"),
	}
	// TODO 需要token校验

	// 获取用户信息
	userInfo := dao.GetUserInfo(&request)
	ctx.JSON(http.StatusOK, pb.DouyinUserResponse{
		StatusCode: 0,
		StatusMsg:  "获取用户信息成功",
		// 由于omitempty关键字无法识别并忽略嵌套结构体的字段空值，返回的json结果会包含时间的空值
		User: userInfo,
	})
}
