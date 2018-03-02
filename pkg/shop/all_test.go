package shop

import (
	"context"
	"testing"

	"github.com/timezstyle/best_price/pkg/schema"
)

func TestShop_Find(t *testing.T) {
	type args struct {
		ctx         context.Context
		productName string
	}
	tests := []struct {
		name    string
		c       schema.Shop
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{"Carrefour find pen", NewCarrefour(), args{context.Background(), "pen"}, false},
		{"RtMart find pen", NewRtMart(), args{context.Background(), "pen"}, false},
		{"Shopee find pen", NewShopee("http://localhost:4444/wd/hub"), args{context.Background(), "pen"}, false},
		{"Taobao find pen", NewTaobao(), args{context.Background(), "pen"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRet, err := tt.c.Find(tt.args.ctx, tt.args.productName)
			if (err != nil) != tt.wantErr {
				t.Errorf("Shop.Find() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if len(gotRet) == 0 {
				t.Errorf("Products is empty")
			}
			for i := range gotRet {
				product := gotRet[i]
				if product.Name == "" {
					t.Errorf("Name is blank")
				}
				if product.Link == "" {
					t.Errorf("Link is blank")
				}
				if product.Price == 0 {
					t.Errorf("Price is 0")
				}
				if product.PictureURL == "" {
					t.Errorf("PictureURL is blank")
				}
			}
		})
	}
}
