package message

import (
	"context"
)

type QueryV4Request struct {
	QueryType string `json:"queryType"` // 查询类型。支持多个组合，英文逗号间隔。例如1,2,3
	DelFlag   int    `json:"delFlag"`   // 1：成功消费删除消息,0：或者不传递不删除消息
}

// 8.2 消息通知查询接口 v4 版本

// 每次查询最多返回100条数据
// 数据最多保留60天
// 查询成功默认消息已经处理

func (m *MessageService) QueryV4(ctx context.Context, token string, req *QueryV4Request) (resp *QueryResponse, err error) {
	err = m.http.DoAuthQuery(ctx, "GET", "/open/api/message/query/notify/statusQueryV4", token, req, &resp)

	return resp, err
}
