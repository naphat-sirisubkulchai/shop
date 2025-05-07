package inventoryhandler

import inventoryusecase "github.com/naphat-sirisubkulchai/shop/modules/inventory/inventoryUsecase"

type(
	inventoryGrpcHandler struct{
		inventoryUsecase inventoryusecase.InventoryUsecaseService
	}
)

func NewInventoryGrpcHandler(inventoryUsecase inventoryusecase.InventoryUsecaseService) *inventoryGrpcHandler{
	return &inventoryGrpcHandler{
		inventoryUsecase:inventoryUsecase,
	}
}