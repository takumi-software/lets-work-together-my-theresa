package infrastructure

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

var jsonFile *os.File

func init() {
	jsonFile, _ = os.Open("products.json")
}

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
	byteValue, _ := ioutil.ReadAll(jsonFile)
	err := json.Unmarshal(byteValue, &ps)
	return ps, err
}
