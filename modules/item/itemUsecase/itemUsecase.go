package itemusecase

import itemrepository "github.com/naphat-sirisubkulchai/shop/modules/item/itemRepository"

type(
	ItemUsecaseService interface {}
	itemUsecase struct{
		itemRepository itemrepository.ItemRepositoryService
	}

)

func NewItemUsecase(itemRepository itemrepository.ItemRepositoryService)ItemUsecaseService{
	return &itemUsecase{itemRepository}

}