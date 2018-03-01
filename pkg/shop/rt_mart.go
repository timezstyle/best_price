package shop

import (
	"bytes"
	"context"
	"net/url"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/timezstyle/best_price/pkg/schema"
)

type RtMart struct {
}

func NewRtMart() *RtMart {
	return &RtMart{}
}

func (c *RtMart) Find(ctx context.Context, productName string) (ret []schema.Product, err error) {
	var (
		b   []byte
		doc *goquery.Document

		method = "GET"
		path   = "http://www.rt-mart.com.tw/direct/index.php?"
	)

	q := url.Values{}
	q.Set("prod_keyword", productName)
	q.Set("action", "product_search")
	q.Set("prod_size", "")
	q.Set("p_data_num", "30")
	q.Set("usort", "prod_selling_price,ASC")

	b, _, err = search(ctx, method, path+q.Encode(), "", nil)
	if err != nil {
		return
	}

	doc, err = goquery.NewDocumentFromReader(bytes.NewReader(b))
	if err != nil {
		return
	}

	ret = []schema.Product{}
	doc.Find(".classify_prolistBox .indexProList").Each(func(i int, s *goquery.Selection) {
		// For each item found
		var (
			name       string
			price      float64
			pictureURL string
			link       string
		)
		name = s.Find(".for_proname").Text()
		price, _ = strconv.ParseFloat(strings.Replace(s.Find(".for_pricebox").Text(), "$", "", -1), 64)
		pictureURL, _ = s.Find(".for_imgbox img").Attr("src")
		link, _ = s.Find(".for_proname a").Attr("href")

		p := schema.Product{
			Name:       name,
			PictureURL: pictureURL,
			Price:      price,
			Link:       link,
		}
		ret = append(ret, p)
	})
	return
}
