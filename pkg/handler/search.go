package handler

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"sort"
	"sync"
	"time"

	"github.com/timezstyle/best_price/pkg/schema"
)

func find(ctx context.Context, shop schema.Shop, productName string, resultCh chan []schema.Product) {
	products, err := shop.Find(ctx, productName)
	if err != nil {
		log.Println(err)
	}
	resultCh <- products
}

func search(shops []schema.Shop, timeout time.Duration) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		q := req.URL.Query()
		productName := q.Get("productName")
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
		w.Header().Set("Content-Type", "application/json")
		b, _ := json.Marshal(results)
		w.Write(b)
	})
}
