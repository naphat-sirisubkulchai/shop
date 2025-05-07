package server

import (
	paymenthandler "github.com/naphat-sirisubkulchai/shop/modules/payment/paymentHandler"
	paymentrepository "github.com/naphat-sirisubkulchai/shop/modules/payment/paymentRepository"
	paymentusecase "github.com/naphat-sirisubkulchai/shop/modules/payment/paymentUsecase"
)

func (s *server) paymentService() {
	repo := paymentrepository.NewPaymentRepository(s.db)
	usecase := paymentusecase.NewPaymentUsecase(repo)
	httpHandler := paymenthandler.NewPaymentHttpHandler(s.cfg, usecase)
	grpcHandler := paymenthandler.NewPaymentQueueHandler(s.cfg,usecase)

	_=httpHandler
	_=grpcHandler
	payment := s.app.Group("/payment_v1")
	payment.GET("",s.healthCheckService)
}