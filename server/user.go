package server

import (
	"log"

	userhandler "github.com/naphat-sirisubkulchai/shop/modules/user/userHandler"
	userrepository "github.com/naphat-sirisubkulchai/shop/modules/user/userRepository"
	userusecase "github.com/naphat-sirisubkulchai/shop/modules/user/userUsecase"
	userPb "github.com/naphat-sirisubkulchai/shop/modules/user/userPb"
	"github.com/naphat-sirisubkulchai/shop/pkg/grpcconn"
)

func (s *server) userService() {
	repo := userrepository.NewUserRepository(s.db)
	usecase := userusecase.NewUserUsecase(repo)
	httpHandler := userhandler.NewUserHttpHandler(s.cfg, usecase)
	grpcHandler := userhandler.NewUserGrpcHandler(usecase)
	queueHandler := userhandler.NewUserQueueHandler(s.cfg,usecase)

	go func() {
		grpcServer, lis := grpcconn.NewGrpcServer(&s.cfg.Jwt, s.cfg.Grpc.UserUrl)

		userPb.RegisterUserGrpcServiceServer(grpcServer, grpcHandler)

		log.Printf("User gRPC server listening on %s", s.cfg.Grpc.UserUrl)
		grpcServer.Serve(lis)
	}()


	_=grpcHandler
	_=queueHandler
	user := s.app.Group("/user_v1")
	user.GET("",s.healthCheckService)
	user.POST("/user/register",httpHandler.CreateUser)
	user.GET("/user/:user_id",httpHandler.FindOneUserProfile)
	user.POST("/user/add-money",httpHandler.AddUserMoney)
	user.GET("/user/account/:user_id",httpHandler.GetUserSavingAccount)
}