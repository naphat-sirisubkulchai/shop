package inventoryhandler

import (
	"github.com/naphat-sirisubkulchai/shop/config"
	inventoryusecase "github.com/naphat-sirisubkulchai/shop/modules/inventory/inventoryUsecase"
)

type(
	InventoryHttpHandlerService interface{}

	inventoryHttpHandler struct{
		cfg *config.Config
		inventoryUsecase inventoryusecase.InventoryUsecaseService
	}
)

func NewInventoryHttpHandler(cfg *config.Config,inventoryUsecase inventoryusecase.InventoryUsecaseService) InventoryHttpHandlerService{
	return &inventoryHttpHandler{
		cfg :cfg,
		inventoryUsecase: inventoryUsecase,
	}
}