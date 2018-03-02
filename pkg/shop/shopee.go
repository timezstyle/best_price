package shop

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"

	"github.com/tebeka/selenium"
	"github.com/timezstyle/best_price/pkg/schema"
	"github.com/timezstyle/best_price/pkg/util"
)

type Shopee struct {
	cookie string
	token  string
}

func NewShopee(seleniumHubAddr string) *Shopee {
	c := &Shopee{}
	var (
		webDriver selenium.WebDriver
		err       error
	)
	caps := selenium.Capabilities{"browserName": "chrome"}
	if webDriver, err = selenium.NewRemote(caps, seleniumHubAddr); err != nil {
		log.Fatalf("Failed to open session: %s\n", err)
	}
	defer webDriver.Quit()

	webDriver.Get("https://shopee.tw")
	cookies, err := webDriver.GetCookies()
	if err != nil {
		log.Fatalf("get cookie failed: %s\n", err)
	}
	for i := range cookies {
		cookie := cookies[i]
		if cookie.Name == "csrftoken" {
			c.token = cookie.Value
		}
		c.cookie += cookie.Name + "=" + cookie.Value + ";"
	}
	// agent, err := webDriver.ExecuteScript("return navigator.userAgent", []interface{}{})
	// if err != nil {
	// 	log.Fatalf("get user agent failed: %s\n", err)
	// }
	// log.Println(agent)
	return c
}

func (c *Shopee) Find(ctx context.Context, productName string) (ret []schema.Product, err error) {
	var (
		b    []byte
		resp []shopeeItemResponse

		method = "GET"
		path   = "https://shopee.tw/api/v1/search_items/?"
	)

	q := url.Values{}
	q.Set("keyword", productName)
	q.Set("by", "price")
	q.Set("order", "asc")
	q.Set("newest", "0")
	q.Set("limit", "50")

	referer := path + q.Encode()
	b, _, err = util.Search(ctx, method, referer, "", nil)
	if err != nil {
		return
	}

	idResp := shopeeIDResponse{}
	err = json.Unmarshal(b, &idResp)
	if err != nil {
		return
	}

	var itemsJSON []byte
	itemsJSON, err = json.Marshal(map[string]interface{}{
		"item_shop_ids": idResp.Items,
	})
	if err != nil {
		return
	}

	q = url.Values{}
	q.Set("keyword", productName)

	h := http.Header{}
	h.Set("Content-Type", "application/json")
	h.Set("Cookie", c.cookie)
	h.Set("x-csrftoken", c.token)
	h.Set("referer", "https://shopee.tw/search/?"+q.Encode())
	b, _, err = util.Search(ctx, "POST", "https://shopee.tw/api/v1/items/", string(itemsJSON), &h)
	if err != nil {
		return
	}

	err = json.Unmarshal(b, &resp)
	if err != nil {
		return
	}

	ret = []schema.Product{}
	for i := range resp {
		product := resp[i]

		var finalPrice float64
		finalPrice = float64(product.Price / 100000)
		if err != nil {
			return
		}

		p := schema.Product{
			Name:       product.Name,
			Price:      finalPrice,
			PictureURL: fmt.Sprintf("https://cfshopeetw-a.akamaihd.net/file/%s_tn", product.Image),
			Link:       fmt.Sprintf("https://shopee.tw/%s-i.%d.%d", url.QueryEscape(product.Name), product.Shopid, product.Itemid),
		}
		ret = append(ret, p)
	}
	return
}

type shopeeIDResponse struct {
	Items []struct {
		Itemid     int   `json:"itemid"`
		Shopid     int   `json:"shopid"`
		Logisticid []int `json:"logisticid"`
	} `json:"items"`
}

type shopeeItemResponse struct {
	Itemid int    `json:"itemid"`
	Image  string `json:"image"`
	Shopid int    `json:"shopid"`
	Name   string `json:"name"`
	Price  int    `json:"price"`
}
