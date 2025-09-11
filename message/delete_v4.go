package message

import (
	"context"
)

type DeleteV4Request struct {
	Ids []string `json:"ids"` // 消息ID，支持批量删除，英文逗号间隔，最大100个
}

// 9.2 删除推送消息 v4 版本

func (m *MessageService) DeleteV4(ctx context.Context, token string, req *DeleteV4Request) (resp *DeleteResponse, err error) {
	err = m.http.DoAuthJSON(ctx, "POST", "/open/api/message/query/notify/batchDelV4", token, req, &resp)

	return resp, err
}
