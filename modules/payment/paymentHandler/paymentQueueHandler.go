package paymenthandler

import (
	"github.com/naphat-sirisubkulchai/shop/config"
	"github.com/naphat-sirisubkulchai/shop/modules/payment/paymentUsecase"
)

type(

	PaymentQueueHandlerService interface{}

	paymentQueueHandler struct{
		cfg *config.Config
		paymentUsecase paymentusecase.PaymentUsecaseService
	}
)

func NewPaymentQueueHandler(cfg *config.Config, 	paymentUsecase paymentusecase.PaymentUsecaseService) PaymentQueueHandlerService {
	return &paymentQueueHandler{
		cfg:cfg,
		paymentUsecase: paymentUsecase,
	}
}
