package promotions

type price struct {
	original           uint64
	final              uint64
	discountPercentage int
	currency           string
}
