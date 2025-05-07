package server

import (
	"log"
	authhandler "github.com/naphat-sirisubkulchai/shop/modules/auth/authHandler"
	authrepository "github.com/naphat-sirisubkulchai/shop/modules/auth/authRepository"
	authusecase "github.com/naphat-sirisubkulchai/shop/modules/auth/authUsecase"
	authPb "github.com/naphat-sirisubkulchai/shop/modules/auth/authPb"
	"github.com/naphat-sirisubkulchai/shop/pkg/grpcconn"
)
func (s *server) authService() {
	repo := authrepository.NewAuthRepository(s.db)
	usecase := authusecase.NewAuthUsecase(repo)
	httpHandler := authhandler.NewAuthHttpHandler(s.cfg, usecase)
	grpcHandler := authhandler.NewAuthGrpcHandler(usecase)


	go func() {
		grpcServer, lis := grpcconn.NewGrpcServer(&s.cfg.Jwt, s.cfg.Grpc.AuthUrl)

		authPb.RegisterAuthGrpcServiceServer(grpcServer, grpcHandler)

		log.Printf("Auth gRPC server listening on %s", s.cfg.Grpc.AuthUrl)
		grpcServer.Serve(lis)
	}()

	_=httpHandler
	_=grpcHandler
	auth := s.app.Group("/auth_v1")
	auth.GET("",s.healthCheckService)

	auth.POST("/auth/login", httpHandler.Login)
	auth.POST("/auth/refresh-token", httpHandler.RefreshToken)
	auth.POST("/auth/logout", httpHandler.Logout)
}
	