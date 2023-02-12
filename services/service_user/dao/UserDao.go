package dao

import (
	"douyin-template/model/pb"
	"douyin-template/services/service_user/db"
	"douyin-template/utils"
)

// IsUserExist 注册前查询表中是否存在该user
func IsUserExist(username string) bool {
	var user pb.User
	row := db.Db.Where("name=?", username).First(&user).RowsAffected
	// 无查询结果
	if row > 0 {
		return true
	}
	return false
}

// InsertUser 添加用户
func InsertUser(request *pb.DouyinUserRegisterRequest) (int64, string) {
	// 生成唯一id
	id := utils.NewSnowId()
	db.Db.Create(&pb.User{
		Id:            id,
		Name:          request.Username,
		Password:      request.Password,
		FollowCount:   0,
		FollowerCount: 0,
	})
	return id, request.GetUsername()
}

// VerifyUser 验证密码
func VerifyUser(request *pb.DouyinUserLoginRequest) (int64, string) {
	var user pb.User
	row := db.Db.Where("name=? and password=?", request.GetUsername(), request.GetPassword()).First(&user).RowsAffected
	if row == 1 {
		return user.Id, user.GetName()
	}
	return 0, ""
}

// GetUserInfo 查询用户信息
func GetUserInfo(request *pb.DouyinUserRequest) pb.User {
	user := pb.User{}
	db.Db.Select([]string{"id", "name", "follow_count", "follower_count"}).First(&user, request.GetUserId())
	//userDto := pb.User{
	//	Id:            user.GetId(),
	//	Name:          user.GetName(),
	//	FollowCount:   user.GetFollowCount(),
	//	FollowerCount: user.GetFollowerCount(),
	//	CreatedAt:     time.Time{},
	//}
	return user
}
