namespace go user

struct douyin_user_register_request {
  // 注册用户名，最短2个字符，最长32个字符
  1: required string username (vt.min_size = "2", vt.max_size = "32", api.vd = "len($)>1 && len($)<32")
  // 密码，最短6个字符，最长32个字符，使用 api.vd 写正则会有点问题，做为替代在 util 里实现了正则的验证
  2: required string password (vt.min_size = "6", vt.max_size = "32", api.vd = "len($)>5 && len($)<32")
}

struct douyin_user_register_response {
  1: required i64 status_code // 状态码，0-成功，其他值-失败
  2: optional string status_msg // 返回状态描述
  3: required i64 user_id // 用户id
  4: required string token // 用户鉴权token
}

struct douyin_user_login_request {
  1: required string username (vt.min_size = "2", vt.max_size = "32", api.vd = "len($)>1 && len($)<33") // 登录用户名，最短2个字符，最长32个字符
  2: required string password (vt.min_size = "6", vt.max_size = "32", api.vd = "len($)>5 && len($)<33") // 登录密码，最短6个字符，最长32个字符
}

struct douyin_user_login_response {
  1: required i64 status_code // 状态码，0-成功，其他值-失败
  2: optional string status_msg // 返回状态描述
  3: required i64 user_id  // 用户id
  4: required string token // 用户鉴权token
}

struct douyin_user_request {
  1: required i64 user_id (vt.gt = "0", api.vd = "$>0") // 用户id
  2: required string token (api.header = "token", api.query = "token") // 用户鉴权token
}

struct douyin_user_response {
  1: required i64 status_code // 状态码，0-成功，其他值-失败
  2: optional string status_msg // 返回状态描述
  3: required UserInfo user // 用户信息
}

struct douyin_relation_action_request {
  1: required string token (api.header = "token", api.query = "token") // 用户鉴权token
  2: required i64 to_user_id (vt.gt = "0", api.vd="$>0") // 对方用户id
  3: required i8 action_type (vt.in = "1", vt.in = "2", api.vd = "$==1||$==2") // 1-关注，2-取消关注
}

struct douyin_relation_action_response {
  1: required i64 status_code // 状态码，0-成功，其他值-失败
  2: optional string status_msg // 返回状态描述
}

struct douyin_relation_follow_list_request {
  1: required i64 user_id (vt.gt = "0", api.vd="$>0") // 用户id
  2: required string token (api.header = "token", api.query = "token") // 用户鉴权token
}

struct douyin_relation_follow_list_response {
  1: required i64 status_code // 状态码，0-成功，其他值-失败
  2: optional string status_msg // 返回状态描述
  3: list<FollowUser> user_list // 用户信息列表
}

// 关注用户信息，关注总数和粉丝总数在跳转到具体用户信息页面时再重新获取
struct FollowUser {
  1: required i64 id // 用户id
  2: required string name  // 用户名称
  3: required bool is_follow  // true-已关注，false-未关注
  4: required string avatar  // 用户头像Url
}

struct douyin_relation_follower_list_request {
  1: required i64 user_id (vt.gt = "0", api.vd="$>0")  // 用户id
  2: required string token (api.header = "token", api.query = "token") // 用户鉴权token
}

struct douyin_relation_follower_list_response {
  1: required i64 status_code // 状态码，0-成功，其他值-失败
  2: optional string status_msg  // 返回状态描述
  3: list<FollowerUser> user_list  // 用户列表
}

// 粉丝用户信息，关注总数和粉丝总数在跳转到具体用户信息页面时再重新获取
struct FollowerUser {
  1: required i64 id // 用户id
  2: required string name  // 用户名称
  3: required bool is_follow  // true-已关注，false-未关注
  4: required string avatar  // 用户头像Url
}

struct douyin_relation_friend_list_request {
  1: required i64 user_id (vt.gt = "0", api.vd="$>0")  // 用户id
  2: required string token (api.header = "token", api.query = "token")  // 用户鉴权token
}

struct douyin_relation_friend_list_response {
  1: required i64 status_code  // 状态码，0-成功，其他值-失败
  2: optional string status_msg  // 返回状态描述
  3: list<FriendUser> user_list // 用户列表
}

struct User {
  1: required i64 id // 用户id
  2: required string name  // 用户名称
  3: optional i64 follow_count  // 关注总数
  4: optional i64 follower_count  // 粉丝总数
  5: required bool is_follow  // true-已关注，false-未关注
  6: required string avatar  // 用户头像Url
}

struct UserInfo {
  1: required i64 id // 用户id
  2: required string name  // 用户名称
  3: required i64 follow_count  // 关注总数
  4: required i64 follower_count  // 粉丝总数
  5: required bool is_follow  // true-已关注，false-未关注
  6: required string avatar  // 用户头像Url
  7: required string background_image //用户个人页顶部大图
  8: required string signature //个人简介
  9: required i64 total_favorited //获赞数量
  10: required i64 work_count  // 用户作品数
  11: required i64 favorite_count  // 用户点赞的视频数
}

struct FriendUser {
  1: required i64 id // 用户id
  2: required string name  // 用户名称
  3: required bool is_follow  // true-已关注，false-未关注
  4: required string avatar  // 用户头像Url
  5: optional string message // 和该好友的最新聊天消息
  6: required i8 msgType // message消息的类型，0 => 当前请求用户接收的消息， 1 => 当前请求用户发送的消息
}

service UserService {
    // User
    douyin_user_register_response Register(1: douyin_user_register_request req) (api.post="/douyin/user/register/")
    douyin_user_login_response Login(1: douyin_user_login_request req) (api.post="/douyin/user/login/")
    douyin_user_response GetUserInfo(1: douyin_user_request req) (api.get="/douyin/user/")
    // Relation
    douyin_relation_action_response Follow(1: douyin_relation_action_request req) (api.post="/douyin/relation/action/")
    douyin_relation_follow_list_response GetFollowList(1: douyin_relation_follow_list_request req) (api.get="/douyin/relation/follow/list/")
    douyin_relation_follower_list_response GetFollowerList(1: douyin_relation_follower_list_request req) (api.get="/douyin/relation/follower/list/")
    douyin_relation_friend_list_response GetFriendList(1: douyin_relation_friend_list_request req) (api.get="/douyin/relation/friend/list/")
}