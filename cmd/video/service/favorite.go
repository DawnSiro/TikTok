package service

import (
	"context"
	"douyin/kitex_gen/video"
	"douyin/pkg/constant"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"strconv"
	"strings"

	"douyin/cmd/video/pack"
	"douyin/dal/db"
	"douyin/pkg/errno"
	"douyin/pkg/global"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/go-redis/redis"
)

type FavoriteService struct {
	ctx context.Context
}

// NewFavoriteService new favoriteService
func NewFavoriteService(ctx context.Context) *FavoriteService {
	return &FavoriteService{
		ctx: ctx,
	}
}

// FavoriteVideo this is a func for add Favorite or reduce Favorite
func (s *FavoriteService) FavoriteVideo(userID, videoID uint64) (*video.DouyinFavoriteActionResponse, error) {
	var builder strings.Builder
	builder.WriteString(strconv.FormatUint(videoID, 10))
	builder.WriteString("_video_like")
	videoLikeKey := builder.String()

	var builder1 strings.Builder
	builder1.WriteString(strconv.FormatUint(userID, 10))
	builder1.WriteString("_user_like_count")
	videoLikeCountKey := builder1.String()

	userVideoLikeCount, err := global.VideoRC.Get(videoLikeCountKey).Result()
	var videoLikeCountInt int
	if err == redis.Nil {
		global.VideoRC.Set(videoLikeCountKey, "0", constant.VideoLikeLimitTime)
	} else {
		videoLikeCountInt, _ = strconv.Atoi(userVideoLikeCount)
		if videoLikeCountInt >= constant.VideoLikeLimit {
			return &video.DouyinFavoriteActionResponse{
				StatusCode: errno.VideoLikeLimitError.ErrCode,
				StatusMsg:  &errno.VideoLikeLimitError.ErrMsg,
			}, nil
		}
	}

	likeCount, err := global.VideoRC.Get(videoLikeKey).Result()
	if err == redis.Nil {
		likeInt64, err := db.SelectVideoFavoriteCountByVideoID(s.ctx, videoID)
		if err != nil {
			hlog.Error("service.favorite.FavoriteVideo err:", err.Error())
			return nil, err
		}
		global.VideoRC.Set(videoLikeKey, likeInt64, 0)
	}
	var likeUint64 uint64
	if likeCount != "" {
		likeUint64, err = strconv.ParseUint(likeCount, 10, 64)
		if err != nil {
			hlog.Error("service.favorite.FavoriteVideo err:", err.Error())
			return nil, err
		}
	}

	err = db.FavoriteVideo(s.ctx, userID, videoID)
	if err != nil {
		hlog.Error("service.favorite.FavoriteVideo err:", err.Error())
		return nil, err
	}
	// 如果 DB 层事务回滚了，err 就不为 nil，Redis 里的数据就不会更新
	global.VideoRC.Set(videoLikeKey, likeUint64+1, 0)
	// 更新单位时间内的点赞数量
	userVideoLikeCountTime, err := global.VideoRC.TTL(videoLikeCountKey).Result()
	global.VideoRC.Set(videoLikeCountKey, videoLikeCountInt+1, userVideoLikeCountTime)

	return &video.DouyinFavoriteActionResponse{
		StatusCode: errno.Success.ErrCode,
	}, nil
}

func (s *FavoriteService) CancelFavoriteVideo(userID, videoID uint64) (*video.DouyinFavoriteActionResponse, error) {
	var builder strings.Builder
	builder.WriteString(strconv.FormatUint(videoID, 10))
	builder.WriteString("_video_like")
	videoLikeKey := builder.String()

	likeCount, err := global.VideoRC.Get(videoLikeKey).Result()
	if err == redis.Nil {
		likeInt64, err := db.SelectVideoFavoriteCountByVideoID(s.ctx, videoID)
		if err != nil {
			klog.Error("service.favorite.CancelFavoriteVideo err:", err.Error())
			return nil, err
		}
		global.VideoRC.Set(videoLikeKey, likeInt64, 0)
	}

	var likeUint64 uint64
	if likeCount != "" {
		likeUint64, err = strconv.ParseUint(likeCount, 10, 64)
		if err != nil {
			klog.Error("service.video.CancelFavoriteVideo err:", err.Error())
			return nil, err
		}
	}

	err = db.CancelFavoriteVideo(s.ctx, userID, videoID)
	if err != nil {
		klog.Error("service.favorite.CancelFavoriteVideo err:", err.Error())
		return nil, err
	}
	// 如果 DB 层事务回滚了，err 就不为 nil，Redis 里的数据就不会更新
	global.VideoRC.Set(videoLikeKey, likeUint64-1, 0)

	return &video.DouyinFavoriteActionResponse{
		StatusCode: errno.Success.ErrCode,
	}, nil
}

func (s *FavoriteService) GetFavoriteList(userID, selectUserID uint64) (*video.DouyinFavoriteListResponse, error) {
	videoList, err := db.SelectFavoriteVideoDataListByUserID(s.ctx, userID, selectUserID)
	if err != nil {
		return nil, err
	}
	return &video.DouyinFavoriteListResponse{
		StatusCode: errno.Success.ErrCode,
		VideoList:  pack.VideoDataList(videoList),
	}, nil
}
