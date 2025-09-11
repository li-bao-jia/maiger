package message

import (
	"context"
)

type DeleteRequest struct {
	Ids []string `json:"iidd"` // 消息ID，支持批量删除，英文逗号间隔，最大100个
}

type DeleteResponse struct {
	Data    bool   `json:"data"`
	Code    string `json:"code"`    // 返回编码
	Message string `json:"message"` // 返回说明
}

// 9.1 删除推送消息

func (m *MessageService) Delete(ctx context.Context, token string, req *DeleteRequest) (resp *DeleteResponse, err error) {
	err = m.http.DoAuthForm(ctx, "POST", "/open/api/message/query/notify/batchDel", token, req, &resp)

	return resp, err
}
