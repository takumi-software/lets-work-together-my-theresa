package promotions

type bySKUDiscount struct{}

func (bsd *bySKUDiscount) Discount(p *product) int {
	if p.Sku == "000003" {
		return 15 // 15%
	}
	return 0 // do nothing
}
