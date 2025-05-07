package migration

import (
	"context"
	"log"

	"github.com/naphat-sirisubkulchai/shop/config"
	"github.com/naphat-sirisubkulchai/shop/modules/item"
	"github.com/naphat-sirisubkulchai/shop/pkg/database"
	"github.com/naphat-sirisubkulchai/shop/pkg/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func itemDbConn(pctx context.Context, cfg *config.Config) *mongo.Database {
	return database.DbConn(pctx, cfg).Database("item_db")
}

func ItemMigrate(pctx context.Context, cfg *config.Config) {
	db := itemDbConn(pctx, cfg)
	defer db.Client().Disconnect(pctx)

	col := db.Collection("items")
	indexs, _ := col.Indexes().CreateMany(pctx, []mongo.IndexModel{
		{Keys: bson.D{{"_id", 1}}},
		{Keys: bson.D{{"title", 1}}},
	})
	for _, index := range indexs {
		log.Printf("Index: %s", index)
	}

	documents := func() []any {
		roles := []*item.Item{
			{
				Title:       "Guitar1",
				Price:       1000,
				ImageUrl:    "https://i.imgur.com/oqZpxBu.png",
				UsageStatus: true,
				Defect:      "mee tam ni",
				CreatedAt:   utils.LocalTime(),
				UpdatedAt:   utils.LocalTime(),
			},
			{
				Title:       "GuitarProMax",
				Price:       500,
				ImageUrl:    "https://i.imgur.com/oqZpxBu.png",
				UsageStatus: true,
				Defect:      "mee tam ni u na",
				CreatedAt:   utils.LocalTime(),
				UpdatedAt:   utils.LocalTime(),
			},
			{
				Title:       "Piano",
				Price:       100,
				ImageUrl:    "https://i.imgur.com/oqZpxBu.png",
				UsageStatus: true,
				Defect:      "me tam ni nid noey",
				CreatedAt:   utils.LocalTime(),
				UpdatedAt:   utils.LocalTime(),
			},
		}

		docs := make([]any, 0)
		for _, r := range roles {
			docs = append(docs, r)
		}
		return docs
	}()

	results, err := col.InsertMany(pctx, documents, nil)
	if err != nil {
		panic(err)
	}
	log.Println("Migrate item completed: ", results)
}