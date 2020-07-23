package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type IndexDoc struct {
	Title      string   `bson:"title,omitempty"`
	UpdateTime int      `bson:"updateTime,omitempty"`
	CreateTime int      `bson:"createTime,omitempty"`
	Content    string   `bson:"content,omitempty"`
	ViewCount  int      `bson:"viewCount,omitempty"`
	Category   []string `bson:"category,omitempty"`
	Author     string   `bson:"author,omitempty"`
	ID         string   `bson:"_id"`
}

func index(c *gin.Context) {

	dbClient := GetClient()
	findOptions := options.Find()
	findOptions.SetSort(bson.D{{"createTime", -1}})
	findOptions.SetLimit(5)
	cTest1 := dbClient.Database(Config.DB.Blog_db).Collection(Config.DB.Blog_text_col)
	cursor, _ := cTest1.Find(context.TODO(), bson.M{}, findOptions)

	var doc []IndexDoc
	for cursor.Next(context.TODO()) {
		var temp IndexDoc
		err := cursor.Decode(&temp)
		if err != nil {
			log.Fatal(err)
		}
		doc = append(doc, temp)

	}

	c.IndentedJSON(200, doc)
	fmt.Println(doc)

}

func home(w http.ResponseWriter, r *http.Request) {
	dbClient := GetClient()
	if r.Method == "GET" {

		cTest1 := dbClient.Database(Config.DB.Blog_db).Collection(Config.DB.Blog_text_col)
		cursor, err := cTest1.Find(context.TODO(), bson.M{})
		var episodes []bson.M
		if err = cursor.All(context.TODO(), &episodes); err != nil {
			log.Fatal(err)
		}
		fmt.Println(episodes)
		fmt.Fprintf(w, "hello put")
	} else if r.Method == "PUT" {
		fmt.Fprintf(w, "home!! put")
	} else if r.Method == "POST" {
		text := blogText{
			Title:      "第er篇文章",
			UpdateTime: int(time.Now().Unix() * 1000),
			CreateTime: int(time.Now().Unix() * 1000),
			Content:    "nsjdhksjdlkajsd",
			ViewCount:  67,
			Category:   []string{"tech", "first"},
			Author:     "vibe",
		}

		cTest1 := dbClient.Database(Config.DB.Blog_db).Collection(Config.DB.Blog_text_col)
		_, err := cTest1.InsertOne(context.TODO(), text)

		if err != nil {
			log.Fatal(err)
		}
		fmt.Fprintf(w, "home!! post")
	} else if r.Method == "DELETE" {
		fmt.Fprintf(w, "home!! delete")
	}
	defer dbClient.Disconnect(context.TODO())
}
