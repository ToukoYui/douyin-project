package pb

type Favorite struct {
	Id         int64 `json:"id,omitempty"`       // 视频唯一标识
	UserId     int64 `json:"user_id,omitempty"`  // 视频作者id
	VideoId    int64 `json:"video_id,omitempty"` // 视频播放地址
	IsFavorite bool  `json:"is_favorite"`        // 视频封面地址
}

func (x *Favorite) TableName() string {
	return "favorite"
}
