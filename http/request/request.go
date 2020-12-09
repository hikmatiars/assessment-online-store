package request

type AddToCart struct {
	ProductId int `json:"product_id"`
	Quantity  int `json:"quantity"`
}