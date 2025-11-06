<div align=center>
  <p align="center"> Maiger SDK for Go</p>
  <p align="center">Maiger SDK 用于与 Maiger 开放平台 API 交互</p>
</div>

<div align=center>
  <p align="center">
    <a href="https://github.com/li-bao-jia">
      <img src="https://img.shields.io/badge/go-1.21.8-blue" alt="Build Status">
    </a>
    <a href="https://github.com/li-bao-jia">
      <img src="https://img.shields.io/github/license/li-bao-jia/maiger" alt="License">
    </a>
  </p>
</div>

## 项目概述

- 初衷：没有SDK，那就封装一个，减少重复工作
- 设计：模块化，每个模块对应一个功能模块
- 特点：完整的请求结构体、返回结构体，减少API封装请求、结构参数比对工作
- 功能：用于对接迈戈供应链平台API，提供了完整的商品、订单、售后、消息通知等功能模块


## 安装 Installation

```
go get github.com/li-bao-jia/maiger@latest
```

## 快速开始 Quick Start

### 初始化客户端

```go
package main

import (
    "context"
	
    "github.com/li-bao-jia/maiger"
)

func main() {
    // 创建客户端实例 * 替换为实际的 API 基础 URL 
    client := maiger.NewClient("https://api.maiger.com")

    // 调用链条 和 外部控制
    ctx := context.Background()
	
    // 获取访问令牌
    tokenResp, err := client.Token.GetToken(ctx, "your_client_id", "your_client_secret")
    if err != nil {
        panic(err)
    }
    
    token := tokenResp.AccessToken
    
    // 使用 token 调用其他 API
    // ...
}
```

## 服务模块

SDK 包含以下主要服务模块：

- **Token** - 认证服务
- **Product** - 商品服务  
- **Order** - 订单服务
- **AfterSale** - 售后服务
- **Message** - 消息通知服务

预留 定制商城专属服务

- **User** - 用户服务
- **Activity** - 活动服务
- **Coupon** - 优惠券服务


## 备注

因为权限影响，着手开发这个这个SDK时，没有定制商城专属服务权限，所以没有封装定制商城专属服务。这里仅预留了服务模块，如果你有需求，可以通过联系方式联系我，增加定制商城专属服务。

## 详细 API 文档

### 1. Token 服务 (client.Token)

用于获取和刷新访问令牌。

#### 方法列表

| 方法                                                        | 描述     |
|-----------------------------------------------------------|--------|
| `GetToken(ctx, clientId, clientSecret)`                   | 获取访问令牌 |
| `RefreshToken(ctx, clientId, clientSecret, refreshToken)` | 刷新访问令牌 |

#### 使用示例

```go
// 获取访问令牌
tokenResp, err := client.Token.GetToken(ctx, "your_client_id", "your_client_secret")
if err != nil {
    return err
}

// 刷新令牌
refreshResp, err := client.Token.RefreshToken(ctx, "your_client_id", "your_client_secret", tokenResp.RefreshToken)
```

### 2. Product 服务 (client.Product)

提供商品相关的查询功能，包括品牌、分类、商品详情等。

#### 方法列表

| 方法                                 | 描述                 | 
|------------------------------------|--------------------|
| `BrandPage(ctx, token, req)`       | 4.1 查询品牌列表         | 
| `CategoryPage(ctx, token, req)`    | 4.2 查询分类列表         |
| `GoodsPage(ctx, token, req)`       | 4.3 查询商品列表(按SKU维度) |
| `GoodsDetail(ctx, token, req)`     | 4.4 查询商品详情         |
| `GoodsDetailSpec(ctx, token, req)` | 4.5 查询商品描述         |
| `Region(ctx, token, req)`          | 4.6 查询区域信息         |
| `GoodsPriceStock(ctx, token, req)` | 4.7 批量查询价格库存       |
| `CategoryAll(ctx, token, req)`     | 4.8 查询全量分类信息       |
| `JdPrice(ctx, token, req)`         | 4.9 查询京东商品价格       |
| `Similar(ctx, token, req)`         | 4.10 查询相似商品        |

