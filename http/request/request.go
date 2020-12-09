package request

type AddToCart struct {
	ProductId int `json:"product_id"`
	Quantity  int `json:"quantity"`
}

type Checkout struct {
	Products []ObjProductCheckout `json:"products"`
}

type ObjProductCheckout struct {
	ProductId int `json:"product_id"`
	Quantity  int `json:"quantity"`
}