package main

import (
	"flag"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/timezstyle/best_price/pkg/handler"
	"github.com/timezstyle/best_price/pkg/schema"
	"github.com/timezstyle/best_price/pkg/shop"
)

func main() {
	var (
		port            string
		reqTimeout      time.Duration
		seleniumHubAddr string
	)
	flag.StringVar(&port, "port", ":3000", "listen port")
	flag.DurationVar(&reqTimeout, "request_timeout", 3*time.Second, "request's timeout when send to shop")
	flag.StringVar(&seleniumHubAddr, "selenium_hub_addr", "http://localhost:4444/wd/hub", "selenium hub addr")
	flag.Parse()

	root := handler.Root{
		Router: mux.NewRouter(),
		Shops: []schema.Shop{
			// add shop here
			shop.NewRtMart(),
			shop.NewCarrefour(),
			shop.NewShopee(seleniumHubAddr),
			shop.NewTaobao(),
		},
		ReqTimeout: reqTimeout,
	}
	root.Setup()

	log.Println("listening at ", port)
	http.ListenAndServe(port, root.Router)

}
