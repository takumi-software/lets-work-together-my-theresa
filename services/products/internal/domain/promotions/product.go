package promotions

import "github.com/takumi-software/lets-work-together-my-theresa/services/products/internal/infrastructure"

type product struct {
	Sku      string
	Name     string
	Category string
	Price    price
}

func NewProduct(sku, name, category string, pr uint64) product {
	product := product{
		Sku:      sku,
		Name:     name,
		Category: category,
	}

	discountPercentage := product.Discount()
	finalPrice := pr
	if discountPercentage > 0 {
		finalPrice = pr * uint64(100/discountPercentage)
	}

	product.Price = price{
		original:           pr,
		final:              finalPrice,
		discountPercentage: discountPercentage,
		currency:           setCurrency(),
	}

	return product
}

type Products []product

func Fetch(filter ProductsFilter) (Products, error) {
	var products Products
	raw, err := infrastructure.Find()
	if err != nil {
		return nil, err
	}
	for _, p := range raw.Products {
		products = append(products, NewProduct(p.SKU, p.Name, p.Category, p.Price))
	}
	return products, nil
}

func (p *product) Discount() int {
	for _, rule := range discountRules {
		discount := rule.Discount(p)
		if discount > 0 {
			return discount
		}
	}
	return 0
}
