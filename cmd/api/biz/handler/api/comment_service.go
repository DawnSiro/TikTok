// Code generated by hertz generator.

package api

import (
	"context"
	api "douyin/cmd/api/biz/model/api"
	"douyin/cmd/api/biz/rpc"
	"douyin/kitex_gen/comment"
	"douyin/pkg/errno"
	"douyin/pkg/global"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// CommentAction .
// @router /douyin/comment/action/ [POST]
func CommentAction(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.DouyinCommentActionRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.JSON(consts.StatusOK, &api.DouyinResponse{
			StatusCode: errno.UserRequestParameterError.ErrCode,
			StatusMsg:  err.Error(),
		})
		return
	}

	hlog.Info("handler.comment_service.CommentAction Request:", req)
	userID := c.GetUint64(global.Config.JWTConfig.IdentityKey)
	hlog.Info("handler.comment_service.CommentAction GetUserID:", userID)
	var resp *comment.DouyinCommentActionResponse
	// 这里注意走 ActionType 对应的逻辑的时候要注意判断相关字段是否为空
	resp, err = rpc.CommentAction(context.Background(), &comment.DouyinCommentActionRequest{
		Token:       req.Token,
		VideoId:     req.VideoID,
		ActionType:  req.ActionType,
		CommentText: req.CommentText,
		CommentId:   req.CommentID,
	})
	if err != nil {
		errNo := errno.ConvertErr(err)
		c.JSON(consts.StatusOK, &api.DouyinCommentActionResponse{
			StatusCode: errNo.ErrCode,
			StatusMsg:  &errNo.ErrMsg,
		})
		return
	}

	c.JSON(consts.StatusOK, resp)
}

// GetCommentList .
// @router /douyin/comment/list/ [GET]
func GetCommentList(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.DouyinCommentListRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.JSON(consts.StatusOK, &api.DouyinResponse{
			StatusCode: errno.UserRequestParameterError.ErrCode,
			StatusMsg:  err.Error(),
		})
		return
	}

	hlog.Info("handler.comment_service.GetCommentList Request:", req)
	userID := c.GetUint64(global.Config.JWTConfig.IdentityKey)
	hlog.Info("handler.comment_service.GetCommentList GetUserID:", userID)
	resp, err := rpc.GetCommentList(context.Background(), &comment.DouyinCommentListRequest{
		Token:   req.Token,
		VideoId: req.VideoID,
	})
	if err != nil {
		errNo := errno.ConvertErr(err)
		c.JSON(consts.StatusOK, &api.DouyinCommentListResponse{
			StatusCode: errNo.ErrCode,
			StatusMsg:  &errNo.ErrMsg,
		})
		return
	}

	c.JSON(consts.StatusOK, resp)
}
