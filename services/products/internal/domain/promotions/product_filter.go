package promotions

type ProductsFilter struct {
	category string
	price    uint64
}

func NewProductsFilter(category string, price uint64) ProductsFilter {
	return ProductsFilter{
		category: category,
		price:    price,
	}
}