#### 使用示例

```go
// 查询品牌列表
brandReq := &product.BrandPageRequest{
    Lang:     "中文",
    PageNum:  1,
    PageSize: 20,
}
brandResp, err := client.Product.BrandPage(ctx, token, brandReq)

// 查询商品详情
detailReq := &product.GoodsDetailRequest{
    Lang:    "中文",
    GoodsId: "goods123",
}
detailResp, err := client.Product.GoodsDetail(ctx, token, detailReq)
```

### 3. Order 服务 (client.Order)

提供订单相关的完整功能，包括下单、支付、查询、取消等。

#### 方法列表

| 方法                                | 描述                |
|-----------------------------------|-------------------|
| `Submit(ctx, token, req)`         | 5.1 提交订单          |
| `Pay(ctx, token, req)`            | 5.2 订单支付          |
| `Cancel(ctx, token, req)`         | 5.3 取消未支付订单       |
| `Confirm(ctx, token, req)`        | 5.4 确认收货          |
| `ConfirmBatch(ctx, token, req)`   | 5.5 批量确认收货        |
| `Track(ctx, token, req)`          | 5.6 查询订单配送信息      |
| `Detail(ctx, token, req)`         | 5.7 订单详情          |
| `Status(ctx, token, req)`         | 5.8 订单状态查询        |
| `CardNumber(ctx, token, req)`     | 5.9 查看虚拟订单卡券、卡密   |
| `TrackDelivery(ctx, token, req)`  | 5.10 根据物流单号查询配送信息 |
| `提供接口，配置后被动接收虚拟订单充值消息`            | 5.11 虚拟订单充值消息通知   |
| `Freight(ctx, token, req)`        | 5.12 运费查询         |
| `CheckAreaLimit(ctx, token, req)` | 5.13 查询商品区域限制     |
| `BoDetail(ctx, token, req)`       | 5.14 主订单查询详情      |

#### 使用示例

```go
// 提交订单
submitReq := &order.SubmitRequest{
    SkuList: []order.SkuItem{
        {
            Num:   2,
            SkuId: "sku123",
            Price: 99.99,
        },
    },
    ThirdOrder: "third_order_123",
    Name:       "张三",
    Mobile:     "13800138000",
    Address:    "北京市朝阳区xxx街道",
}
submitResp, err := client.Order.Submit(ctx, token, submitReq)

// 支付订单
payReq := &order.PayRequest{
    ThirdOrder: "third_order_123",
}
payResp, err := client.Order.Pay(ctx, token, payReq)
```

### 4. AfterSale 服务 (client.AfterSale)

提供售后服务的完整功能，包括创建售后单、查询售后信息等。

#### 方法列表

| 方法                                     | 描述                        |
|----------------------------------------|---------------------------|
| `QueryApply(ctx, token, req)`          | 3.1 查询售后类型以及包裹信息          |
| `Create(ctx, token, req)`              | 3.2 创建售后单                 |
| `ViewProgress(ctx, token, req)`        | 3.3 查询售后进度 ⚠️(废弃)         |
| `QueryRefund(ctx, token, req)`         | 3.4 查看退款单                 |
| `Cancellation(ctx, token, req)`        | 3.5 取消售后                  |
| `Sign(ctx, token, req)`                | 3.6 确认收货                  |
| `AddLogistics(ctx, token, req)`        | 3.7 添加物流信息                |
| `QueryLogistics(ctx, token, req)`      | 3.8 根据订单商品ID获取发货包裹信息      |
| `CreateRefund(ctx, token, req)`        | 3.9 创建退款售后(新版)            |
| `AfterProgress(ctx, token, req)`       | 3.10 查询售后进度(新版)           |
| `FindAfterDetail(ctx, token, req)`     | 3.11 查看售后详情(旧版)           |
| `QueryTradeNo(ctx, token, req)`        | 3.12 根据退款流水查询售后详情（定制商城专用） |
| `FindAfterDetailV2(ctx, token, req)`   | 3.13 查看售后详情(新版)           |
| `QueryLogisticsTrack(ctx, token, req)` | 3.14 查询物流轨迹(新版)           |

