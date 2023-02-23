package dao

import (
	"context"
	"douyin-template/model"
	"douyin-template/model/pb"
	"douyin-template/services/consts"
	"douyin-template/services/service_video/db"
	"douyin-template/services/service_video/rpc"
	"douyin-template/utils"
	"fmt"
	"github.com/goccy/go-json"
	"strconv"
	"time"
)

var ctx = context.Background()

func existsKey(key string) bool {
	//返回1表示存在，0表示不存在
	result, _ := utils.RedisDb.Do(context.Background(), "exists", key).Result()
	exist := result.(int64)
	fmt.Println(exist)
	if exist == 1 {
		return false
	}
	return true

}

// GetVideoInfoList 获取视频信息
func GetVideoInfoList(request *model.DouyinFeedRequest) ([]*model.VideoDto, int64) {
	// 根据LatestTime降序查询视频列表
	var videoList []model.Video
	latestTime := time.Unix(request.GetLatestTime(), 0)
	fmt.Println("请求时间为：", latestTime)
	/*添加缓存*/
	//1.检查缓存中是否有数据
	data := utils.RedisDb.Get(ctx, consts.VIDEO_CACHE_KEY)
	if existsKey(consts.VIDEO_CACHE_KEY) {
		//3. 无数据,请求数据并返回
		//3.1 先请求数据库
		/***********/
		row := db.Db.Where("created_time<=?", latestTime).Order("created_time desc").Find(&videoList).RowsAffected
		if row <= 0 {
			fmt.Printf("无查询结果,批量查询失败,查询时间参数：%v", latestTime)
		}
		fmt.Printf("videoList结果大小：%d,结果列表:%#v\n", len(videoList), videoList)
		//3.2 再重建缓存
		jsonString, err := json.Marshal(videoList)
		fmt.Println(jsonString)
		if err != nil {
			return nil, 0
		}
		result, err := utils.RedisDb.Set(ctx, consts.VIDEO_CACHE_KEY, string(jsonString), 10*time.Minute).Result()
		if err != nil {
			return nil, 0
		}
		fmt.Printf("执行结果%v", result)
	} else {
		//否则的话就是有数据,那么就将这个数据解析
		bytes := []byte(data.Val())
		err := json.Unmarshal(bytes, &videoList) //要传指针
		fmt.Printf("缓存中有数据:%s\n", videoList)
		if err != nil {
			return nil, 0
		}
	}

	// 获取视频列表中的最早发布的时间
	resultList := make([]*model.VideoDto, len(videoList))
	last := len(videoList) - 1
	if last < 0 {
		return resultList, latestTime.Unix()
	}
	earlyTime := videoList[last].CreatedAt

	for i, item := range videoList {
		//根据user_id查询User对象
		//2.调用user服务获取user对象
		myClaim, _ := utils.ParseToken(request.GetToken())
		userId, err := strconv.ParseInt(myClaim.UserId, 10, 64)
		if err != nil {
			fmt.Sprintf("字符串id转int64失败：%v", err)
		}
		response, err := rpc.VideoToUserRpcClient.GetUserInfo(context.Background(), &model.DouyinUserRequest{
			UserId: userId,
			Token:  request.GetToken(),
		})
		if err != nil {
			fmt.Sprintf("调用user服务失败：%v", err)
		}
		user := response.GetUser()
		// 查询favorite表获取is_favorite
		isFavoriteRequestmodel := model.IsFavoriteRequest{
			UserId:  user.GetId(),
			VideoId: item.GetId(),
		}
		isFavorite, err := rpc.VideoToFavoriteRpcClient.CheckIsFavorite(context.Background(), &isFavoriteRequestmodel)
		if err != nil {
			fmt.Sprintf("video调用favorit失败：%v", err)
		}

		//favorite := pb.Favorite{}

		isFav := false
		if isFavorite.GetIsFavorite() == 1 {
			isFav = true
		}
		// 包装视频列表结果
		resultList[i] = &model.VideoDto{
			Id:            item.GetId(),
			Author:        user,
			PlayUrl:       item.GetPlayUrl(),
			CoverUrl:      item.GetCoverUrl(),
			FavoriteCount: item.GetFavoriteCount(),
			CommentCount:  item.GetCommentCount(),
			IsFavorite:    isFav,
			Title:         item.GetTitle(),
		}
	}
	return resultList, earlyTime.Unix()
}

