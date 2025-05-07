package inventoryusecase

import (
	inventoryrepository "github.com/naphat-sirisubkulchai/shop/modules/inventory/inventoryRepository"
)

type(
	InventoryUsecaseService interface{}

	inventoryUsecase struct{
		inventoryRepository inventoryrepository.InventoryRepositoryService
	}
)

func NewInventoryUsecase(inventoryRepository inventoryrepository.InventoryRepositoryService) InventoryUsecaseService{
	return &inventoryUsecase{
		inventoryRepository: inventoryRepository,
	}
}