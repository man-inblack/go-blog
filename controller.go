package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func index(c *gin.Context) {
	dbClient, ctx := connectDb()
	findOptions := options.Find()
	findOptions.SetSort(bson.D{{"createTime", -1}})
	findOptions.SetLimit(5)
	cTest1 := dbClient.Database(Config.DB.Blog_db).Collection(Config.DB.Blog_text_col)
	cursor, _ := cTest1.Find(ctx, bson.M{}, findOptions)
	var result []bson.M
	if err := cursor.All(ctx, &result); err != nil {
		log.Fatal(err)
	}
	fmt.Println(result)
}

func home(w http.ResponseWriter, r *http.Request) {
	dbClient, ctx := connectDb()
	if r.Method == "GET" {

		cTest1 := dbClient.Database(Config.DB.Blog_db).Collection(Config.DB.Blog_text_col)
		cursor, err := cTest1.Find(ctx, bson.M{})
		var episodes []bson.M
		if err = cursor.All(ctx, &episodes); err != nil {
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
		_, err := cTest1.InsertOne(ctx, text)

		if err != nil {
			log.Fatal(err)
		}
		fmt.Fprintf(w, "home!! post")
	} else if r.Method == "DELETE" {
		fmt.Fprintf(w, "home!! delete")
	}
	defer dbClient.Disconnect(ctx)
}

// func test1(w http.ResponseWriter, r *http.Request) {
// 	content := make([]byte, r.ContentLength)
// 	r.Body.Read(content)
// 	getTextById()
// }

// func getTextById(textId primitive.ObjectID) {
// 	dbClient, ctx := connectDb(url)
// 	blogCol := dbClient.Database(blogDB).Collection(blogTextCol)
// 	cursor, err := blogCol.FindOne(ctx, bson.M{_id: textId})
// 	var episodes []bson.M
// 	if err = cursor.All(ctx, &episodes); err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Println(episodes)

// }
