package paymentusecase

import paymentrepository "github.com/naphat-sirisubkulchai/shop/modules/payment/paymentRepository"

type(
	PaymentUsecaseService interface{}
	paymentUsecase struct{
		paymentRepository paymentrepository.PaymentRepositoryService
	}

)
func NewPaymentUsecase(paymentRepository paymentrepository.PaymentRepositoryService) PaymentUsecaseService{
	return &paymentUsecase{
		paymentRepository:paymentRepository,
	}
}