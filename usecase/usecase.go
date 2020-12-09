package usecase

import(
	"assessment-online-store/entity"
	"context"
)


//Init Use case interface
type InterfaceUseCase interface {
	GetListProductUseCase() []*entity.ProductInventory
}

type UseCase struct {
	Context context.Context
	Inventories []*entity.ProductInventory
}

func NewUseCase(ctx context.Context, inventories []*entity.ProductInventory) InterfaceUseCase {
	return &UseCase{
		Context: ctx,
		Inventories : inventories,
	}
}

func (uc *UseCase)GetListProductUseCase() []*entity.ProductInventory {
	return uc.Inventories
}