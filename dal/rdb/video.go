package rdb

import (
	"douyin/pkg/constant"
	"douyin/pkg/global"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"strconv"
	"time"
)

// ------------ Redis ------------

// VideoInfo 视频固定的信息
type VideoInfo struct {
	PublishTime time.Time
	AuthorID    uint64
	PlayURL     string
	CoverURL    string
	Title       string
}

type VideoCount struct {
	FavoriteCount int64
	CommentCount  int64
}

// FavoriteVideoIDZSet 用户点赞视频ID 集合
type FavoriteVideoIDZSet struct {
	VideoID     uint64
	CreatedTime float64
}

// PublishVideoIDZSet 用户发布视频ID 集合
type PublishVideoIDZSet struct {
	VideoID     uint64
	CreatedTime float64
}

// LoadFavoriteVideoID 加载用户点赞视频列表到 Redis
func LoadFavoriteVideoID(userID uint64, videoIDs []uint64) error {
	key := constant.FavoriteVideoIDRedisSetPrefix + strconv.FormatUint(userID, 10)
	err := global.VideoRC.ZAdd(key).Err()
	if err != nil {
		hlog.Error("rdb.video.FavoriteVideo err:", err.Error())
		return err
	}
	return nil
}

// AddFavoriteVideoID 添加 Redis 用户点赞视频 ID
func AddFavoriteVideoID(userID, videoID uint64) error {
	key := constant.FavoriteVideoIDRedisSetPrefix + strconv.FormatUint(userID, 10)
	err := global.VideoRC.ZAdd(key).Err()
	if err != nil {
		hlog.Error("rdb.video.FavoriteVideo err:", err.Error())
		return err
	}
	return nil
}

// DelFavoriteVideoID 删除 Redis 用户点赞视频 ID
func DelFavoriteVideoID(userID, videoID uint64) error {
	key := constant.FavoriteVideoIDRedisSetPrefix + strconv.FormatUint(userID, 10)
	err := global.VideoRC.SRem(key, videoID).Err()
	if err != nil {
		hlog.Error("rdb.video.DelFavoriteVideo err:", err.Error())
		return err
	}
	return nil
}

// SetAuthorID 通过视频ID获取
func SetAuthorID(authorID, videoID uint64) error {
	key := constant.VideoAuthorIDRedisPrefix + strconv.FormatUint(videoID, 10)
	return global.VideoRC.Set(key, authorID, constant.VideoAuthorIDExpiration).Err()
}

// GetAuthorID 通过视频ID获取作者ID
func GetAuthorID(videoID uint64) (uint64, error) {
	key := constant.VideoAuthorIDRedisPrefix + strconv.FormatUint(videoID, 10)
	authorIDString, err := global.VideoRC.Get(key).Result()
	if err != nil {
		return 0, nil
	}
	return strconv.ParseUint(authorIDString, 10, 64)
}
