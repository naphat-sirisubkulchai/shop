package itemhandler

import (
	"github.com/naphat-sirisubkulchai/shop/config"
	itemusecase "github.com/naphat-sirisubkulchai/shop/modules/item/itemUsecase"
)

type(
	ItemHttpHandlerService interface {}
	itemHttpHandler struct{
		cfg *config.Config
		itemUsecase itemusecase.ItemUsecaseService
	}

)

func NewItemHttpHandler(cfg *config.Config,itemUsecase itemusecase.ItemUsecaseService) ItemHttpHandlerService{
	return &itemHttpHandler{
		cfg:cfg,
		itemUsecase: itemUsecase,
	}

}