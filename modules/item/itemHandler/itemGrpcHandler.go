package itemhandler

import (
	"context"
	itemPb		"github.com/naphat-sirisubkulchai/shop/modules/item/itemPb"
	itemusecase "github.com/naphat-sirisubkulchai/shop/modules/item/itemUsecase"
)


type(
	itemGrpcHandler struct{
		itemUsecase itemusecase.ItemUsecaseService
		itemPb.UnimplementedItemGrpcServiceServer
	}

)

func NewItemGrpcHandler(itemUsecase itemusecase.ItemUsecaseService) *itemGrpcHandler{
	return &itemGrpcHandler{
		itemUsecase: itemUsecase,
	}
}
func (g *itemGrpcHandler) FindItemsInIds(ctx context.Context, req *itemPb.FindItemsInIdsReq) (*itemPb.FindItemsInIdsRes, error) {
	return nil,nil
}