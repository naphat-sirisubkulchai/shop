package server

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/naphat-sirisubkulchai/shop/config"
	middlewarehandler "github.com/naphat-sirisubkulchai/shop/modules/middleware/middlewareHandler"
	middlewarerepository "github.com/naphat-sirisubkulchai/shop/modules/middleware/middlewareRepository"
	middlewareusecase "github.com/naphat-sirisubkulchai/shop/modules/middleware/middlewareUsecase"
	"go.mongodb.org/mongo-driver/mongo"
)

type(
	server struct{
		app *echo.Echo
		db *mongo.Client
		cfg *config.Config
		middleware middlewarehandler.MiddlewareHandlerService

	}
)

func newMiddleware(cfg *config.Config) middlewarehandler.MiddlewareHandlerService{
	repo := middlewarerepository.NewMiddlewareRepository()
	usecase := middlewareusecase.NewMiddlewareUsecase(repo)
	return middlewarehandler.NewMiddlewareHandler(cfg,usecase)
}

func (s *server)gracefulShutdown(pctx context.Context, quit <-chan os.Signal){
	log.Println("Start service: %s",s.cfg.App.Name)
	<-quit
	log.Println("Shutting service %s", s.cfg.App.Name)

	ctx,cancel := context.WithTimeout(pctx,10*time.Second)
	defer cancel()

	if err:= s.app.Shutdown(ctx);err!= nil{
		log.Fatalf("Error: %v",err)
	}
}
func (s *server)httpListening(){
	if err:= s.app.Start(s.cfg.App.Url); err!=nil&& err!= http.ErrServerClosed{
		log.Fatal("Error: %v",err)
	}
}

func Start(pctx context.Context,cfg *config.Config, db *mongo.Client){
	s:=&server{
		app: echo.New(),
		db:db,
		cfg:cfg,
		middleware: newMiddleware(cfg),
	}

	s.app.Use(middleware.TimeoutWithConfig(middleware.TimeoutConfig{
		Skipper: middleware.DefaultSkipper,
		ErrorMessage:"Error: Request Timeout",
		Timeout: 30* time.Second,
	}))

	s.app.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		Skipper: middleware.DefaultSkipper,
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET,echo.POST,echo.PUT,echo.PATCH,echo.DELETE},
	}))

	s.app.Use(middleware.BodyLimit("2M"))

	switch s.cfg.App.Name{	
	case "auth":
		s.authService()
	case "user":
		s.userService()
	case "item":
		s.itemService()
	case "inventory":
		s.inventoryService()
	case "payment":
		s.paymentService()
	}

	quit := make(chan os.Signal,1)
	signal.Notify(quit,syscall.SIGINT,syscall.SIGTERM)
	
	s.app.Use(middleware.Logger())
	
	go s.gracefulShutdown(pctx,quit)

	s.httpListening()
}