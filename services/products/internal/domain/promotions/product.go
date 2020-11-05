package promotions

import (
	"github.com/takumi-software/lets-work-together-my-theresa/services/products/internal/infrastructure"
)

type product struct {
	Sku      string
	Name     string
	Category string
	Price    Price
}

func NewPromotionProduct(sku, name, category string, pr uint64) product {
	product := product{
		Sku:      sku,
		Name:     name,
		Category: category,
	}

	discountPercentage := product.Discount()
	finalPrice := pr
	if discountPercentage > 0 && finalPrice > 0 {
		finalPrice = finalPrice - ((uint64(discountPercentage) * finalPrice) / 100)
	}

	product.Price = Price{
		Original:           pr,
		Final:              finalPrice,
		DiscountPercentage: discountPercentage,
		Currency:           setCurrency(),
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
		if meetsFilters(p, filter) {
			products = append(products, NewPromotionProduct(p.SKU, p.Name, p.Category, p.Price))
		}
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

func meetsFilters(product infrastructure.Product, filter ProductsFilter) bool {
	results := true
	if filter.category != "" && filter.category != product.Category {
		results = false
	}
	if filter.price != 0 && filter.price != product.Price {
		results = false
	}
	return results
}
