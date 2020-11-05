package promotions

type bootsCategoryDiscount struct{}

func (bcd *bootsCategoryDiscount) Discount(p *product) int {
	if p.Category == "boots" {
		return 30 // 30%
	}
	return 0 // do nothing
}
