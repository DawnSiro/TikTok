namespace go video

enum ErrCode {
	Success                                  = 0     // 一切正常
	Client                                   = 10001 // 用户端错误 一级宏观错误码
	UserRegistration                         = 10100 // 用户注册错误 二级宏观错误码
	UsernameVerificationFailed               = 10110 // 用户名校验失败 三级宏观错误码
	UsernameAlreadyExists                    = 10111
	PasswordVerificationFailed               = 10120 // 密码校验失败 三级宏观错误码
	PasswordLengthNotEnough                  = 10121
	PasswordStrengthNotEnough                = 10122
	UserLogin                                = 10200 // 用户登陆异常 二级宏观错误码
	UserAccountDoesNotExist                  = 10201
	UserPassword                             = 10210
	PasswordNumberOfTimesExceeds             = 10211
	UserIdentityVerificationFailed           = 10220 // 用户身份校验失败 （Token错误等）
	UserLoginHasExpired                      = 10230
	AccessPermission                         = 10300 // 访问权限异常 二级宏观错误码
	DeletePermission                         = 10310 // 删除权限异常 普通用户不能删除别人的评论
	VideoLikeLimit                           = 10320 //
	UserRequestParameter                     = 10400 // 用户请求参数错误 二级宏观错误码
	RepeatOperationError                     = 10410 // 用户重复操作
	IllegalUserInput                         = 10430
	ContainsProhibitedSensitiveWords         = 10431
	UserUploadFile                           = 10500 // 用户上传文件异常 二级宏观错误码
	FileTypeUploadedNotMatch                 = 10501
	FileTypeUploadedNotSupport               = 10502
	VideoUploadedTooLarge                    = 10504
	Service                                  = 20000 // 未知异常
	SystemExecution                          = 20001 // 系统执行出错 一级宏观错误码
	SystemExecutionTimeout                   = 20100 // 系统执行超时 二级宏观错误码
	SystemDisasterToleranceFunctionTriggered = 20200 // 系统容灾功能被触发 二级宏观错误码
	SystemResource                           = 20300 // 系统资源异常 二级宏观错误码
	CallingThirdPartyService                 = 30001 // 调用第三方服务出错 一级宏观错误码
	MiddlewareService                        = 30100 // 中间件服务出错 二级宏观错误码
	RPCService                               = 30110
	RPCServiceNotFind                        = 30111
	RPCServiceNotRegistered                  = 30112
	InterfaceNotExist                        = 30113
	CacheService                             = 30120
	KeyLengthExceedsLimit                    = 30121
	ValueLengthExceedsLimit                  = 30122
	StorageCapacityFull                      = 30123
	UnsupportedDataFormat                    = 30124
	DatabaseService                          = 30200 // 数据库服务出错 二级宏观错误码
	TableDoesNotExist                        = 30211
	ColumnDoesNotExist                       = 30212
	DatabaseDeadlock                         = 30231
}

struct douyin_response {
  1: required i64 status_code // 状态码，0-成功，其他值-失败
  2: required string status_msg // 返回状态描述
}

struct douyin_feed_request {
  1: optional i64 latest_time // 可选参数，限制返回视频的最新投稿时间戳，精确到秒，不填表示当前时间
  2: optional string token (api.header = "token", api.query = "token") // 可选参数，登录用户设置
}

struct douyin_feed_response {
  1: required i64 status_code   // 状态码，0-成功，其他值-失败
  2: optional string status_msg // 返回状态描述
  3: list<Video> video_list     // 视频列表
  4: optional i64 next_time     // 本次返回的视频中，发布最早的时间，作为下次请求时的latest_time
}

struct douyin_publish_action_request {
  1: required string token (api.header = "token", api.form = "token") // 用户鉴权token
//  2: optional binary data // 视频数据
  2: required string title (vt.min_size = "1", vt.max_size = "63", api.vd = "len($)>0&&len($)<64") // 视频标题
}

struct douyin_publish_action_response {
  1: required i64 status_code  // 状态码，0-成功，其他值-失败
  2: optional string status_msg  // 返回状态描述
}

struct douyin_publish_list_request {
  1: required i64 user_id (vt.gt = "0", api.vd="$>0")  // 用户id
  2: required string token (api.header = "token", api.query = "token")  // 用户鉴权token
}

struct douyin_publish_list_response {
  1: required i64 status_code  // 状态码，0-成功，其他值-失败
  2: optional string status_msg  // 返回状态描述
  3: list<Video> video_list  // 用户发布的视频列表
}


struct douyin_favorite_action_request {
  1: required string token (api.header = "token", api.query = "token")  // 用户鉴权token
  2: required i64 video_id (vt.gt = "0", api.vd="$>0")  // 视频id
  3: required i8 action_type (vt.in = "1", vt.in = "2", api.vd = "$==1||$==2") // 1-点赞，2-取消点赞
}

struct douyin_favorite_action_response {
  1: required i64 status_code // 状态码，0-成功，其他值-失败
  2: optional string status_msg // 返回状态描述
}

struct douyin_favorite_list_request {
  1: required i64 user_id (vt.gt = "0", api.vd="$>0") // 用户id
  2: required string token (api.header = "token", api.query = "token") // 用户鉴权token
}

struct douyin_favorite_list_response {
  1: required i64 status_code // 状态码，0-成功，其他值-失败
  2: optional string status_msg // 返回状态描述
  3: list<Video> video_list // 用户点赞视频列表
}

struct Video {
  1: required i64 id // 视频唯一标识
  2: required UserInfo author // 视频作者信息
  3: required string play_url // 视频播放地址
  4: required string cover_url // 视频封面地址
  5: required i64 favorite_count // 视频的点赞总数
  6: required i64 comment_count // 视频的评论总数
  7: required bool is_favorite // true-已点赞，false-未点赞
  8: required string title // 视频标题
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

service VideoService {
    // Feed
    douyin_feed_response GetFeed(1: douyin_feed_request req) (api.get="/douyin/feed/")
    // Publish
    douyin_publish_action_response PublishAction(1: douyin_publish_action_request req) (api.post="/douyin/publish/action/")
    douyin_publish_list_response GetPublishVideos(1: douyin_publish_list_request req) (api.get="/douyin/publish/list/")
    // Favorite
    douyin_favorite_action_response FavoriteVideo(1: douyin_favorite_action_request req) (api.post="/douyin/favorite/action/")
    douyin_favorite_list_response GetFavoriteList(1: douyin_favorite_list_request req) (api.get="/douyin/favorite/list/")
}