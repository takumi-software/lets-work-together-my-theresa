package adapter

import (
	"fmt"

	structpb "github.com/golang/protobuf/ptypes/struct"

	productspb "github.com/takumi-software/lets-work-together-my-theresa/protos/go/my-theresa/products"
	"github.com/takumi-software/lets-work-together-my-theresa/services/products/internal/domain/promotions"
)

func productsToProto(products promotions.Products) []*productspb.Product {
	var outProducts []*productspb.Product

	for _, product := range products {
		outProducts = append(outProducts, &productspb.Product{
			Sku:      product.Sku,
			Name:     product.Name,
			Category: product.Category,
			Price: &productspb.Price{
				Original:           float32(product.Price.Original),
				Final:              float32(product.Price.Final),
				DiscountPercentage: formatDiscount(product.Price.DiscountPercentage),
				Currency:           product.Price.Currency,
			},
		})
	}
	return outProducts
}

func formatDiscount(discount int) *structpb.Value {
	if discount == 0 {
		return &structpb.Value{Kind: &structpb.Value_NullValue{NullValue: structpb.NullValue_NULL_VALUE}}
	}
	return &structpb.Value{Kind: &structpb.Value_StringValue{StringValue: fmt.Sprint(discount, "%")}}
}
