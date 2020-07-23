package main

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func connectDb() (*mongo.Client, context.Context) {
	log.Println(Config.DB.Server, Config.DB.Blog_db, Config.DB.Blog_text_col, "done!")
	client, err := mongo.NewClient(options.Client().ApplyURI(Config.DB.Server))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	return client, ctx
}
