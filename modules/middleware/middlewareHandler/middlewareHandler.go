package middlewarehandler

import (
	"github.com/naphat-sirisubkulchai/shop/config"
	"github.com/naphat-sirisubkulchai/shop/modules/middleware/middlewareUsecase"
)


type(
	MiddlewareHandlerService interface{}

	middlewareHandler struct{
		cfg *config.Config
		middlewareUsecase middlewareusecase.MiddlewareUsecaseService
	}
)

func NewMiddlewareHandler(cfg *config.Config, middlewareUsecase middlewareusecase.MiddlewareUsecaseService ) MiddlewareHandlerService{
	return &middlewareHandler{cfg,middlewareUsecase}
}