package dao

import (
	"context"
	"douyin-template/model"
	"douyin-template/services/consts"
	"douyin-template/utils"
	"fmt"
	"github.com/go-redis/redis/v8"
	"strconv"
	"time"
)

var ctx = context.Background()

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
	/*
		rowsAffected := db.Db.Model(&model.Favorite{}).
			Where("user_id=? and video_id=? ", userId, request.GetVideoId()).
			Update("is_favorite", request.GetActionType()).RowsAffected
		//rowsAffected := db.Db.Model(&favorite).Update("is_favorite", request.GetActionType()).RowsAffected
		if rowsAffected <= 0 {
			fmt.Sprintf("点赞更新操作失败")
			return false
		}
		return true
	*/
	/*基于redis实现点赞功能*/
	err = handleIsLike(request.VideoId, request.ActionType == 1, userId)
	if err != nil {
		return false
	}
	return true
}

// 处理用户点赞的状态
func handleIsLike(videoId int64, isFavorite bool, userId int64) (err error) {
	/*组装对应的key*/
	//视频id->用户ID
	videoLikedKey := consts.VIDEO_LIKED_KEY + strconv.FormatInt(videoId, 10)
	//用户id->视频list
	userLikedKey := consts.VIDEO_USER_LIKED_KEY + strconv.FormatInt(userId, 10)
	/*检查一下这个用户是否在set中*/
	isMember, err := utils.RedisDb.SIsMember(ctx, videoLikedKey, userId).Result()
	if err != nil {
		return err
	}
	if isFavorite {
		//1. 判断类型:用户点赞,那么就是要将这个用户加入到set中
		if isMember {
			//如果已经在集合中了,那么就是无效操作,直接返回即可
			return nil
		}
		//1.1 否则的话,就将这个用户加入到redis中
		success, err := utils.RedisDb.SAdd(ctx, videoLikedKey, userId).Result()
		if success != 1 { //操作成功
			return err
		}
		//1.2 同时,将该视频加入到用户的点赞列表中
		_, err = utils.RedisDb.ZAdd(ctx, userLikedKey, &redis.Z{
			Score:  float64(time.Now().Unix()),
			Member: videoId,
		}).Result()
		if err != nil {
			return err
		}
		return nil
	}
	//2. 如果是取消点赞,那么将用户从set中删除
	if !isMember { //如果用户都在不在集合中,又想取消点赞,那么肯定是错误的操作
		return nil
	}
	//2.1 否则的话就将这个用户从里面删除
	success, err := utils.RedisDb.SRem(ctx, videoLikedKey, userId).Result()
	if success != 1 || err != nil {
		return err
	}
	//2.2 并且将这个元素从集合中删除
	success, err = utils.RedisDb.ZRem(ctx, userLikedKey, videoId).Result()
	if success != 1 || err != nil {
		return err
	}
	return nil
}

// GetUserIdAndVideoId 查询被点赞过的视频id和用户id
func GetUserIdAndVideoId(request *model.DouyinFavoriteListRequest) []model.Favorite {
	var likedVideoIdSli []model.Favorite
	/*
		rowsAffected := db.Db.Select("user_id", "video_id").Where("user_id =? and is_favorite=?", request.GetUserId(), 1).Find(&likedVideoIdSli).RowsAffected
		if rowsAffected <= 0 {
			fmt.Println("查询点赞视频id错误")
		}
	*/
	//1. 组装key
	key := consts.VIDEO_USER_LIKED_KEY + strconv.FormatInt(request.UserId, 10)
	//2. 找到对应的Zset
	list := utils.RedisDb.ZRange(ctx, key, 0, -1).Val()
	//3. 将对应的数据赋值上去
	likedVideoIdSli = make([]model.Favorite, len(list))
	for i, item := range list {
		//3.1 封装videoId
		videoId, err := strconv.ParseInt(item, 10, 64)
		if err != nil {
			return nil
		}
		likedVideoIdSli[i].VideoId = videoId
		//3.2 封装userId
		likedVideoIdSli[i].UserId = request.UserId
		//3.3 封装isFavorite
		likedVideoIdSli[i].IsFavorite = 1
	}
	return likedVideoIdSli
}

func IsLike(userId int64, videoId int64) bool {
	videoLikedKey := consts.VIDEO_LIKED_KEY + strconv.FormatInt(videoId, 10)
	isMember, err := utils.RedisDb.SIsMember(ctx, videoLikedKey, userId).Result()
	fmt.Printf("用户点过赞")
	if err != nil {
		return false
	}
	return isMember
}
