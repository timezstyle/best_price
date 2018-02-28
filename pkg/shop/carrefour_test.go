package shop

import (
	"context"
	"testing"
)

func TestCarrefour_Find(t *testing.T) {
	c := NewCarrefour()
	type args struct {
		ctx         context.Context
		productName string
	}
	tests := []struct {
		name    string
		c       *Carrefour
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{"find pen", c, args{context.Background(), "pen"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Carrefour{}
			gotRet, err := c.Find(tt.args.ctx, tt.args.productName)
			if (err != nil) != tt.wantErr {
				t.Errorf("Carrefour.Find() error = %v, wantErr %v", err, tt.wantErr)
				return
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
