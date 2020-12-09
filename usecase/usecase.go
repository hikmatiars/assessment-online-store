package usecase

import(
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
	index := req.ProductId
	if req.Quantity >= uc.Inventories[index].ProductStock {
		return http.StatusUnprocessableEntity, errors.New("product out of stock")
	}

	uc.Inventories[index].ProductStock -= req.Quantity

	uc.Cart = append(uc.Cart, &entity.Cart{
		ProductId: uc.Inventories[index].ProductId,
		ProductName: uc.Inventories[index].ProductName,
		Quantity: uc.Inventories[index].ProductStock,
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