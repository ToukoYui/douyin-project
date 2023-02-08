package pb

type DouyinPublishListRequest struct {
	UserId int64  `json:"user_id"`
	Token  string `json:"token"`
}

type DouyinPublishListResponse struct {
	StatusCode    int32      `json:"status_code"`
	StatusMessage string     `json:"status_message,omitempty"`
	VideoList     []VideoDto `json:"video_list,omitempty"`
}
