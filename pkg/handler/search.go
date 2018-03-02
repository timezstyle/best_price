package handler

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"reflect"
	"sort"
	"strconv"
	"sync"
	"time"

	"github.com/timezstyle/best_price/pkg/schema"
)

func GetPositiveIntOrDefault(input string, def int) (ret int) {
	var err error
	ret, err = strconv.Atoi(input)
	if err != nil {
		ret = def
	}
	if ret < 0 {
		ret = def
	}
	return
}

func find(ctx context.Context, shop schema.Shop, productName string, resultCh chan []schema.Product) {
	products, err := shop.Find(ctx, productName)
	if err != nil {
		log.Println(reflect.TypeOf(shop), err)
	}
	resultCh <- products
}

func search(shops []schema.Shop, timeout time.Duration) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		q := req.URL.Query()
		productName := q.Get("product_name")
		offset := GetPositiveIntOrDefault(q.Get("offset"), 0)
		limit := GetPositiveIntOrDefault(q.Get("limit"), 30)
		results := []schema.Product{}

		ctx, cancel := context.WithTimeout(req.Context(), timeout)
		defer cancel()

		resultCh := make(chan []schema.Product)
		defer close(resultCh)

		var wg sync.WaitGroup
		go func() {
			for result := range resultCh {
				results = append(results, result...)
				wg.Done()
			}
		}()
		for i := range shops {
			wg.Add(1)
			shop := shops[i]
			go find(ctx, shop, productName, resultCh)
		}
		wg.Wait()

		sort.Sort(schema.SortByPrice(results))

		resultsLen := len(results)
		lastestLen := offset + limit
		if resultsLen >= lastestLen {
			results = results[offset:lastestLen]
		} else {
			if resultsLen >= offset {
				results = results[offset:]
			} else {
				results = []schema.Product{}
			}
		}

		w.Header().Set("Content-Type", "application/json")
		b, _ := json.Marshal(map[string]interface{}{
			"total":   resultsLen,
			"results": results,
		})
		w.Write(b)
	})
}
