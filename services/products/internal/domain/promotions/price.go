package promotions

type Price struct {
	Original           uint64
	Final              uint64
	DiscountPercentage int
	Currency           string
}
