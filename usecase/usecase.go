package usecase

import (
	"assessment-online-store/entity"
	"assessment-online-store/http/request"
	"context"
	"errors"
	"net/http"
)

//Init Use case interface
type InterfaceUseCase interface {
	GetListProductUseCase() []*entity.ProductInventory
	AddCartUseCase(req request.AddToCart ) (int, error)
	GetListCartUseCase() (int, []*entity.Cart)
}

type UseCase struct {
	Context context.Context
	Inventories []*entity.ProductInventory
	Cart []*entity.Cart
}

func NewUseCase(ctx context.Context, inventories []*entity.ProductInventory, cart []*entity.Cart) InterfaceUseCase {
	return &UseCase{
		Context: ctx,
		Inventories : inventories,
		Cart: cart,
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
	})

	return http.StatusOK, nil
}

func (uc *UseCase) GetListCartUseCase() (int, []*entity.Cart) {
	if len(uc.Cart) == 0 {
		return http.StatusNoContent, uc.Cart
	}

	return http.StatusOK, uc.Cart
}