// CreateVideoInfo 添加视频信息
func CreateVideoInfo(video *model.Video) {
	ctx := context.Background()
	//1.删除视频流缓存
	utils.RedisDb.Del(ctx, consts.VIDEO_CACHE_KEY)
	//2.删除用户视频缓存
	//2.1 组装key
	key := consts.VIDEO_SINGLE_CACHE_KEY + strconv.FormatInt(video.UserId, 10)
	utils.RedisDb.Del(ctx, key)
	row := db.Db.Create(video).RowsAffected
	if row != 1 {
		panic(any("添加视频信息失败！！！"))
	}
}

// GetPublishVideoList 获取单个用户的发布视频列表
func GetPublishVideoList(request *model.DouyinPublishListRequest) []*model.VideoDto {
	ctx := context.Background()
	/*添加缓存*/
	//1. 组装key
	key := consts.VIDEO_SINGLE_CACHE_KEY + strconv.FormatInt(request.UserId, 10)
	//2. 查询是否有数据
	exists := existsKey(key)
	// 查询该用户发布过的视频信息
	var videoList []model.Video
	if !exists {
		//3.1 重建缓存
		db.Db.Where("user_id=?", request.UserId).Find(&videoList)
		//3.2 序列化缓存数据
		jsonString, err := json.Marshal(videoList)
		if err != nil {
			return nil
		}
		fmt.Printf("缓存中没有数据!%s", jsonString)
		//3.3 将数据加入到缓存中
		utils.RedisDb.Set(ctx, key, jsonString, time.Second*10)
	} else {
		//4. 走缓存
		data := utils.RedisDb.Get(ctx, key)
		bytes, err := data.Bytes()
		if err != nil {
			return nil
		}
		//4.1 将缓存中的对象反序列化为videoList
		err = json.Unmarshal(bytes, videoList)
		fmt.Printf("缓存中有数据,%s", videoList)
		if err != nil {
			return nil
		}
	}

	// 遍历包装
	resultList := make([]*model.VideoDto, len(videoList))
	fmt.Println("调用User服务查询对象")
	for i, item := range videoList {
		//调用user服务查询对象，根据user_id查询User对象，此处可以通过批量处理优化
		userInfoResp, err := rpc.VideoToUserRpcClient.GetUserInfo(ctx, &model.DouyinUserRequest{
			UserId: item.GetUserId(),
			Token:  request.GetToken(),
		})
		if err != nil {
			fmt.Sprintf("调用User服务查询对象失败%v", err)
		}
		user := *userInfoResp.GetUser()

		// 查询favorite表获取is_favorite   todo 调用like服务查询is_favorite
		favorite := pb.Favorite{}
		db.Db.Select("is_favorite").Where("user_id=? and video_id=?", user.GetId(), item.GetId()).First(&favorite)
		// 包装视频列表结果
		resultList[i] = &model.VideoDto{
			Id:            item.GetId(),
			Author:        &user,
			PlayUrl:       item.GetPlayUrl(),
			CoverUrl:      item.GetCoverUrl(),
			FavoriteCount: item.GetFavoriteCount(),
			CommentCount:  item.GetCommentCount(),
			IsFavorite:    favorite.IsFavorite,
			Title:         item.GetTitle(),
		}
	}
	return resultList

}

// GetLikedVideo 根据userid和videoid查询video
func GetLikedVideo(request *model.DouyinUseridAndVideoid) *model.Video {
	video := model.Video{}
	db.Db.Where("user_id=? and id=?", request.GetUserId(), request.GetVideoId()).First(&video)
	return &video
}