#### 使用示例

```go
// 查询可申请的售后类型
queryReq := &after.QueryApplyRequest{
    OrderSn: "order123",
    SkuId:   "sku123",
}
queryResp, err := client.AfterSale.QueryApply(ctx, token, queryReq)

// 创建售后单
createReq := []after.CreateRequest{
    {
        OrderSn:      "order123",
        SkuId:        "sku123",
        ReturnType:   0, // 0-退货/退款
        Reasons:      "商品质量问题",
        QuestionDesc: "商品有瑕疵",
    },
}
createResp, err := client.AfterSale.Create(ctx, token, &createReq)
```

### 5. Message 服务 (client.Message)

提供消息通知相关功能，用于接收平台推送的各种状态变更通知。

#### 方法列表

| 方法                          | 描述                   |
|-----------------------------|----------------------|
| `Query(ctx, token, req)`    | 8.1 消息通知查询接口 ⚠️ (旧版) |
| `QueryV4(ctx, token, req)`  | 8.2 消息通知查询接口 v4 版本   |
| `Delete(ctx, token, req)`   | 9.1 删除推送消息 ⚠️ (旧版)   |
| `DeleteV4(ctx, token, req)` | 9.2 删除推送消息 v4 版本     |

#### 消息类型说明

| 类型编号 | 描述             |
|------|----------------|
| 0    | 商品上下架变更        |
| 1    | 商品库存变更         |
| 2    | 商品基本信息变更       |
| 3    | 新增商品消息         |
| 4    | 商品价格变更         |
| 5    | 订单已发货变更        |
| 6    | 订单已取消变更        |
| 7    | 订单已完成变更        |
| 8    | 售后-客服已审核变更     |
| 9    | 售后-客诉人已退件变更    |
| 10   | 售后-供应商已收货变更    |
| 11   | 售后-供应商已发货变更    |
| 12   | 售后-客诉人已收件变更    |
| 13   | 售后-物流信息更新      |
| 14   | 售后-售后已退款变更     |
| 15   | 售后-退款已审核变更     |
| 16   | 售后-售后已取消变更     |
| 17   | 售后-售后已完成变更     |
| 18   | 售后-创建售后单通知变更   |
| 19   | 售后-售补偿售后单已审核变更 |
| 20   | 售后-补寄审核        |

#### 使用示例

```go
// 查询消息通知 (推荐使用 V4 版本)
queryReq := &message.QueryV4Request{
    QueryType: "1,2,3", // 查询库存变更、基本信息变更、新增商品消息
    DelFlag:   1,       // 1: 成功消费删除消息, 0: 不删除消息
}
queryResp, err := client.Message.QueryV4(ctx, token, queryReq)

// 删除已处理的消息 (推荐使用 V4 版本)
deleteReq := &message.DeleteV4Request{
    Ids: []string{"msg1", "msg2", "msg3"},
}
deleteResp, err := client.Message.DeleteV4(ctx, token, deleteReq)
```

## 版本说明

### ⚠️ 已废弃的方法

以下方法已被标记为废弃，建议使用新版本：

| 废弃方法                        | 推荐替代方法                        | 说明         |
|-----------------------------|-------------------------------|------------|
| `AfterSale.FindAfterDetail` | `AfterSale.FindAfterDetailV2` | 售后详情查询新版本  |
| `Message.Query`             | `Message.QueryV4`             | 消息查询 V4 版本 |
| `Message.Delete`            | `Message.DeleteV4`            | 消息删除 V4 版本 |

### 字段废弃说明

在商品详情中，以下字段已废弃：
- `BarCode` - 条形码字段已作废

## 错误处理

所有 API 调用都会返回标准的响应格式：

