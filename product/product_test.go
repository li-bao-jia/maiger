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
		GoodsId: "57c6e274-b882-4241-b3ab-3d5a754d997c",
	})

	if err != nil {
		t.Fatalf("GoodsDetail error: %v", err)
	}
	fmt.Println("GoodsDetail:", detailResp.Data)
}

// 测试 4.5 获取商品描述

func TestGoodsDetailSpec(t *testing.T) {
	detailSpecResp, err := goodsService.GoodsDetailSpec(ctx, accessToken, &GoodsDetailSpecRequest{
		GoodsId: "0d0189c1-e062-4ffa-9db0-b8c6e42549f3",
	})

	if err != nil {
		t.Fatalf("GoodsDetailSpec error: %v", err)
	}
	fmt.Println("GoodsDetailSpec:", detailSpecResp.Data)
}

// 测试 4.6 查询区域信息

func TestRegion(t *testing.T) {
	regionResp, err := goodsService.Region(ctx, accessToken, &RegionRequest{
		ParentId: 1,
	})

	if err != nil {
		t.Fatalf("Region error: %v", err)
	}
	fmt.Println("Region:", regionResp.Data)
}

// 测试 4.7 批量查询价格库存

func TestGoodsPriceStock(t *testing.T) {
	goodsPriceStockResp, err := goodsService.GoodsPriceStock(ctx, accessToken, &GoodsPriceStockRequest{
		SkuIds: []string{"2a2e54bd-389c-4bf2-945b-312fa9099f35"},
	})

	if err != nil {
		t.Fatalf("GoodsPriceStock error: %v", err)
	}
	fmt.Println("GoodsPriceStock:", goodsPriceStockResp.Data)
}

// 测试 4.8 查询全量分类信息

func TestCategoryAll(t *testing.T) {
	categoryAllResp, err := goodsService.CategoryAll(ctx, accessToken, &CategoryAllRequest{
		Lang: "中文",
	})

	if err != nil {
		t.Fatalf("CategoryAll error: %v", err)
	}
	fmt.Println("CategoryAll:", categoryAllResp.Data)
}

// 测试 4.9 查询京东商品价格

func TestJdPrice(t *testing.T) {
	jdPriceResp, err := goodsService.JdPrice(ctx, accessToken, &[]JdPriceRequest{
		{
			GoodsUrl:     "https://item.jd.com/100023408309.html",
			ThirdGoodsNo: "100023408309",
		},
	})

	if err != nil {
		t.Fatalf("JdPrice error: %v", err)
	}
	fmt.Println("JdPrice:", jdPriceResp.Data)
}

// 测试 4.10 查询相似商品

func TestSimilar(t *testing.T) {
	similarResp, err := goodsService.Similar(ctx, accessToken, &SimilarRequest{
		SkuId: "4575b922-0e8a-484d-908e-63e49c63755b",
	})

	if err != nil {
		t.Fatalf("Similar error: %v", err)
	}
	fmt.Println("Similar:", similarResp.Data)
}
