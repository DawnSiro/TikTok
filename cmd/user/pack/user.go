package pack

import (
	"douyin/dal/model"
	"douyin/kitex_gen/user"
	"douyin/pkg/errno"

	"github.com/cloudwego/kitex/pkg/klog"
)

func UserInfo(u *model.User, isFollow bool) *user.UserInfo {
	if u == nil {
		klog.Error("pack.user.UserInfo err:", errno.ServiceError)
		return nil
	}
	return &user.UserInfo{
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
}

func User(u *model.User, isFollow bool) *user.User {
	if u == nil {
		klog.Error("pack.user.User err:", errno.ServiceError)
		return nil
	}
	followCount := u.FollowingCount
	followerCount := u.FollowerCount
	return &user.User{
		Id:            int64(u.ID),
		Name:          u.Username,
		FollowCount:   &followCount,
		FollowerCount: &followerCount,
		IsFollow:      isFollow,
		Avatar:        u.Avatar,
	}
}

func FriendUser(u *model.User, isFollow bool, messageContent string, msgType uint8) *user.FriendUser {
	if u == nil {
		klog.Error("pack.user.UserInfo err:", errno.ServiceError)
		return nil
	}
	return &user.FriendUser{
		Id:       int64(u.ID),
		Name:     u.Username,
		IsFollow: isFollow,
		Avatar:   u.Avatar,
		Message:  &messageContent,
		MsgType:  int8(msgType),
	}
}

func RelationData(data *model.RelationUserData) *user.User {
	if data == nil {
		return nil
	}
	followCount := int64(data.FollowingCount)
	followerCount := int64(data.FollowerCount)
	return &user.User{
		Id:            int64(data.UID),
		Name:          data.Username,
		FollowCount:   &followCount,
		FollowerCount: &followerCount,
		IsFollow:      data.IsFollow,
		Avatar:        data.Avatar,
	}
}

func RelationDataList(dataList []*model.RelationUserData) []*user.User {
	res := make([]*user.User, 0, len(dataList))
	for i := 0; i < len(dataList); i++ {
		res = append(res, RelationData(dataList[i]))
	}
	return res
}
