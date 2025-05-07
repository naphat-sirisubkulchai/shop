package server

import (
	inventoryhandler "github.com/naphat-sirisubkulchai/shop/modules/inventory/inventoryHandler"
	inventoryrepository "github.com/naphat-sirisubkulchai/shop/modules/inventory/inventoryRepository"
	inventoryusecase "github.com/naphat-sirisubkulchai/shop/modules/inventory/inventoryUsecase"
)

func (s *server) inventoryService() {
	repo := inventoryrepository.NewInventoryRepository(s.db)
	usecase := inventoryusecase.NewInventoryUsecase(repo)
	httpHandler := inventoryhandler.NewInventoryHttpHandler(s.cfg, usecase)
	grpcHandler := inventoryhandler.NewInventoryGrpcHandler(usecase)

	_=httpHandler
	_=grpcHandler
	inventory := s.app.Group("/inventory_v1")
	inventory.GET("",s.healthCheckService)
}
	