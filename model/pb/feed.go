package pb

import "time"

func (x *Video) TableName() string {
	return "video"
}

type DouyinFeedRequest struct {
	LatestTime int64  `json:"latest_time,omitempty"` // 可选参数，限制返回视频的最新投稿时间戳，精确到秒，不填表示当前时间
	Token      string `json:"token,omitempty"`       // 可选参数，登录用户设置
}

func (x *DouyinFeedRequest) GetLatestTime() int64 {
	if x != nil {
		return x.LatestTime
	}
	return 0
}

func (x *DouyinFeedRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type DouyinFeedResponse struct {
	StatusCode int32       `json:"status_code"`          // 状态码，0-成功，其他值-失败
	StatusMsg  string      `json:"status_msg,omitempty"` // 返回状态描述
	VideoList  *[]VideoDto `json:"video_list,omitempty"` // 视频列表
	NextTime   int64       `json:"next_time,omitempty"`  // 本次返回的视频中，发布最早的时间，作为下次请求时的latest_time
}

func (*DouyinFeedResponse) ProtoMessage() {}

func (x *DouyinFeedResponse) GetStatusCode() int32 {
	if x != nil {
		return x.StatusCode
	}
	return 0
}

func (x *DouyinFeedResponse) GetStatusMsg() string {
	if x != nil {
		return x.StatusMsg
	}
	return ""
}

func (x *DouyinFeedResponse) GetVideoList() *[]VideoDto {
	if x != nil {
		return x.VideoList
	}
	return nil
}

func (x *DouyinFeedResponse) GetNextTime() int64 {
	if x != nil {
		return x.NextTime
	}
	return 0
}

type VideoDto struct {
	//state         protoimpl.MessageState
	//sizeCache     protoimpl.SizeCache
	//unknownFields protoimpl.UnknownFields

	Id            int64  `json:"id,omitempty"`             // 视频唯一标识
	Author        User   `json:"author,omitempty"`         // 视频作者信息
	PlayUrl       string `json:"play_url,omitempty"`       // 视频播放地址
	CoverUrl      string `json:"cover_url,omitempty"`      // 视频封面地址
	FavoriteCount int64  `json:"favorite_count,omitempty"` // 视频的点赞总数
	CommentCount  int64  `json:"comment_count,omitempty"`  // 视频的评论总数
	IsFavorite    bool   `json:"is_favorite"`              // true-已点赞，false-未点赞
	Title         string `json:"title,omitempty"`          // 视频标题
}

func (*VideoDto) ProtoMessage() {}

func (x *VideoDto) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *VideoDto) GetAuthor() User {
	if x != nil {
		return x.Author
	}
	return User{}
}

func (x *VideoDto) GetPlayUrl() string {
	if x != nil {
		return x.PlayUrl
	}
	return ""
}

func (x *VideoDto) GetCoverUrl() string {
	if x != nil {
		return x.CoverUrl
	}
	return ""
}

func (x *VideoDto) GetFavoriteCount() int64 {
	if x != nil {
		return x.FavoriteCount
	}
	return 0
}

func (x *VideoDto) GetCommentCount() int64 {
	if x != nil {
		return x.CommentCount
	}
	return 0
}

func (x *VideoDto) GetIsFavorite() bool {
	if x != nil {
		return x.IsFavorite
	}
	return false
}

func (x *VideoDto) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

type Video struct {
	Id            int64     `json:"id,omitempty"`        // 视频唯一标识
	UserId        int64     `json:"user_id,omitempty"`   // 视频作者id
	PlayUrl       string    `json:"play_url,omitempty"`  // 视频播放地址
	CoverUrl      string    `json:"cover_url,omitempty"` // 视频封面地址
	FavoriteCount int64     `json:"favorite_count"`      // 视频的点赞总数
	CommentCount  int64     `json:"comment_count"`       // 视频的评论总数
	Title         string    `json:"title,omitempty"`     // 视频标题
	CreatedAt     time.Time `json:"created_time,omitempty" gorm:"column:created_time"`
}

func (*Video) ProtoMessage() {}

func (x *Video) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Video) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *Video) GetPlayUrl() string {
	if x != nil {
		return x.PlayUrl
	}
	return ""
}

func (x *Video) GetCoverUrl() string {
	if x != nil {
		return x.CoverUrl
	}
	return ""
}

func (x *Video) GetFavoriteCount() int64 {
	if x != nil {
		return x.FavoriteCount
	}
	return 0
}

func (x *Video) GetCommentCount() int64 {
	if x != nil {
		return x.CommentCount
	}
	return 0
}

func (x *Video) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}
