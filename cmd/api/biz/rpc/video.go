package rpc

import (
	"context"
	"douyin/kitex_gen/video"
	"douyin/kitex_gen/video/videoservice"

	"douyin/pkg/constant"
	"douyin/pkg/mw"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/kitex-contrib/obs-opentelemetry/provider"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
	etcd "github.com/kitex-contrib/registry-etcd"
)

var videoClient videoservice.Client

func initVideo() {
	r, err := etcd.NewEtcdResolver([]string{constant.ETCDAddress})
	if err != nil {
		panic(err)
	}
	provider.NewOpenTelemetryProvider(
		provider.WithServiceName(constant.ApiServiceName),
		provider.WithExportEndpoint(constant.ExportEndpoint),
		provider.WithInsecure(),
	)
	c, err := videoservice.NewClient(
		constant.VideoServiceName,
		client.WithResolver(r),
		client.WithMuxConnection(1),
		client.WithMiddleware(mw.CommonMiddleware),
		client.WithInstanceMW(mw.ClientMiddleware),
		client.WithSuite(tracing.NewClientSuite()),
		client.WithClientBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: constant.ApiServiceName}),
	)
	if err != nil {
		panic(err)
	}
	videoClient = c
}

// Favorite

func FavoriteAction(ctx context.Context, req *video.DouyinFavoriteActionRequest) (*video.DouyinFavoriteActionResponse, error) {
	return videoClient.FavoriteVideo(ctx, req)
}

func GetFavoriteList(ctx context.Context, req *video.DouyinFavoriteListRequest) (*video.DouyinFavoriteListResponse, error) {
	return videoClient.GetFavoriteList(ctx, req)
}

// Feed

func GetFeed(ctx context.Context, req *video.DouyinFeedRequest) (*video.DouyinFeedResponse, error) {
	return videoClient.GetFeed(ctx, req)
}

// Publish

func PublishAction(ctx context.Context, req *video.DouyinPublishActionRequest) (*video.DouyinPublishActionResponse, error) {
	return videoClient.PublishAction(ctx, req)
}

func GetPublishVideos(ctx context.Context, req *video.DouyinPublishListRequest) (*video.DouyinPublishListResponse, error) {
	return videoClient.GetPublishVideos(ctx, req)
}
