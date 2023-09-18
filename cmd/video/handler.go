package main

import (
	"context"
	"douyin/cmd/video/service"
	video "douyin/kitex_gen/video"
	"douyin/pkg/constant"
	"douyin/pkg/errno"
	"douyin/pkg/util"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/kitex/pkg/klog"
)

// VideoServiceImpl implements the last service interface defined in the IDL.
type VideoServiceImpl struct{}

// GetFeed implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) GetFeed(ctx context.Context, req *video.DouyinFeedRequest) (resp *video.DouyinFeedResponse, err error) {
	if err := req.IsValid(); err != nil {
		errNo := errno.UserRequestParameterError
		errNo.ErrMsg = err.Error()
		return nil, errNo
	}

	// 解析 Token
	var userID uint64
	if req.Token != nil {
		claim, _ := util.ParseToken(*req.Token)
		userID = claim.ID
	}
	return service.NewFeedService(ctx).GetFeed(req.LatestTime, userID)
}

// PublishAction implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) PublishAction(ctx context.Context, req *video.DouyinPublishActionRequest) (resp *video.DouyinPublishActionResponse, err error) {
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
	return service.NewPublishService(ctx).PublishAction(req.Title, req.Data, claim.ID)
}

// GetPublishVideos implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) GetPublishVideos(ctx context.Context, req *video.DouyinPublishListRequest) (resp *video.DouyinPublishListResponse, err error) {
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
	return service.NewPublishService(ctx).GetPublishVideos(claim.ID, uint64(req.UserId))
}

// FavoriteVideo implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) FavoriteVideo(ctx context.Context, req *video.DouyinFavoriteActionRequest) (resp *video.DouyinFavoriteActionResponse, err error) {
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

	if req.ActionType == constant.Favorite {
		resp, err = service.NewFavoriteService(ctx).FavoriteVideo(claim.ID, uint64(req.VideoId))
	} else if req.ActionType == constant.CancelFavorite {
		resp, err = service.NewFavoriteService(ctx).CancelFavoriteVideo(claim.ID, uint64(req.VideoId))
	} else {
		err = errno.UserRequestParameterError
		hlog.Error("handler.favorite_service.FavoriteVideo err:", err.Error())
	}
	return
}

// GetFavoriteList implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) GetFavoriteList(ctx context.Context, req *video.DouyinFavoriteListRequest) (resp *video.DouyinFavoriteListResponse, err error) {
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
	return service.NewFavoriteService(ctx).GetFavoriteList(claim.ID, uint64(req.UserId))
}
