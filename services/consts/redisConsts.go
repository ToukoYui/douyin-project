package consts

/**
* redis常见业务键
*
* @author: 张庭杰
* @date: 2023年02月21日 19:57
 */

// VIDEO_GET_KEY 视频ID的唯一性ID
const VIDEO_GET_KEY string = "video:getKey:"

// VIDEO_CACHE_KEY 当前视频流的缓存
const VIDEO_CACHE_KEY string = "video:whole_cache:"

// VIDEO_SINGLE_CACHE_KEY 单个用户点赞缓存{userId->videoInfo}
const VIDEO_SINGLE_CACHE_KEY string = "video:single_cache:"

// VIDEO_LIKED_KEY 视频点赞集合{视频ID->UserId}
const VIDEO_LIKED_KEY string = "video:liked:"

// VIDEO_USER_LIKED_KEY 用户ID->视频列表
const VIDEO_USER_LIKED_KEY string = "video:user_liked:"
