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
	Price        int `json:"price"`
	PriceFlashSale int `json:"price_flash_sale"`
}

var (
	Inventories []*ProductInventory
	totalProduct = 5
)

func SeedData() []*ProductInventory {
	priceRandom := []int{
		2000000,
		4000000,
		1250000,
		1300000,
		1200000,
	}

	for i := 1; i < totalProduct; i++ {
		n := util.Random( len(priceRandom) )
		productCode := util.RandomString( 5 )
		productName := fmt.Sprintf("Product Test-%v", i)
		Inventories = append(Inventories, &ProductInventory{
			ProductId: i,
			ProductName: productName,
			ProductCode: productCode,
			ProductStock: 5,
			Price: priceRandom[n],
			PriceFlashSale: 12000, //Flash Sale 12.12
		})
	}

	return Inventories
}