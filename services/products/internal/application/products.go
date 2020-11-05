package application

import (
	products "github.com/takumi-software/lets-work-together-my-theresa/services/products/internal/domain/promotions"

	"golang.org/x/net/context"
)

type productsService struct {
}

// NewService creates an instance of products use case
func NewService() ProductsService {
	return &productsService{}
}

func (ps *productsService) Fetch(ctx context.Context, filter *FetchFilterInput) (*FetchOutput, error) {
	prs, err := products.Fetch(products.NewProductsFilter(filter.Category, filter.Price))
	if err != nil {
		return nil, err
	}
	return &FetchOutput{
		Products: prs,
	}, nil
}
