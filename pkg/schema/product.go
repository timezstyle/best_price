package schema

type Product struct {
	Name       string  `json:"name"`
	Price      float64 `json:"price"`
	PictureURL string  `json:"picture_url"`
	Link       string  `json:"link"`
}

type SortByPrice []Product

func (a SortByPrice) Len() int      { return len(a) }
func (a SortByPrice) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a SortByPrice) Less(i, j int) bool {
	return a[i].Price < a[j].Price
}
