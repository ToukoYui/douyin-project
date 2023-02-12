package controller

import (
	"context"
	"douyin-template/model/pb"
	"douyin-template/services/service_user/dao"
	"douyin-template/utils"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type TokenInfo struct {
	UserId   string `json:"user_id"`
	UserName string `json:"username"`
}

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
	id, name := dao.InsertUser(&request)
	token := utils.CreateToken(id, name) // 生成token

	//  将id和name放进redis，用于token检验时匹配
	tokenInfo := TokenInfo{
		UserId:   strconv.FormatInt(id, 10),
		UserName: name,
	}
	marshal, err := json.Marshal(tokenInfo)
	if err != nil {
		panic(any("json转化失败"))
	}
	utils.RedisDb.Set(context.Background(), strings.Join([]string{"token", token}, ":"), marshal, time.Minute*30)
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
	// 登录成功
	if id, name := dao.VerifyUser(&request); id != 0 {
		//生成Token
		token := utils.CreateToken(id, name)
		//  将id和name放进redis，用于token检验时匹配
		tokenInfo := TokenInfo{
			UserId:   strconv.FormatInt(id, 10),
			UserName: name,
		}
		marshal, err := json.Marshal(tokenInfo)
		if err != nil {
			panic(any("json转化失败"))
		}
		utils.RedisDb.Set(context.Background(), strings.Join([]string{"token", token}, ":"), marshal, time.Minute*30)

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
	// token验证
	if !utils.ValidateToken(ctx.Query("token")) {
		ctx.JSON(http.StatusUnauthorized, pb.DouyinUserResponse{
			StatusCode: 1,
			StatusMsg:  "Token解析错误，验证失败",
		})
		return
	}
	// 包装request
	id, err := strconv.ParseInt(ctx.Query("user_id"), 10, 64)
	if err != nil {
		fmt.Println(err)
		panic(any("id字符串转化失败！！！"))
	}
	request := pb.DouyinUserRequest{
		UserId: id,
		Token:  ctx.Query("token"),
	}
	// 获取用户信息
	userInfo := dao.GetUserInfo(&request)
	ctx.JSON(http.StatusOK, pb.DouyinUserResponse{
		StatusCode: 0,
		StatusMsg:  "获取用户信息成功",
		// 由于omitempty关键字无法识别并忽略嵌套结构体的字段空值，返回的json结果会包含时间的空值
		User: userInfo,
	})
}
