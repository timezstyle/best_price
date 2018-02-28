package schema

import "context"

type Shop interface {
	Find(ctx context.Context, productName string) ([]Product, error)
}
