package adapter

import (
	"context"

	"github.com/takumi-software/lets-work-together-my-theresa/protos/go/my-theresa/products"
	"github.com/takumi-software/lets-work-together-my-theresa/services/products/internal/application"
)

type GRPCServer struct {
	service application.ProductsService
}

func NewGRPCServer(service application.ProductsService) (GRPCServer, error) {
	return GRPCServer{
		service: service,
	}, nil
}

func (g *GRPCServer) Fetch(ctx context.Context, request *products.FetchProductsRequest) (*products.FetchProductsResponse, error) {
	out, err := g.service.Fetch(ctx, &application.FetchFilterInput{
		Category: request.Category,
		Price:    request.Price,
	})
	if err != nil {
		return nil, err
	}
	ps := productsToProto(out.Products)

	return &products.FetchProductsResponse{
		Products: ps,
	}, nil
}