```go
type Response struct {
    Code    string      `json:"code"`    // 返回编码，"100" 表示成功
    Message string      `json:"message"` // 返回说明
    Data    interface{} `json:"data"`    // 具体数据
}
```

成功条件：HTTP 状态码为 200 且 `code` 为 "100"

## 注意事项

1. **认证**：除了获取 token 的接口外，所有 API 都需要在请求头中携带 `OPEN-API-AUTH` 字段
2. **语言支持**：商品相关接口支持多语言，可选项：
   - `中文` (默认)
   - `english`
   - `french`
   - `portuguese`
3. **限制**：
   - 提交订单一次最多支持 50 个商品
   - 订单状态查询最大支持 20 个订单
   - 批量确认收货最大支持 20 个订单
   - 消息查询每次最多返回 100 条数据，数据最多保留 60 天
   - 消息删除最大支持 100 个 ID
4. **超时**：未支付订单 24 小时后系统会自动取消
5. **物流查询**：只支持查询 2 个月内的订单配送信息
6. **虚拟商品**：
   - 虚拟商品下单时必须填写 `RechargeAccount` (充值账号)
   - 支付成功后可通过 `CardNumber` 方法获取卡券、卡密信息

## 许可证

本项目采用开源许可证，具体请查看 LICENSE 文件。

## 贡献

欢迎提交 Issue 和 Pull Request 来改进这个 SDK。如果你有特殊需求，也可以通过联系方式联系我。

## 完整示例

### 完整的订单流程示例

```go
package main

import (
    "context"
    "fmt"
    "log"
    "time"

    "github.com/li-bao-jia/maiger"
    "github.com/li-bao-jia/maiger/product"
    "github.com/li-bao-jia/maiger/order"
)

func main() {
    // 1. 初始化客户端
    client := maiger.NewClient("https://api.maiger.com")
    ctx := context.Background()

    // 2. 获取访问令牌
    tokenResp, err := client.Token.GetToken(ctx, "your_client_id", "your_client_secret")
    if err != nil {
        log.Fatal("获取token失败:", err)
    }
    token := tokenResp.AccessToken

    // 3. 查询商品信息
    goodsReq := &product.GoodsPageRequest{
        Lang:     "中文",
        PageNum:  1,
        PageSize: 10,
    }
    goodsResp, err := client.Product.GoodsPage(ctx, token, goodsReq)
    if err != nil {
        log.Fatal("查询商品失败:", err)
    }

    if len(goodsResp.Data.Products) == 0 {
        log.Fatal("没有找到商品")
    }

    // 4. 查询运费
    freightReq := &order.FreightRequest{
        Address: "北京市朝阳区",
        SkuIdList: []order.SkuIdList{
            {
                SkuId:    goodsResp.Data.Products[0].SkuId,
                Quantity: 1,
            },
        },
    }
    freightResp, err := client.Order.Freight(ctx, token, freightReq)
    if err != nil {
        log.Fatal("查询运费失败:", err)
    }

    // 5. 提交订单
    submitReq := &order.SubmitRequest{
        SkuList: []order.SkuItem{
            {
                Num:   1,
                SkuId: goodsResp.Data.Products[0].SkuId,
            },
        },
        ThirdOrder: "my_order_" + fmt.Sprintf("%d", time.Now().Unix()),
        Name:       "张三",
        Mobile:     "13800138000",
        Address:    "北京市朝阳区xxx街道xxx号",
        Province:   "北京市",
        City:       "北京市",
        County:     "朝阳区",
    }
    submitResp, err := client.Order.Submit(ctx, token, submitReq)
    if err != nil {
        log.Fatal("提交订单失败:", err)
    }

    if len(submitResp.Data) == 0 {
        log.Fatal("订单提交失败")
    }

    orderSn := submitResp.Data[0].OrderSn
    fmt.Printf("订单提交成功，订单号: %s\n", orderSn)

    // 6. 支付订单
    payReq := &order.PayRequest{
        ThirdOrder: submitReq.ThirdOrder,
    }
    payResp, err := client.Order.Pay(ctx, token, payReq)
    if err != nil {
        log.Fatal("支付订单失败:", err)
    }

    if payResp.Data {
        fmt.Println("订单支付成功")
    }

    // 7. 查询订单详情
    detailReq := &order.DetailRequest{
        OrderSn: orderSn,
    }
    detailResp, err := client.Order.Detail(ctx, token, detailReq)
    if err != nil {
        log.Fatal("查询订单详情失败:", err)
    }

    fmt.Printf("订单状态: %d, 支付状态: %d\n",
        detailResp.Data.OrderStatus, detailResp.Data.PayStatus)
}
```

