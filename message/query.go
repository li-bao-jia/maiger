package message

import (
	"context"
)

// 0		商品上下架变更
// 1		商品库存变更
// 2		商品基本信息变更
// 3		新增商品消息
// 4		商品价格变更
// 5		订单已发货变更
// 6		订单已取消变更
// 7		订单已完成变更
// 8		售后-客服已审核变更
// 9		售后-客诉人已退件变更
// 10		售后-供应商已收货变更
// 11		售后-供应商已发货变更
// 12		售后-客诉人已收件变更
// 13		售后-物流信息更新
// 14		售后-售后已退款变更
// 15		售后-退款已审核变更
// 16		售后-售后已取消变更
// 17		售后-售后已完成变更
// 18		售后-创建售后单通知变更
// 19		售后-售补偿售后单已审核变更
// 20		售后-补寄审核

type QueryRequest struct {
	QueryType string `json:"queryType"` // 查询类型。支持多个组合，英文逗号间隔。例如1,2,3
}

type QueryResponse struct {
	Data     QueryResult `json:"data"`
	Code     string      `json:"code"`     // 返回编码
	Message  string      `json:"message"`  // 返回说明
	DataSize int         `json:"dataSize"` // 当前类型的数据量大小
}

type QueryResult struct {
	MaxUpdateTime int            `json:"maxUpdateTime"`
	Contents      []QueryContent `json:"contents"`
}

type QueryContent struct {
	// 公共字段
	QueryType int    `json:"queryType"` // 消息类型枚举值
	IIDD      string `json:"iidd"`      // 主键ID
	ClientID  string `json:"clientId"`  // 客户端id

	// 商品相关
	ThirdSkuId string  `json:"thirdSkuId,omitempty"` // 商品skuId
	GoodsId    string  `json:"goodsId,omitempty"`    // 商品goodsId
	Type       int     `json:"type,omitempty"`       // 上下架状态 0上架 1下架
	SkuStock   int     `json:"skuStock,omitempty"`   // 库存变更后的数量
	SellPrice  float64 `json:"sellPrice,omitempty"`  // 销售价
	GoodsPrice float64 `json:"goodsPrice,omitempty"` // 指导价

	// 订单相关
	OrderSn  string `json:"orderSn,omitempty"`  // 订单编号
	SkuId    string `json:"skuId,omitempty"`    // 订单中的skuId
	Quantity int    `json:"quantity,omitempty"` // 发货数量

	// 售后相关
	AfterId  string   `json:"afterId,omitempty"`  // 售后ID
	Status   int      `json:"status,omitempty"`   // 审核状态 (0通过 1拒绝)
	Reason   string   `json:"reason,omitempty"`   // 拒绝原因
	ImageUrl string   `json:"imageUrl,omitempty"` // 图片证据（逗号分隔）
	Images   []string `json:"-"`                  // 转换后的图片数组（业务可用）
}

// 8.1 消息通知查询接口

// 每次查询最多返回100条数据
// 数据最多保留60天
// 查询成功默认消息已经处理

func (m *MessageService) Query(ctx context.Context, token string, req *QueryRequest) (resp *QueryResponse, err error) {
	err = m.http.DoAuthJSON(ctx, "POST", "/open/api/message/query/notify/statusQuery", token, req, &resp)

	return resp, err
}
