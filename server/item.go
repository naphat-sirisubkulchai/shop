package server

import (
	"log"

	itemhandler "github.com/naphat-sirisubkulchai/shop/modules/item/itemHandler"
	itemPb "github.com/naphat-sirisubkulchai/shop/modules/item/itemPb"
	itemrepository "github.com/naphat-sirisubkulchai/shop/modules/item/itemRepository"
	itemusecase "github.com/naphat-sirisubkulchai/shop/modules/item/itemUsecase"
	"github.com/naphat-sirisubkulchai/shop/pkg/grpcconn"
)

func (s *server) itemService() {
	repo := itemrepository.NewItemRepository(s.db)
	usecase := itemusecase.NewItemUsecase(repo)
	httpHandler := itemhandler.NewItemHttpHandler(s.cfg, usecase)
	grpcHandler := itemhandler.NewItemGrpcHandler(usecase)

	go func() {
		grpcServer, lis := grpcconn.NewGrpcServer(&s.cfg.Jwt, s.cfg.Grpc.ItemUrl)

		itemPb.RegisterItemGrpcServiceServer(grpcServer, grpcHandler)

		log.Printf("Item gRPC server listening on %s", s.cfg.Grpc.ItemUrl)
		grpcServer.Serve(lis)
	}()


	_=httpHandler
	_=grpcHandler
	item := s.app.Group("/item_v1")
	item.GET("",s.healthCheckService)
}