### 消息通知处理示例

```go
package main

import (
    "context"
    "fmt"
    "log"
    "time"

    "github.com/li-bao-jia/maiger"
    "github.com/li-bao-jia/maiger/message"
)

func main() {
    client := maiger.NewClient("https://api.maiger.com")
    ctx := context.Background()

    // 获取token
    tokenResp, err := client.Token.GetToken(ctx, "your_client_id", "your_client_secret")
    if err != nil {
        log.Fatal("获取token失败:", err)
    }
    token := tokenResp.AccessToken

    // 定期轮询消息
    ticker := time.NewTicker(30 * time.Second)
    defer ticker.Stop()

    for {
        select {
        case <-ticker.C:
            // 查询所有类型的消息
            queryReq := &message.QueryV4Request{
                QueryType: "0,1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20",
                DelFlag:   1, // 查询成功后自动删除
            }

            queryResp, err := client.Message.QueryV4(ctx, token, queryReq)
            if err != nil {
                log.Printf("查询消息失败: %v", err)
                continue
            }

            if len(queryResp.Data.Contents) > 0 {
                fmt.Printf("收到 %d 条消息\n", len(queryResp.Data.Contents))

                for _, msg := range queryResp.Data.Contents {
                    handleMessage(msg)
                }
            }
        }
    }
}

func handleMessage(msg message.QueryContent) {
    switch msg.QueryType {
    case 0:
        fmt.Printf("商品上下架变更: SkuId=%s, Type=%d\n", msg.ThirdSkuId, msg.Type)
    case 1:
        fmt.Printf("商品库存变更: SkuId=%s, 新库存=%d\n", msg.ThirdSkuId, msg.SkuStock)
    case 4:
        fmt.Printf("商品价格变更: SkuId=%s, 销售价=%.2f, 指导价=%.2f\n",
            msg.ThirdSkuId, msg.SellPrice, msg.GoodsPrice)
    case 5:
        fmt.Printf("订单已发货: OrderSn=%s, SkuId=%s, 数量=%d\n",
            msg.OrderSn, msg.SkuId, msg.Quantity)
    case 6:
        fmt.Printf("订单已取消: OrderSn=%s\n", msg.OrderSn)
    case 7:
        fmt.Printf("订单已完成: OrderSn=%s\n", msg.OrderSn)
    default:
        fmt.Printf("收到消息类型 %d: %+v\n", msg.QueryType, msg)
    }
}
```

## 状态码说明

### 订单状态 (OrderStatus)

| 状态码 | 描述  |
|-----|-----|
| 0   | 待付款 |
| 1   | 待发货 |
| 3   | 已发货 |
| 5   | 已取消 |
| 11  | 已完成 |

### 支付状态 (PayStatus)

| 状态码 | 描述  |
|-----|-----|
| 0   | 未付款 |
| 2   | 已付款 |

### 售后类型 (ReturnType)

| 类型码 | 描述    |
|-----|-------|
| 0   | 退货/退款 |
| 2   | 换货    |
| 4   | 补寄    |
| 6   | 补偿费用  |
| 7   | 退款    |

### 售后状态 (AfterStatus)

| 状态码 | 描述   |
|-----|------|
| 1   | 待处理  |
| 2   | 处理中  |
| 3   | 已处理  |
| 4   | 已完成  |
| 6   | 取消售后 |

