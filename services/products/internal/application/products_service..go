package application

import (
	"context"

	"github.com/takumi-software/lets-work-together-my-theresa/services/products/internal/domain/promotions"
)

// here we need the interface
type ProductsService interface {
	Fetch(ctx context.Context, input *FetchFilterInput) (*FetchOutput, error)
}

type FetchFilterInput struct {
	Category string
	Price    uint64
}

type FetchOutput struct {
	Products promotions.Products
}
