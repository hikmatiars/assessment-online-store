package entity

import (
	"assessment-online-store/util"
	"fmt"
)

type ProductInventory struct {
	ProductId int `json:"product_id"`
	ProductName string `json:"product_name"`
	ProductCode string `json:"product_code"`
	ProductStock int `json:"product_stock"`
}

var (
	Inventories []*ProductInventory
	totalProduct = 5
)

func SeedData() []*ProductInventory {
	for i := 1; i < totalProduct; i++ {
		productCode := util.RandomString( 5 )
		productName := fmt.Sprintf("Product Test-%v", i)
		Inventories = append(Inventories, &ProductInventory{
			ProductId: i,
			ProductName: productName,
			ProductCode: productCode,
			ProductStock: 5,
		})
	}

	return Inventories
}