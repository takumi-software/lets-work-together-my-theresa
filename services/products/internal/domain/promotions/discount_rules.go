package promotions

var discountRules []DiscountRule

func init() {
	// only thing to take into consideration is to add them in descendant order or
	// TODO  redo ordering in the future?
	discountRules = append(discountRules, &bootsCategoryDiscount{})
	discountRules = append(discountRules, &bySKUDiscount{})
}

type DiscountRule interface {
	Discount(p *product) int
}

type bySKUDiscount struct{}

func (bsd *bySKUDiscount) Discount(p *product) int {
	if p.Sku == "000003" {
		return 15 // 15%
	}
	return 0 // do nothing
}

type bootsCategoryDiscount struct{}

func (bcd *bootsCategoryDiscount) Discount(p *product) int {
	if p.Category == "boots" {
		return 30 // 30%
	}
	return 0 // do nothing
}
