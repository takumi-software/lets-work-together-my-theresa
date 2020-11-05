package promotions

var discountRules []Discounter

func init() {
	// only thing to take into consideration is to add them in descendant order or
	// TODO  redo ordering in the future?
	discountRules = append(discountRules, &bootsCategoryDiscount{})
	discountRules = append(discountRules, &bySKUDiscount{})
}

type Discounter interface {
	Discount(p *product) int
}
