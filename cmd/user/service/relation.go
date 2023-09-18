package service

import (
	"context"
	"douyin/cmd/user/pack"
	"douyin/dal/db"
	"douyin/kitex_gen/user"
	"douyin/pkg/errno"

	"github.com/cloudwego/kitex/pkg/klog"
)

type RelationService struct {
	ctx context.Context
}

func NewRelationService(ctx context.Context) *RelationService {
	return &RelationService{
		ctx: ctx,
	}
}

func (s *RelationService) Follow(userID, toUserID uint64) (*user.DouyinRelationActionResponse, error) {
	if userID == toUserID {
		errNo := errno.UserRequestParameterError
		errNo.ErrMsg = "不能自己关注自己哦"
		klog.Error("service.relation.Follow err:", errNo.Error())
		return nil, errNo
	}
	isFollow := db.IsFollow(s.ctx, userID, toUserID)
	if isFollow {
		klog.Error("service.relation.Follow err:", errno.RepeatOperationError)
		return nil, errno.RepeatOperationError
	}

	//关注操作
	err := db.Follow(s.ctx, userID, toUserID)
	if err != nil {
		klog.Error("service.relation.Follow err:", err.Error())
		return nil, err
	}
	return &user.DouyinRelationActionResponse{
		StatusCode: errno.Success.ErrCode,
	}, nil
}

func (s *RelationService) CancelFollow(userID, toUserID uint64) (*user.DouyinRelationActionResponse, error) {
	if userID == toUserID {
		errNo := errno.UserRequestParameterError
		errNo.ErrMsg = "不能自己取关自己哦"
		klog.Error("service.relation.CancelFollow err:", errNo.Error())
		return nil, errNo
	}
	//取消关注
	err := db.CancelFollow(s.ctx, userID, toUserID)
	if err != nil {
		klog.Error("service.relation.CancelFollow err:", err.Error())
		return nil, err
	}
	return &user.DouyinRelationActionResponse{
		StatusCode: errno.Success.ErrCode,
	}, nil
}

// GetFollowList
// userID 为发送请求的用户ID，从 Token 里取到
// selectUserID 为需要查询的用户的ID，做为请求参数传递
func (s *RelationService) GetFollowList(userID, selectUserID uint64) (*user.DouyinRelationFollowListResponse, error) {
	relationDataList, err := db.SelectFollowDataListByUserID(s.ctx, userID)
	if err != nil {
		klog.Error("service.relation.GetFollowList err:", err.Error())
		return nil, err
	}

	return &user.DouyinRelationFollowListResponse{
		StatusCode: errno.Success.ErrCode,
		UserList:   pack.RelationDataList(relationDataList),
	}, nil
}

func (s *RelationService) GetFollowerList(userID, selectUserID uint64) (*user.DouyinRelationFollowerListResponse, error) {
	relationDataList, err := db.SelectFollowerDataListByUserID(s.ctx, userID)
	if err != nil {
		klog.Error("service.relation.GetFollowList err:", err.Error())
		return nil, err
	}

	return &user.DouyinRelationFollowerListResponse{
		StatusCode: errno.Success.ErrCode,
		UserList:   pack.RelationDataList(relationDataList),
	}, nil
}

func (s *RelationService) GetFriendList(userID uint64) (*user.DouyinRelationFriendListResponse, error) {
	userList, err := db.GetFriendList(s.ctx, userID)
	if err != nil {
		klog.Error("service.relation.GetFollowerList err:", err.Error())
		return nil, err
	}

	friendUserList := make([]*user.FriendUser, 0, len(userList))
	for _, u := range userList {
		msg, err := db.GetLatestMsg(s.ctx, userID, u.ID)
		if err != nil {
			klog.Error("service.relation.GetFollowerList err:", err.Error())
			return nil, err
		}
		friendUserList = append(friendUserList, pack.FriendUser(u, db.IsFollow(s.ctx, userID, u.ID), msg.Content, msg.MsgType))
	}

	return &user.DouyinRelationFriendListResponse{
		StatusCode: errno.Success.ErrCode,
		UserList:   friendUserList,
	}, nil
}
