package service

import (
	"context"
	"douyin-template/model"
	"douyin-template/services/service_comment/dao"
	"douyin-template/services/service_comment/rpc"
	"douyin-template/utils"
	"fmt"
	"log"
	"sort"
	"strconv"
	"sync"
	"time"
)

// OneDayOfHours 时间
var OneDayOfHours = 60 * 60 * 24
var OneMinute = 60 * 1
var OneMonth = 60 * 60 * 24 * 30
var OneYear = 365 * 60 * 60 * 24
var ExpireTime = time.Hour * 48 // 设置Redis数据热度消散时间。
const DateTime = "2006-01-02 15:04:05"

// CommentInfo 查看评论-传出的结构体-service
//type CommentInfo struct {
//	Id      int64 `json:"id,omitempty"`
//	VideoId int64 `json:"video_id"`
//	//UserInfo   model.User `json:"user,omitempty"`
//	Content    string `json:"content,omitempty"`
//	CreateDate string `json:"create_date,omitempty"`
//}
//
//type CommentData struct {
//	Id            int64     `json:"id,omitempty"`
//	UserId        int64     `json:"user_id,omitempty"`
//	Name          string    `json:"name,omitempty"`
//	FollowCount   int64     `json:"follow_count"`
//	FollowerCount int64     `json:"follower_count"`
//	IsFollow      bool      `json:"is_follow"`
//	Content       string    `json:"content,omitempty"`
//	CreateDate    time.Time `json:"create_date,omitempty"`
//}

// CountFromVideoId
// 1、使用video id 查询Comment数量
func CountFromVideoId(videoId int64) (int64, error) {
	//先在缓存中查
	cnt, err := utils.RedisDb.SCard(utils.Ctx, strconv.FormatInt(videoId, 10)).Result()
	if err != nil { //若查询缓存出错，则打印log
		//return 0, err
		log.Println("count from Redis error:", err)
	}
	log.Println("comment count Redis :", cnt)
	//1.缓存中查到了数量，则返回数量值-1（去除0值）
	if cnt != 0 {
		return cnt - 1, nil
	}
	//2.缓存中查不到则去数据库查
	cntDao, err1 := dao.Count(videoId)
	log.Println("comment count dao :", cntDao)
	if err1 != nil {
		log.Println("comment count dao err:", err1)
		return 0, nil
	}
	//将评论id切片存入Redis-第一次存储 V-C set 值：
	go func() {
		//查询评论id list
		cList, _ := dao.CommentIdList(videoId)
		//先在Redis中存储一个-1值，防止脏读
		_, _err := utils.RedisDb.SAdd(utils.Ctx, strconv.Itoa(int(videoId)), -1).Result()
		if _err != nil { //若存储Redis失败，则直接返回
			log.Println("Redis save one vId - cId 0 failed")
			return
		}
		//设置key值过期时间
		_, err := utils.RedisDb.Expire(utils.Ctx, strconv.Itoa(int(videoId)),
			time.Duration(OneMonth)*time.Second).Result()
		if err != nil {
			log.Println("Redis save one vId - cId expire failed")
		}
		//评论id循环存入Redis
		for _, commentId := range cList {
			insertRedisVideoCommentId(strconv.Itoa(int(videoId)), commentId)
		}
		log.Println("count comment save ids in Redis")
	}()
	//返回结果
	return cntDao, nil
}

// Send
// 2、发表评论
func Send(comment model.Comment, token string) (model.CommentDto, error) {
	log.Println("CommentService-Send: running") //函数已运行
	myClaim, _ := utils.ParseToken(token)
	userId, err := strconv.ParseInt(myClaim.UserId, 10, 64)
	if err != nil {
		fmt.Sprintf("字符串id转int64失败：%v", err)
	}

	//1.评论信息存储：
	comment.UserId = userId
	err = dao.InsertComment(comment)
	if err != nil {
		return model.CommentDto{}, err
	}

	//2.调用user服务获取user对象

	response, err := rpc.CommentToUserRpcClient.GetUserInfo(context.Background(), &model.DouyinUserRequest{
		UserId: userId,
		Token:  token,
	})
	if err != nil {
		fmt.Sprintf("调用user服务失败：%v", err)
	}
	user := response.GetUser()

	//3.拼接
	commentDto := model.CommentDto{
		Id:         comment.GetId(),
		User:       user,
		Content:    comment.GetContent(),
		CreateData: comment.GetCreateData(),
	}
	//将此发表的评论id存入Redis
	go func() {
		insertRedisVideoCommentId(strconv.Itoa(int(comment.GetVideoId())), strconv.Itoa(int(comment.GetId())))
		log.Println("send comment save in Redis")
	}()
	//返回结果
	return commentDto, nil
}

// DelComment
// 3、删除评论，传入评论id
func DelComment(commentId int64) error {
	log.Println("CommentService-DelComment: running") //函数已运行
	//1.先查询Redis，若有则删除，返回客户端-再go协程删除数据库；无则在数据库中删除，返回客户端。
	n, err := utils.RedisDb.Exists(utils.Ctx, strconv.FormatInt(commentId, 10)).Result()
	if err != nil {
		log.Println(err)
	}
	if n > 0 { //在缓存中有此值，则找出来删除，然后返回
		vid, err1 := utils.RedisDb.Get(utils.Ctx, strconv.FormatInt(commentId, 10)).Result()
		if err1 != nil { //没找到，返回err
			log.Println("Redis find CV err:", err1)
		}
		//删除，两个Redis都要删除
		del1, err2 := utils.RedisDb.Del(utils.Ctx, strconv.FormatInt(commentId, 10)).Result()
		if err2 != nil {
			log.Println(err2)
		}
		del2, err3 := utils.RedisDb.SRem(utils.Ctx, vid, strconv.FormatInt(commentId, 10)).Result()
		if err3 != nil {
			log.Println(err3)
		}
		log.Println("del comment in Redis success:", del1, del2) //del1、del2代表删除了几条数据

		//使用mq进行数据库中评论的删除-评论状态更新
		//评论id传入消息队列
		//rabbitmq.RmqCommentDel.Publish(strconv.FormatInt(commentId, 10))
		return nil
	}
	//不在内存中，则直接走数据库删除
	return dao.DeleteComment(commentId)
}

