package entity

type Cart struct {
	ProductId int `json:"product_id"`
	ProductName string `json:"product_name"`
	Quantity    int `json:"quantity"`
	Price       int `json:"price"`
}

var Carts []*Cart