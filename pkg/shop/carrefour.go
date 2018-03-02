package shop

import (
	"context"
	"encoding/json"
	"net/http"
	"net/url"
	"strconv"

	"github.com/timezstyle/best_price/pkg/schema"
	"github.com/timezstyle/best_price/pkg/util"
)

type Carrefour struct {
}

func NewCarrefour() *Carrefour {
	return &Carrefour{}
}

func (c *Carrefour) Find(ctx context.Context, productName string) (ret []schema.Product, err error) {
	var (
		b    []byte
		resp carrefourResponse

		method = "POST"
		path   = "https://online.carrefour.com.tw/CarrefourECProduct/GetSearchJson"
	)

	q := url.Values{}
	q.Set("key", productName)
	q.Set("categoryId", "")
	q.Set("orderBy", "10")
	q.Set("pageIndex", "1")
	q.Set("pageSize", "35")
	q.Set("minPrice", "")
	q.Set("maxPrice", "")

	h := http.Header{}
	h.Set("Content-Type", "application/x-www-form-urlencoded")
	b, _, err = util.Search(ctx, method, path, q.Encode(), &h)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return
	}

	ret = []schema.Product{}
	for i := range resp.Content.ProductListModel {
		product := resp.Content.ProductListModel[i]

		var finalPrice float64
		if product.SpecialPrice != "" {
			finalPrice, err = strconv.ParseFloat(product.SpecialPrice, 64)
		} else {
			finalPrice, err = strconv.ParseFloat(product.Price, 64)
		}
		if err != nil {
			return
		}

		p := schema.Product{
			Name:       product.Name,
			Price:      finalPrice,
			PictureURL: product.PictureURL,
			Link:       "https://carrefoureccdn.azureedge.net" + product.SeName,
		}
		ret = append(ret, p)
	}
	return
}

type carrefourResponse struct {
	Success int `json:"success"`
	Content struct {
		CategoryID              int         `json:"CategoryId"`
		CurrentCategoryID       int         `json:"CurrentCategoryId"`
		CurrentCategoryName     interface{} `json:"CurrentCategoryName"`
		ParentCategoryID1       int         `json:"ParentCategoryId1"`
		ParentCategoryName1     interface{} `json:"ParentCategoryName1"`
		ParentCategoryID2       int         `json:"ParentCategoryId2"`
		ParentCategoryName2     interface{} `json:"ParentCategoryName2"`
		ParentCategorySeName2   interface{} `json:"ParentCategorySeName2"`
		ParentCategoryID3       int         `json:"ParentCategoryId3"`
		ParentCategoryName3     interface{} `json:"ParentCategoryName3"`
		ParentCategorySeName3   interface{} `json:"ParentCategorySeName3"`
		ParentCategoryName4     interface{} `json:"ParentCategoryName4"`
		ParentCategorySeName4   interface{} `json:"ParentCategorySeName4"`
		IsPromotionAreaPorducts bool        `json:"IsPromotionAreaPorducts"`
		IsNewHomePageProducts   bool        `json:"IsNewHomePageProducts"`
		IsRewardProducts        bool        `json:"IsRewardProducts"`
		StoreActivityBasicID    int         `json:"StoreActivityBasicId"`
		ActivityPictureURL      interface{} `json:"ActivityPictureUrl"`
		Note                    interface{} `json:"Note"`
		NewHomePageID           int         `json:"NewHomePageId"`
		NewHomePageBigTitleID   int         `json:"NewHomePageBigTitleId"`
		NewHomePageBigTitleName interface{} `json:"NewHomePageBigTitleName"`
		RewardID                int         `json:"RewardId"`
		Key                     string      `json:"Key"`
		SearchCategoryID        string      `json:"searchCategoryId"`
		MetaTitle               interface{} `json:"MetaTitle"`
		MetaKeywords            interface{} `json:"MetaKeywords"`
		MetaDescription         interface{} `json:"MetaDescription"`
		ProductListModel        []struct {
			ID                   int         `json:"Id"`
			PictureURL           string      `json:"PictureUrl"`
			Name                 string      `json:"Name"`
			XingHao              interface{} `json:"XingHao"`
			SeName               string      `json:"SeName"`
			SubTitle             interface{} `json:"SubTitle"`
			ActiveBeginTime      interface{} `json:"ActiveBeginTime"`
			ActiveEndTime        interface{} `json:"ActiveEndTime"`
			ActiveName           interface{} `json:"ActiveName"`
			ActiveBackColor      interface{} `json:"ActiveBackColor"`
			ActiveFontSize       int         `json:"ActiveFontSize"`
			ActiveURL            interface{} `json:"ActiveUrl"`
			UnitCode             interface{} `json:"UnitCode"`
			ItemQtyPerPack       int         `json:"ItemQtyPerPack"`
			ItemQtyPerPackFormat string      `json:"ItemQtyPerPackFormat"`
			ProductVolume        interface{} `json:"ProductVolume"`
			SpecialPrice         string      `json:"SpecialPrice"`
			Price                string      `json:"Price"`
			ActivityPrice        string      `json:"ActivityPrice"`
			BonusCount           interface{} `json:"BonusCount"`
			ShortDescription     interface{} `json:"ShortDescription"`
			ProductOperation     struct {
				ItemQtyPerPack         int         `json:"ItemQtyPerPack"`
				MaxNumberOnSale        int         `json:"MaxNumberOnSale"`
				IsQuickShippingModel   bool        `json:"IsQuickShippingModel"`
				StockQuantity          interface{} `json:"StockQuantity"`
				IsPreOrderShippingMode bool        `json:"IsPreOrderShippingMode"`
				IsOffIslandMode        bool        `json:"IsOffIslandMode"`
				IsCommonOrder          bool        `json:"IsCommonOrder"`
				ID                     int         `json:"Id"`
			} `json:"ProductOperation"`
			PromotionAreaType              interface{} `json:"PromotionAreaType"`
			DisplayID                      int         `json:"DisplayId"`
			HomePictureURL                 interface{} `json:"HomePictureUrl"`
			QucikShippingProductListPicURL string      `json:"QucikShippingProductListPicUrl"`
			SpecialStoreProductListPicURL  string      `json:"SpecialStoreProductListPicUrl"`
			OffIslandProductListPicURL     interface{} `json:"OffIslandProductListPicUrl"`
			PromotionProductPicURL         string      `json:"PromotionProductPicUrl"`
			ProductNumShow                 interface{} `json:"ProductNumShow"`
			IsWish                         bool        `json:"IsWish"`
			Specification                  string      `json:"Specification"`
		} `json:"ProductListModel"`
		ProductIds       string `json:"ProductIds"`
		ProductIdsTop20  string `json:"ProductIdsTop20"`
		CurrentPageIndex int    `json:"CurrentPageIndex"`
		PageSize         int    `json:"PageSize"`
		OrderByID        int    `json:"OrderById"`
		Count            int    `json:"Count"`
		ID               int    `json:"Id"`
	} `json:"content"`
}
