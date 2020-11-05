package infrastructure

import (
	"encoding/json"
)

type Product struct {
	SKU      string `json:"sku"`
	Name     string `json:"name"`
	Category string `json:"category"`
	Price    uint64 `json:"price"`
}

type Raw struct {
	Products []Product `json:"products"`
}

func Find() (Raw, error) {
	var ps Raw
	memory := `{
	  "products": [
		{
		  "sku": "000001",
		  "name": "BV Lean leather ankle boots",
		  "category": "boots",
		  "price": 89000
		},
		{
		  "sku": "000002",
		  "name": "BV Lean leather ankle boots",
		  "category": "boots",
		  "price": 99000
		},
		{
		  "sku": "000003",
		  "name": "Ashlington leather ankle boots",
		  "category": "boots",
		  "price": 71000
		},
		{
		  "sku": "000004",
		  "name": "Naima embellished suede sandals",
		  "category": "sandals",
		  "price": 79500
		},
		{
		  "sku": "000005",
		  "name": "Nathane leather sneakers",
		  "category": "sneakers",
		  "price": 59000
		}
	  ]
	}`

	byteValue := []byte(memory)

	err := json.Unmarshal(byteValue, &ps)
	return ps, err
}
