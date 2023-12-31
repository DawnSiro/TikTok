// Code generated by Kitex v0.7.1. DO NOT EDIT.

package userservice

import (
	"context"
	user "douyin/kitex_gen/user"
	client "github.com/cloudwego/kitex/client"
	callopt "github.com/cloudwego/kitex/client/callopt"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	Register(ctx context.Context, req *user.DouyinUserRegisterRequest, callOptions ...callopt.Option) (r *user.DouyinUserRegisterResponse, err error)
	Login(ctx context.Context, req *user.DouyinUserLoginRequest, callOptions ...callopt.Option) (r *user.DouyinUserLoginResponse, err error)
	GetUserInfo(ctx context.Context, req *user.DouyinUserRequest, callOptions ...callopt.Option) (r *user.DouyinUserResponse, err error)
	Follow(ctx context.Context, req *user.DouyinRelationActionRequest, callOptions ...callopt.Option) (r *user.DouyinRelationActionResponse, err error)
	GetFollowList(ctx context.Context, req *user.DouyinRelationFollowListRequest, callOptions ...callopt.Option) (r *user.DouyinRelationFollowListResponse, err error)
	GetFollowerList(ctx context.Context, req *user.DouyinRelationFollowerListRequest, callOptions ...callopt.Option) (r *user.DouyinRelationFollowerListResponse, err error)
	GetFriendList(ctx context.Context, req *user.DouyinRelationFriendListRequest, callOptions ...callopt.Option) (r *user.DouyinRelationFriendListResponse, err error)
}

// NewClient creates a client for the service defined in IDL.
func NewClient(destService string, opts ...client.Option) (Client, error) {
	var options []client.Option
	options = append(options, client.WithDestService(destService))

	options = append(options, opts...)

	kc, err := client.NewClient(serviceInfo(), options...)
	if err != nil {
		return nil, err
	}
	return &kUserServiceClient{
		kClient: newServiceClient(kc),
	}, nil
}

// MustNewClient creates a client for the service defined in IDL. It panics if any error occurs.
func MustNewClient(destService string, opts ...client.Option) Client {
	kc, err := NewClient(destService, opts...)
	if err != nil {
		panic(err)
	}
	return kc
}

type kUserServiceClient struct {
	*kClient
}

func (p *kUserServiceClient) Register(ctx context.Context, req *user.DouyinUserRegisterRequest, callOptions ...callopt.Option) (r *user.DouyinUserRegisterResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.Register(ctx, req)
}

func (p *kUserServiceClient) Login(ctx context.Context, req *user.DouyinUserLoginRequest, callOptions ...callopt.Option) (r *user.DouyinUserLoginResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.Login(ctx, req)
}

func (p *kUserServiceClient) GetUserInfo(ctx context.Context, req *user.DouyinUserRequest, callOptions ...callopt.Option) (r *user.DouyinUserResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetUserInfo(ctx, req)
}

func (p *kUserServiceClient) Follow(ctx context.Context, req *user.DouyinRelationActionRequest, callOptions ...callopt.Option) (r *user.DouyinRelationActionResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.Follow(ctx, req)
}

func (p *kUserServiceClient) GetFollowList(ctx context.Context, req *user.DouyinRelationFollowListRequest, callOptions ...callopt.Option) (r *user.DouyinRelationFollowListResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetFollowList(ctx, req)
}

func (p *kUserServiceClient) GetFollowerList(ctx context.Context, req *user.DouyinRelationFollowerListRequest, callOptions ...callopt.Option) (r *user.DouyinRelationFollowerListResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetFollowerList(ctx, req)
}

func (p *kUserServiceClient) GetFriendList(ctx context.Context, req *user.DouyinRelationFriendListRequest, callOptions ...callopt.Option) (r *user.DouyinRelationFriendListResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetFriendList(ctx, req)
}