### 发货状态 (ShippingStatus)

| 状态码 | 描述   |
|-----|------|
| 0   | 未发货  |
| 1   | 部分发货 |
| 2   | 已发货  |
| 3   | 已签收  |

### 发货类型 (DeliverType)

| 类型码 | 描述     |
|-----|--------|
| 0   | 物流信息   |
| 1   | 商家送货上门 |
| 2   | 用户自取   |

## 高级功能

### 批量操作

```go
// 批量查询商品价格和库存
skuIds := []string{"sku1", "sku2", "sku3"}
priceStockReq := &product.GoodsPriceStockRequest{
    SkuIds: skuIds,
}
priceStockResp, err := client.Product.GoodsPriceStock(ctx, token, priceStockReq)

// 批量查询订单状态
orderSns := []string{"order1", "order2", "order3"}
statusResp, err := client.Order.Status(ctx, token, &orderSns)
```

### 错误处理最佳实践

```go
func handleAPIError(err error, resp interface{}) error {
    if err != nil {
        return fmt.Errorf("API调用失败: %w", err)
    }

    // 检查业务返回码
    switch v := resp.(type) {
    case *product.BrandPageResponse:
        if v.Code != "100" {
            return fmt.Errorf("业务错误: %s - %s", v.Code, v.Message)
        }
    case *order.SubmitResponse:
        if v.Code != "100" {
            return fmt.Errorf("业务错误: %s - %s", v.Code, v.Message)
        }
    // 添加其他响应类型的检查...
    }

    return nil
}

// 使用示例
brandResp, err := client.Product.BrandPage(ctx, token, brandReq)
if err := handleAPIError(err, brandResp); err != nil {
    log.Fatal(err)
}
```

## 常见问题 (FAQ)

### Q: 如何处理token过期？

A: 可以使用RefreshToken方法刷新token，或者重新获取新的token：

```go
// 方法1: 刷新token
refreshResp, err := client.Token.RefreshToken(ctx, clientId, clientSecret, oldRefreshToken)

// 方法2: 重新获取token
newTokenResp, err := client.Token.GetToken(ctx, clientId, clientSecret)
```

### Q: 订单提交后如何处理拆单？

A: 提交订单后系统会自动拆单，返回的数据中包含拆单信息：

```go
submitResp, err := client.Order.Submit(ctx, token, submitReq)
for _, orderData := range submitResp.Data {
    fmt.Printf("主单号: %s, 子订单号: %s\n", orderData.MainOrderSn, orderData.OrderSn)
}
```

### Q: 如何查询商品的详细规格信息？

A: 使用GoodsDetail方法可以获取完整的商品信息，包括SKU规格：

```go
detailResp, err := client.Product.GoodsDetail(ctx, token, &product.GoodsDetailRequest{
    GoodsId: "goods123",
})

for _, sku := range detailResp.Data.GoodsSkuBaseInfoVoList {
    fmt.Printf("SKU: %s, 规格: ", sku.SkuID)
    for _, spec := range sku.GoodsSpecList {
        fmt.Printf("%s:%s ", spec.SpecName, spec.SpecValue)
    }
    fmt.Println()
}
```

## ❤️ 支持与赞助

如果您觉得这个项目有帮助，请考虑赞助它的开发。非常感谢您的支持！

<table>
  <tr>
    <td><img src="/assets/wechat.jpg" width="200" alt="WeChat Pay" /></td>
    <td><img src="/assets/alipay.jpg" width="200" alt="Alipay" /></td>
  </tr>
</table>

### 我们的赞助商

非常感谢所有赞助商的慷慨支持！

-


### 广告

- 定制：如果您有定制需求、各种软件开发需求，或者您想封装某些SDK，请通过以下方式联系我

### 联系方式

- 开发者: BaoJia Li

- QQ: 751818588

- QQ群: 232185834

- 邮箱: livsyitian@163.com