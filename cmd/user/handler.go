package main

import (
	"context"
	"douyin/pkg/constant"
	"github.com/cloudwego/kitex/pkg/klog"

	"douyin/cmd/user/service"
	"douyin/kitex_gen/user"
	"douyin/pkg/errno"
	"douyin/pkg/util"
)

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct{}

// Register implements the UserServiceImpl interface.
func (s *UserServiceImpl) Register(ctx context.Context, req *user.DouyinUserRegisterRequest) (*user.DouyinUserRegisterResponse, error) {
	if err := req.IsValid(); err != nil {
		errNo := errno.UserRequestParameterError
		errNo.ErrMsg = err.Error()
		return nil, errNo
	}
	return service.NewUserService(ctx).Register(req.Username, req.Password)
}

// Login implements the UserServiceImpl interface.
func (s *UserServiceImpl) Login(ctx context.Context, req *user.DouyinUserLoginRequest) (*user.DouyinUserLoginResponse, error) {
	if err := req.IsValid(); err != nil {
		errNo := errno.UserRequestParameterError
		errNo.ErrMsg = err.Error()
		return nil, errNo
	}
	return service.NewUserService(ctx).Login(req.Username, req.Password)
}

// GetUserInfo implements the UserServiceImpl interface.
func (s *UserServiceImpl) GetUserInfo(ctx context.Context, req *user.DouyinUserRequest) (*user.DouyinUserResponse, error) {
	if err := req.IsValid(); err != nil {
		errNo := errno.UserRequestParameterError
		errNo.ErrMsg = err.Error()
		return nil, errNo
	}

	claim, err := util.ParseToken(req.Token)
	if err != nil {
		return nil, err
	}

	return service.NewUserService(ctx).GetUserInfo(claim.ID, uint64(req.UserId))
}

// Follow implements the UserServiceImpl interface.
func (s *UserServiceImpl) Follow(ctx context.Context, req *user.DouyinRelationActionRequest) (resp *user.DouyinRelationActionResponse, err error) {
	if err := req.IsValid(); err != nil {
		errNo := errno.UserRequestParameterError
		errNo.ErrMsg = err.Error()
		return nil, errNo
	}

	// 解析 Token，再验一次权
	claim, err := util.ParseToken(req.Token)
	if err != nil {
		klog.Error("handler.handler.CommentAction err:", err.Error())
		return nil, err
	}
	if req.ActionType == constant.Follow {
		resp, err = service.NewRelationService(ctx).Follow(claim.ID, uint64(req.ToUserId))
	} else if req.ActionType == constant.CancelFollow {
		resp, err = service.NewRelationService(ctx).CancelFollow(claim.ID, uint64(req.ToUserId))
	}
	err = errno.UserRequestParameterError
	klog.Error("handler.relation_service.Follow err:", err.Error())
	return
}

// GetFollowList implements the UserServiceImpl interface.
func (s *UserServiceImpl) GetFollowList(ctx context.Context, req *user.DouyinRelationFollowListRequest) (resp *user.DouyinRelationFollowListResponse, err error) {
	if err := req.IsValid(); err != nil {
		errNo := errno.UserRequestParameterError
		errNo.ErrMsg = err.Error()
		return nil, errNo
	}

	// 解析 Token，再验一次权
	claim, err := util.ParseToken(req.Token)
	if err != nil {
		klog.Error("handler.handler.CommentAction err:", err.Error())
		return nil, err
	}
	return service.NewRelationService(ctx).GetFollowList(claim.ID, uint64(req.UserId))
}

// GetFollowerList implements the UserServiceImpl interface.
func (s *UserServiceImpl) GetFollowerList(ctx context.Context, req *user.DouyinRelationFollowerListRequest) (resp *user.DouyinRelationFollowerListResponse, err error) {
	if err := req.IsValid(); err != nil {
		errNo := errno.UserRequestParameterError
		errNo.ErrMsg = err.Error()
		return nil, errNo
	}

	// 解析 Token，再验一次权
	claim, err := util.ParseToken(req.Token)
	if err != nil {
		klog.Error("handler.handler.CommentAction err:", err.Error())
		return nil, err
	}
	return service.NewRelationService(ctx).GetFollowerList(claim.ID, uint64(req.UserId))
}

// GetFriendList implements the UserServiceImpl interface.
func (s *UserServiceImpl) GetFriendList(ctx context.Context, req *user.DouyinRelationFriendListRequest) (resp *user.DouyinRelationFriendListResponse, err error) {
	if err := req.IsValid(); err != nil {
		errNo := errno.UserRequestParameterError
		errNo.ErrMsg = err.Error()
		return nil, errNo
	}

	// 解析 Token，再验一次权
	claim, err := util.ParseToken(req.Token)
	if err != nil {
		klog.Error("handler.handler.GetFriendList err:", err.Error())
		return nil, err
	}
	if claim.ID == uint64(req.UserId) {
		return nil, errno.UserRequestParameterError
	}
	return service.NewRelationService(ctx).GetFriendList(claim.ID)
}
