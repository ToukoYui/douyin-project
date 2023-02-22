package controller

import (
	"context"
	"douyin-template/model"
	"douyin-template/services/service_user/dao"
	"douyin-template/utils"
	"encoding/json"
	"fmt"
	"golang.org/x/time/rate"
	"strconv"
	"strings"
	"time"
)

type Server struct {
	// 继承 protoc-gen-go-grpc 生成的服务端代码
	model.UnimplementedUserSrvServer
}

func (s *Server) Login(ctx context.Context, request *model.DouyinUserLoginRequest) (*model.DouyinUserLoginResponse, error) {
	fmt.Println("进入login")
	userName := request.GetUsername()
	if !requestIsAllow(5, userName) {
		return &model.DouyinUserLoginResponse{
			StatusCode: 2,
			StatusMsg:  "操作过于频繁!",
		}, nil
	}

	if !dao.IsUserExist(userName) {
		return &model.DouyinUserLoginResponse{
			StatusCode: 1,
			StatusMsg:  "该用户不存在",
		}, nil
	}
	// 登录成功
	if id, name := dao.VerifyUser(request); id != 0 {
		fmt.Println("登录用户id：", id)
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

		return &model.DouyinUserLoginResponse{
			StatusCode: 0,
			StatusMsg:  "登录成功",
			UserId:     id,
			Token:      token,
		}, nil
	}

	return &model.DouyinUserLoginResponse{
		StatusCode: 1,
		StatusMsg:  "密码错误",
	}, nil
}

func (s *Server) Register(ctx context.Context, request *model.DouyinUserRegisterRequest) (*model.DouyinUserRegisterResponse, error) {
	userName := request.GetUsername()
	fmt.Println(userName)
	if !requestIsAllow(5, userName) {
		return &model.DouyinUserRegisterResponse{
			StatusCode: 2,
			StatusMsg:  "操作过于频繁!",
		}, nil
	}
	//请求打到数据库上了
	if dao.IsUserExist(userName) {
		return &model.DouyinUserRegisterResponse{
			StatusCode: 1,
			StatusMsg:  "注册失败，该用户名已存在",
		}, nil
	}
	// 将user添加进表中
	id, name := dao.InsertUser(request)
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
	return &model.DouyinUserRegisterResponse{
		StatusCode: 0,
		StatusMsg:  "注册成功",
		UserId:     id,
		Token:      token,
	}, nil
}

type TokenInfo struct {
	UserId   string `json:"user_id"`
	UserName string `json:"username"`
}

func (s *Server) GetUserInfo(ctx context.Context, request *model.DouyinUserRequest) (*model.DouyinUserResponse, error) {
	// token验证
	if request.GetToken() != "" {
		if !utils.ValidateToken(request.GetToken()) {
			return &model.DouyinUserResponse{
				StatusCode: 1,
				StatusMsg:  "Token解析错误，验证失败",
			}, nil
		}
	}
	// 获取用户信息
	userInfo := dao.GetUserInfo(request)
	return &model.DouyinUserResponse{
		StatusCode: 0,
		StatusMsg:  "获取用户信息成功",
		// 由于omitempty关键字无法识别并忽略嵌套结构体的字段空值，返回的json结果会包含时间的空值
		User: &userInfo,
	}, nil
}

/*
*

	r : 每rate毫秒放入一个令牌
	capacity : 令牌桶的大小
	identify : 标志接口请求者的一个信息
*/
func requestIsAllow(capacity int, identify string) bool {
	/*登录接口限流实现*/
	//1. 创建新的限流器
	//参数说明
	// r: 每10ms可以接受一次注册
	// b: 桶中可以放100的令牌
	// key:限制userName
	limiter := utils.NewLimiter(rate.Every(1000*time.Millisecond), capacity, identify)
	//3. 检查是否超过限流的限制
	return limiter.Allow()
}
