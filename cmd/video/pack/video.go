package pack

import (
	"douyin/dal/model"
	"douyin/kitex_gen/video"
	"douyin/pkg/errno"

	"github.com/cloudwego/kitex/pkg/klog"
)

func Video(v *model.Video, u *model.User, isFollow, isFavorite bool) *video.Video {
	if v == nil || u == nil {
		klog.Error("pack.video.Video err:", errno.ServiceError)
		return nil
	}
	author := &video.UserInfo{
		Id:              int64(u.ID),
		Name:            u.Username,
		FollowCount:     u.FollowingCount,
		FollowerCount:   u.FollowerCount,
		IsFollow:        isFollow,
		Avatar:          u.Avatar,
		BackgroundImage: u.BackgroundImage,
		Signature:       u.Signature,
		TotalFavorited:  u.TotalFavorited,
		WorkCount:       u.WorkCount,
		FavoriteCount:   u.FavoriteCount,
	}
	return &video.Video{
		Id:            int64(v.ID),
		Author:        author,
		PlayUrl:       v.PlayURL,
		CoverUrl:      v.CoverURL,
		FavoriteCount: v.FavoriteCount,
		CommentCount:  v.CommentCount,
		IsFavorite:    isFavorite,
		Title:         v.Title,
	}
}

func VideoData(data *model.VideoData) *video.Video {
	if data == nil {
		return nil
	}
	followCount := data.FollowCount
	followerCount := data.FollowerCount
	author := &video.UserInfo{
		Id:              int64(data.UID),
		Name:            data.Username,
		FollowCount:     followCount,
		FollowerCount:   followerCount,
		IsFollow:        data.IsFollow,
		Avatar:          data.Avatar,
		BackgroundImage: data.BackgroundImage,
		Signature:       data.Signature,
		TotalFavorited:  data.TotalFavorited,
		WorkCount:       data.WorkCount,
		FavoriteCount:   data.UserFavoriteCount,
	}
	return &video.Video{
		Id:            int64(data.VID),
		Author:        author,
		PlayUrl:       data.PlayURL,
		CoverUrl:      data.CoverURL,
		FavoriteCount: data.FavoriteCount,
		CommentCount:  data.CommentCount,
		IsFavorite:    data.IsFavorite,
		Title:         data.Title,
	}
}

func VideoDataList(dataList []*model.VideoData) []*video.Video {
	res := make([]*video.Video, 0, len(dataList))
	for i := 0; i < len(dataList); i++ {
		res = append(res, VideoData(dataList[i]))
	}
	return res
}
