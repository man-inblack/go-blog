package main

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

func connectDb() {
	log.Println(Config.DB.Server, Config.DB.Blog_db, Config.DB.Blog_text_col, "done!")

	// client, err := mongo.NewClient(options.Client().ApplyURI(Config.DB.Server))
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	clientOption := options.Client().ApplyURI(Config.DB.Server).SetConnectTimeout(10 * time.Second)
	var err error
	client, err = mongo.Connect(context.TODO(), clientOption)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	// err = client.Connect(ctx)
	// if err != nil {
	// 	log.Fatal(err)
	// }
}

func GetClient() *mongo.Client {
	if client == nil {
		connectDb()
	}
	return client
}
