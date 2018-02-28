package handler

import (
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/timezstyle/best_price/pkg/schema"
)

type Root struct {
	Shops      []schema.Shop
	Router     *mux.Router
	ReqTimeout time.Duration
}

func (r *Root) Setup() {
	r.Router.HandleFunc("/search", search(r.Shops, r.ReqTimeout)).
		Methods(http.MethodGet)
}
