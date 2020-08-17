package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/gin-gonic/gin"
	"github.com/gomarkdown/markdown"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type text struct {
	Title      string   `bson:"title,omitempty"`
	UpdateTime int      `bson:"updateTime,omitempty"`
	CreateTime int      `bson:"createTime,omitempty"`
	Content    string   `bson:"content,omitempty"`
	ViewCount  int      `bson:"viewCount,omitempty"`
	Category   []string `bson:"category,omitempty"`
	Author     string   `bson:"author,omitempty"`
	ID         string   `bson:"_id"`
	// IsDelete   bool     `bson:"isDelete,omitempty"`
}

type textList struct {
	Page     int64
	PageSize int64
	Total    int64
	Data     []text
}

func index(c *gin.Context) {

	dbClient := GetClient()
	findOptions := options.Find()
	findOptions.SetSort(bson.M{"createTime": -1})
	findOptions.SetLimit(5)
	cTest1 := dbClient.Database(Config.DB.Blog_db).Collection(Config.DB.Blog_text_col)
	cursor, _ := cTest1.Find(context.TODO(), bson.M{}, findOptions)

	var doc []text
	for cursor.Next(context.TODO()) {
		var temp text
		err := cursor.Decode(&temp)
		if err != nil {
			log.Fatal(err)
		}
		temp.Content = string(markdown.ToHTML([]byte(temp.Content), nil, nil))
		doc = append(doc, temp)

	}

	c.JSON(200, doc)
	fmt.Println(doc)

}

func getTextByID(c *gin.Context) {
	textID, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		log.Fatal(err)
	}
	dbClient := GetClient()
	blogTextClo := dbClient.Database(Config.DB.Blog_db).Collection(Config.DB.Blog_text_col)
	var result text
	err = blogTextClo.FindOne(context.TODO(), bson.M{"_id": textID}).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}
	result.Content = string(markdown.ToHTML([]byte(result.Content), nil, nil))
	c.JSON(200, result)
}

func getTextList(c *gin.Context) {

	isPage := c.DefaultQuery("is_page", "true") == "true"
	page, _ := strconv.ParseInt(c.DefaultQuery("page", "1"), 10, 64)
	pageSize, _ := strconv.ParseInt(c.DefaultQuery("page_size", "10"), 10, 64)

	findOptions := options.Find()
	findOptions.SetSort(bson.M{"createTime": -1})
	if isPage {
		findOptions.SetSkip((page - 1) * pageSize)
		findOptions.SetLimit(pageSize)
	}
	dbClient := GetClient()
	cTest1 := dbClient.Database(Config.DB.Blog_db).Collection(Config.DB.Blog_text_col)
	cursor, _ := cTest1.Find(context.TODO(), bson.M{"isDelete": false}, findOptions)
	textCount, _ := cTest1.CountDocuments(context.TODO(), bson.M{"isDelete": false})
	var doc []text
	for cursor.Next(context.TODO()) {
		var temp text
		err := cursor.Decode(&temp)
		if err != nil {
			log.Fatal(err)
		}
		doc = append(doc, temp)

	}
	result := textList{page, pageSize, textCount, doc}
	c.JSON(200, result)
	fmt.Println(result)
}

func insertBlogText(c *gin.Context) {
	
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