// GetList
// 4、查看评论列表-返回评论list
func GetList(videoId int64, userId int64, token string) ([]*model.CommentDto, error) {
	log.Println("CommentService-GetList: running") //函数已运行
	//调用dao，先查评论，再循环查用户信息：
	//1.先查询评论列表信息
	commentList, err := dao.GetCommentList(videoId)
	if err != nil {
		log.Println("CommentService-GetList: return err: " + err.Error()) //函数返回提示错误信息
		return nil, err
	}
	//当前有0条评论
	if commentList == nil {
		return nil, nil
	}

	//提前定义好切片长度
	commentInfoList := make([]*model.CommentDto, len(commentList))

	wg := &sync.WaitGroup{}
	wg.Add(len(commentList))
	idx := 0
	for _, comment := range commentList {
		//2.调用方法组装评论信息，再append
		var commentData model.CommentDto
		//将评论信息进行组装，添加想要的信息,插入从数据库中查到的数据
		go func(comment model.Comment) {
			oneComment(&commentData, &comment, userId, token)
			//3.组装list
			//commentInfoList = append(commentInfoList, commentData)
			commentInfoList[idx] = &commentData
			idx = idx + 1
			wg.Done()
		}(comment)
	}
	wg.Wait()
	//评论排序-按照主键排序
	sort.Sort(CommentSlice(commentInfoList))
	//------------------------法二结束----------------------------

	//协程查询Redis中是否有此记录，无则将评论id切片存入Redis
	go func() {
		//1.先在缓存中查此视频是否已有评论列表
		cnt, err1 := utils.RedisDb.SCard(utils.Ctx, strconv.FormatInt(videoId, 10)).Result()
		if err1 != nil { //若查询缓存出错，则打印log
			//return 0, err
			log.Println("count from Redis error:", err)
		}
		//2.缓存中查到了数量大于0，则说明数据正常，不用更新缓存
		if cnt > 0 {
			return
		}
		//3.缓存中数据不正确，更新缓存：
		//先在Redis中存储一个-1 值，防止脏读
		_, _err := utils.RedisDb.SAdd(utils.Ctx, strconv.Itoa(int(videoId)), -1).Result()
		if _err != nil { //若存储Redis失败，则直接返回
			log.Println("Redis save one vId - cId 0 failed")
			return
		}
		//设置key值过期时间
		_, err2 := utils.RedisDb.Expire(utils.Ctx, strconv.Itoa(int(videoId)),
			time.Duration(OneMonth)*time.Second).Result()
		if err2 != nil {
			log.Println("Redis save one vId - cId expire failed")
		}
		//将评论id循环存入Redis
		for _, _comment := range commentInfoList {
			insertRedisVideoCommentId(strconv.Itoa(int(videoId)), strconv.Itoa(int(_comment.Id)))
		}
		log.Println("comment list save ids in Redis")
	}()

	log.Println("CommentService-GetList: return list success") //函数执行成功，返回正确信息
	return commentInfoList, nil
}

// 在redis中存储video_id对应的comment_id 、 comment_id对应的video_id
func insertRedisVideoCommentId(videoId string, commentId string) {
	//在redis-RdbVCid中存储video_id对应的comment_id
	_, err := utils.RedisDb.SAdd(utils.Ctx, videoId, commentId).Result()
	if err != nil { //若存储redis失败-有err，则直接删除key
		log.Println("redis save send: vId - cId failed, key deleted")
		utils.RedisDb.Del(utils.Ctx, videoId)
		return
	}
	//在redis-RdbCVid中存储comment_id对应的video_id
	_, err = utils.RedisDb.Set(utils.Ctx, commentId, videoId, 0).Result()
	if err != nil {
		log.Println("redis save one cId - vId failed")
	}
}

// 此函数用于给一个评论赋值：评论信息+用户信息 填充
func oneComment(comment *model.CommentDto, com *model.Comment, userId int64, token string) {
	var wg sync.WaitGroup
	wg.Add(1)

	//2.调用user服务获取user对象
	myClaim, _ := utils.ParseToken(token)
	userId, err := strconv.ParseInt(myClaim.UserId, 10, 64)
	if err != nil {
		fmt.Sprintf("字符串id转int64失败：%v", err)
	}
	response, err := rpc.CommentToUserRpcClient.GetUserInfo(context.Background(), &model.DouyinUserRequest{
		UserId: userId,
		Token:  token,
	})
	if err != nil {
		fmt.Sprintf("调用user服务失败：%v", err)
	}
	user := response.GetUser()
	comment.Id = com.GetId()
	comment.Content = com.GetContent()
	comment.CreateData = com.GetCreateData()
	comment.User = user

	wg.Done()
	wg.Wait()
}

// CommentSlice 此变量以及以下三个函数都是做排序-准备工作
type CommentSlice []*model.CommentDto

func (a CommentSlice) Len() int { //重写Len()方法
	return len(a)
}
func (a CommentSlice) Swap(i, j int) { //重写Swap()方法
	a[i], a[j] = a[j], a[i]
}
func (a CommentSlice) Less(i, j int) bool { //重写Less()方法
	return a[i].Id > a[j].Id
}
