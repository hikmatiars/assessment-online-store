package usecase

import (
	"assessment-online-store/entity"
	"assessment-online-store/http/request"
	"assessment-online-store/util"
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"
)

//Init Use case interface
type InterfaceUseCase interface {
	GetListProductUseCase() []*entity.ProductInventory
	AddCartUseCase(req request.AddToCart ) (int, error)
	GetListCartUseCase() (int, []*entity.Cart)
	CheckoutItemUseCase(req request.Checkout) (int, []map[string]interface{}, error)
}

type UseCase struct {
	Context context.Context
	Inventories []*entity.ProductInventory
	Cart []*entity.Cart
	TimeFlashSale time.Time
}

func NewUseCase(ctx context.Context, inventories []*entity.ProductInventory, cart []*entity.Cart, flashSale time.Time) InterfaceUseCase {
	return &UseCase{
		Context: ctx,
		Inventories : inventories,
		Cart: cart,
		TimeFlashSale: flashSale,
	}
}

func (uc *UseCase) GetListProductUseCase() []*entity.ProductInventory {
	return uc.Inventories
}

func (uc *UseCase) AddCartUseCase(req request.AddToCart ) (int, error) {
	var index int

	//get index with sampe product id
	for i, inventory := range uc.Inventories {
		if inventory.ProductId == req.ProductId {
			index = i
			break
		}
	}

	if req.Quantity > uc.Inventories[index].ProductStock {
		return http.StatusUnprocessableEntity, errors.New("product out of stock")
	}

	uc.Inventories[index].ProductStock -= req.Quantity

	if len(uc.Cart) > 0 {
		//checking product id is exist or not on list cart
		for i, val := range uc.Cart {
			if val.ProductId == req.ProductId {
				uc.Cart[i].Quantity += req.Quantity
				break
			}
		}

		return http.StatusOK, nil
	}

	uc.Cart = append(uc.Cart, &entity.Cart{
		ProductId: uc.Inventories[index].ProductId,
		ProductName: uc.Inventories[index].ProductName,
		Quantity: req.Quantity,
		Price: uc.Inventories[index].Price,
		PriceFlashSale: uc.Inventories[index].PriceFlashSale,
	})

	return http.StatusOK, nil
}

func (uc *UseCase) GetListCartUseCase() (int, []*entity.Cart) {
	if len(uc.Cart) == 0 {
		return http.StatusNoContent, uc.Cart
	}

	return http.StatusOK, uc.Cart
}

func (uc *UseCase) CheckoutItemUseCase(req request.Checkout) (int, []map[string]interface{}, error) {
	var dataCheckout []map[string]interface{}

	if len(uc.Cart) == 0 {
		return http.StatusUnprocessableEntity, dataCheckout, errors.New("list cart empty")
	}

	for _, itemCheckout := range req.Products {

		index := GetIndexProductInventory( itemCheckout.ProductId, uc.Inventories )

		for _, cart := range uc.Cart {
			var price int

			if itemCheckout.ProductId == cart.ProductId {

				if itemCheckout.Quantity > cart.Quantity {
					//validate if request quantity more than item on cart
					errs := fmt.Sprintf("qty product checkout item %s is out stock", cart.ProductName)
					return http.StatusUnprocessableEntity, dataCheckout, errors.New(errs)
				}

				if util.DatePassed( uc.TimeFlashSale, time.Now() ) {
					price = cart.PriceFlashSale
				}else {
					price = cart.Price
				}

				dataCheckout = append(dataCheckout, map[string]interface{}{
					"product_id" : cart.ProductId,
					"product_name" : cart.ProductName,
					"quantity" : itemCheckout.Quantity,
					"price" : price,
				})
				//if success checkout reduce stock
				cart.Quantity -= itemCheckout.Quantity
				uc.Inventories[index].ProductStock -= itemCheckout.Quantity
				break
			}
		}
	}

	if len(dataCheckout) == 0 {
		return http.StatusNoContent, dataCheckout, errors.New("product id on cart not found")
	}

	return http.StatusOK, dataCheckout, nil

}

func GetIndexProductInventory( productId int, inventories []*entity.ProductInventory ) int {
	for i, val := range inventories {
		if productId == val.ProductId {
			return i
		}
	}

	return 0
}