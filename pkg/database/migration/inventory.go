package migration

import (
	"context"
	"log"

	"github.com/naphat-sirisubkulchai/shop/config"
	"github.com/naphat-sirisubkulchai/shop/pkg/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func inventoryDbConn(pctx context.Context, cfg *config.Config) *mongo.Database {
	return database.DbConn(pctx, cfg).Database("inventory_db")
}

func InventoryMigrate(pctx context.Context, cfg *config.Config) {
	db := inventoryDbConn(pctx, cfg)
	defer db.Client().Disconnect(pctx)

	col := db.Collection("users_inventory")

	indexs, _ := col.Indexes().CreateMany(pctx, []mongo.IndexModel{
		{Keys: bson.D{{"user_id", 1}, {"item_id", 1}}},
	})
	for _, index := range indexs {
		log.Printf("Index: %s", index)
	}

	col = db.Collection("users_inventory_queue")

	results, err := col.InsertOne(pctx, bson.M{"offset": -1}, nil)
	if err != nil {
		panic(err)
	}
	log.Println("Migrate inventory completed: ", results)
}