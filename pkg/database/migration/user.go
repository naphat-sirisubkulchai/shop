package migration

import (
	"context"
	"log"

	"github.com/naphat-sirisubkulchai/shop/config"
	"github.com/naphat-sirisubkulchai/shop/modules/user"
	"github.com/naphat-sirisubkulchai/shop/pkg/database"
	"github.com/naphat-sirisubkulchai/shop/pkg/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

func userDbConn(pctx context.Context, cfg *config.Config) *mongo.Database {
	return database.DbConn(pctx, cfg).Database("user_db")
}

func UserMigrate(pctx context.Context, cfg *config.Config) {
	db := userDbConn(pctx, cfg)
	defer db.Client().Disconnect(pctx)

	col := db.Collection("user_transactions")

	// indexs
	indexs, _ := col.Indexes().CreateMany(pctx, []mongo.IndexModel{
		{Keys: bson.D{{"_id", 1}}},
		{Keys: bson.D{{"user_id", 1}}},
	})
	log.Println(indexs)

	col = db.Collection("users")

	// indexs
	indexs, _ = col.Indexes().CreateMany(pctx, []mongo.IndexModel{
		{Keys: bson.D{{"_id", 1}}},
		{Keys: bson.D{{"email", 1}}},
	})
	log.Println(indexs)

	documents := func() []any {
		roles := []*user.User{
			{
				Email: "user1@korkor.com",
				Password: func() string {
					// Hashing password
					hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("123456"), bcrypt.DefaultCost)
					return string(hashedPassword)
				}(),
				Username: "User1",
				UserRoles: []user.UserRole{
					{
						RoleTitle: "user",
						RoleCode:  0,
					},
				},
				CreatedAt: utils.LocalTime(),
				UpdatedAt: utils.LocalTime(),
			},
			{
				Email: "user2@korkor.com",
				Password: func() string {
					// Hashing password
					hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("123456"), bcrypt.DefaultCost)
					return string(hashedPassword)
				}(),
				Username: "User2",
				UserRoles: []user.UserRole{
					{
						RoleTitle: "user",
						RoleCode:  0,
					},
				},
				CreatedAt: utils.LocalTime(),
				UpdatedAt: utils.LocalTime(),
			},
			{
				Email: "user3@korkor.com",
				Password: func() string {
					// Hashing password
					hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("123456"), bcrypt.DefaultCost)
					return string(hashedPassword)
				}(),
				Username: "User3",
				UserRoles: []user.UserRole{
					{
						RoleTitle: "user",
						RoleCode:  0,
					},
				},
				CreatedAt: utils.LocalTime(),
				UpdatedAt: utils.LocalTime(),
			},
			{
				Email: "admin1@korkor.com",
				Password: func() string {
					// Hashing password
					hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("123456"), bcrypt.DefaultCost)
					return string(hashedPassword)
				}(),
				Username: "User3",
				UserRoles: []user.UserRole{
					{
						RoleTitle: "user",
						RoleCode:  0,
					},
					{
						RoleTitle: "admin",
						RoleCode:  1,
					},
				},
				CreatedAt: utils.LocalTime(),
				UpdatedAt: utils.LocalTime(),
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
	log.Println("Migrate auth completed: ", results)

	userTransactions := make([]any, 0)
	for _, p := range results.InsertedIDs {
		userTransactions = append(userTransactions, &user.UserTransaction{
			UserId:  "user:" + p.(primitive.ObjectID).Hex(),
			Amount:    1000,
			CreatedAt: utils.LocalTime(),
		})
	}

	col = db.Collection("user_transactions")
	results, err = col.InsertMany(pctx, userTransactions, nil)
	if err != nil {
		panic(err)
	}
	log.Println("Migrate user_transactions completed: ", results)

	col = db.Collection("user_transactions_queue")
	result, err := col.InsertOne(pctx, bson.M{"offset": -1}, nil)
	if err != nil {
		panic(err)
	}
	log.Println("Migrate user_transactions_queue completed: ", result)
}