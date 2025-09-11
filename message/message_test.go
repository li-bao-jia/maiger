package message

import (
	"context"
	"fmt"
	"testing"

	"github.com/li-bao-jia/maiger/http"
)

var (
	domain      = "https://www.xxxxxxxx.com"
	accessToken = "your_token"

	ctx            = context.Background()
	httpClient     = http.NewHTTPClient(domain)
	messageService = NewMessageService(httpClient)
)

// 测试 8.1 消息通知查询接口

func TestQuery(t *testing.T) {
	queryResp, err := messageService.Query(ctx, accessToken, &QueryRequest{
		QueryType: "0,1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20",
	})

	if err != nil {
		t.Fatalf("Query error: %v", err)
	}
	fmt.Println("Query:", queryResp.Data)
}
