// Code generated by hertz generator.

package api

import (
	"bytes"
	"context"
	"io"
	"mime/multipart"

	"douyin/cmd/api/biz/model/api"
	"douyin/cmd/api/biz/rpc"
	"douyin/kitex_gen/publish"
	"douyin/pkg/errno"
	"douyin/pkg/global"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// PublishAction .
// @router /douyin/publish/action/ [POST]
func PublishAction(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.DouyinPublishActionRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.JSON(consts.StatusOK, &api.DouyinResponse{
			StatusCode: errno.UserRequestParameterError.ErrCode,
			StatusMsg:  err.Error(),
		})
		return
	}

	hlog.Info("handler.publish_service.PublishAction Request:", req)
	fileHeader, err := c.FormFile("data")
	if err != nil {
		msg := err.Error()
		c.JSON(consts.StatusOK, &api.DouyinPublishActionResponse{
			StatusCode: errno.UserRequestParameterError.ErrCode,
			StatusMsg:  &msg,
		})
		return
	}

	hlog.Info("handler.publish_service.PublishAction Filename:", fileHeader.Filename)
	file, err := fileHeader.Open()
	defer func(file multipart.File) {
		err := file.Close()
		if err != nil {
			hlog.Error("handler.publish_service.PublishAction err:", err)
			errNo := errno.ConvertErr(err)
			c.JSON(consts.StatusOK, &api.DouyinPublishActionResponse{
				StatusCode: errNo.ErrCode,
				StatusMsg:  &errNo.ErrMsg,
			})
			return
		}
	}(file)
	// 将文件转化为字节流
	buf := bytes.NewBuffer(nil)
	if _, err := io.Copy(buf, file); err != nil {
		errNo := errno.ConvertErr(err)
		c.JSON(consts.StatusOK, &api.DouyinPublishActionResponse{
			StatusCode: errNo.ErrCode,
			StatusMsg:  &errNo.ErrMsg,
		})
		return
	}

	userID := c.GetUint64(global.Config.JWTConfig.IdentityKey)
	hlog.Info("handler.feed_service.GetFeed GetUserID:", userID)
	resp, err := rpc.PublishAction(context.Background(), &publish.DouyinPublishActionRequest{
		Token: req.Token,
		Data:  buf.Bytes(),
		Title: req.Title,
	})
	if err != nil {
		errNo := errno.ConvertErr(err)
		c.JSON(consts.StatusOK, &api.DouyinPublishActionResponse{
			StatusCode: errNo.ErrCode,
			StatusMsg:  &errNo.ErrMsg,
		})
		return
	}

	c.JSON(consts.StatusOK, resp)
}

// GetPublishVideos .
// @router /douyin/publish/list/ [GET]
func GetPublishVideos(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.DouyinPublishListRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.JSON(consts.StatusOK, &api.DouyinResponse{
			StatusCode: errno.UserRequestParameterError.ErrCode,
			StatusMsg:  err.Error(),
		})
		return
	}

	hlog.Info("handler.publish_service.GetPublishVideos Request:", req)
	resp, err := rpc.GetPublishVideos(context.Background(), &publish.DouyinPublishListRequest{
		UserId: req.UserID,
		Token:  req.Token,
	})
	if err != nil {
		errNo := errno.ConvertErr(err)
		c.JSON(consts.StatusOK, &api.DouyinPublishListResponse{
			StatusCode: errNo.ErrCode,
			StatusMsg:  &errNo.ErrMsg,
		})
		return
	}

	c.JSON(consts.StatusOK, resp)
}