package middlewareusecase

import middlewarerepository "github.com/naphat-sirisubkulchai/shop/modules/middleware/middlewareRepository"

type(
	MiddlewareUsecaseService interface{}

	middlewareUsecase struct{
		middlewareRepository middlewarerepository.MiddlewareRepositoryService
	}
)

func NewMiddlewareUsecase(middlewareRepository middlewarerepository.MiddlewareRepositoryService) MiddlewareUsecaseService {
	return &middlewareUsecase{middlewareRepository}
}