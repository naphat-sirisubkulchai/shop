package main

import (
	"context"
	"log"
	"os"

	"github.com/naphat-sirisubkulchai/shop/config"
	"github.com/naphat-sirisubkulchai/shop/pkg/database"
	"github.com/naphat-sirisubkulchai/shop/server"
)
func main(){
	ctx := context.Background() 
	_= ctx 

	//Initialize config 
	cfg := config.LoadConfig(func () string{
		 if len(os.Args) <2{
			log.Fatal("Error .env path is required ")
		 }
		 return os.Args[1] 
	}())

	db := database.DbConn(ctx, &cfg)
	log.Println(db)
	defer db.Disconnect(ctx)
	log.Println(cfg)

	server.Start(ctx,&cfg,db)
}