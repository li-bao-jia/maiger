package product

import (
	"context"
	"fmt"
	"testing"

	"github.com/li-bao-jia/maiger/http"
)

var (
	domain      = "https://www.xxxxxxxx.com"
	accessToken = "your_token"

	ctx          = context.Background()
	httpClient   = http.NewHTTPClient(domain)
	goodsService = NewGoodsService(httpClient)
)

// 测试 4.1 查询品牌列表

func TestBrandPage(t *testing.T) {
	brandResp, err := goodsService.BrandPage(ctx, accessToken, &BrandPageRequest{
		Lang:      "中文",
		PageNum:   2,
		PageSize:  60,
		BrandId:   "",
		BrandName: "",
	})

	if err != nil {
		t.Fatalf("BrandPage error: %v", err)
	}
	fmt.Println("BrandPage:", brandResp.Data.Brands)
}

// 测试 4.2 查询分类列表

func TestCategoryPage(t *testing.T) {
	categoryResp, err := goodsService.CategoryPage(ctx, accessToken, &CategoryPageRequest{
		Lang:     "中文",
		PageNum:  2,
		PageSize: 20,
		ParentId: "",
	})
	if err != nil {
		t.Fatalf("CategoryPage error: %v", err)
	}
	fmt.Println("CategoryPage:", categoryResp.Data.Categories)
}

// 测试 4.3 查询商品列表

func TestGoodsPage(t *testing.T) {
	goodsResp, err := goodsService.GoodsPage(ctx, accessToken, &GoodsPageRequest{
		Lang:      "中文",
		PageNum:   2,
		PageSize:  100,
		CatId:     "",
		BrandId:   "",
		GoodsName: "",
	})

	if err != nil {
		t.Fatalf("GoodsPage error: %v", err)
	}
	fmt.Println("GoodsPage:", goodsResp.Data.Products)
}

// 测试 4.4 查询商品详情

func TestGoodsDetail(t *testing.T) {
	detailResp, err := goodsService.GoodsDetail(ctx, accessToken, &GoodsDetailRequest{
		Lang:    "中文",
		SkuId:   "",
		GoodsId: "0d0189c1-e062-4ffa-9db0-b8c6e42549f3",
	})

	if err != nil {
		t.Fatalf("GoodsPage error: %v", err)
	}
	fmt.Println("GoodsPage:", detailResp.Data)
}
