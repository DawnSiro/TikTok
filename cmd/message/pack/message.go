package pack

import (
	"douyin/dal/model"
	"douyin/kitex_gen/message"
	"douyin/pkg/errno"

	"github.com/cloudwego/kitex/pkg/klog"
)

func Messages(ms []*model.Message) []*message.Message {
	if ms == nil {
		klog.Error("pack.message.Messages err:", errno.ServiceError)
		return nil
	}
	res := make([]*message.Message, 0)
	for i := 0; i < len(ms); i++ {
		res = append(res, Message(ms[i]))
	}
	return res
}

func Message(m *model.Message) *message.Message {
	if m == nil {
		klog.Error("pack.message.Messages err:", errno.ServiceError)
		return nil
	}
	createTime := m.CreatedTime.UnixMilli()
	return &message.Message{
		Id:         int64(m.ID),
		ToUserId:   int64(m.ToUserID),
		FromUserId: int64(m.FromUserID),
		Content:    m.Content,
		CreateTime: &createTime,
	}
}
