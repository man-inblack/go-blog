package main

import (
	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})
	r.GET("/", index)

	// fmt.Println("blog begin  ...")
	// http.HandleFunc("/", home)
	// http.HandleFunc("/blog", test1)
	// http.ListenAndServe(":8086", nil)
	//defer MongoClient.Disconnect(Ctx)
	r.Run(":8086")
}

// func home(w http.ResponseWriter, r *http.Request) {
// 	dbClient, ctx := connectDb(url)
// 	if r.Method == "GET" {

// 		cTest1 := dbClient.Database(blogDB).Collection(blogTextCol)
// 		cursor, err := cTest1.Find(ctx, bson.M{})
// 		var episodes []bson.M
// 		if err = cursor.All(ctx, &episodes); err != nil {
// 			log.Fatal(err)
// 		}
// 		fmt.Println(episodes)
// 		fmt.Fprintf(w, "hello put")
// 	} else if r.Method == "PUT" {
// 		fmt.Fprintf(w, "home!! put")
// 	} else if r.Method == "POST" {
// 		text := blogText{
// 			Title:      "第er篇文章",
// 			UpdateTime: int(time.Now().Unix() * 1000),
// 			CreateTime: int(time.Now().Unix() * 1000),
// 			Content:    "nsjdhksjdlkajsd",
// 			ViewCount:  67,
// 			Category:   []string{"tech", "first"},
// 			Author:     "vibe",
// 		}

// 		cTest1 := dbClient.Database(blogDB).Collection(blogTextCol)
// 		_, err := cTest1.InsertOne(ctx, text)

// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 		fmt.Fprintf(w, "home!! post")
// 	} else if r.Method == "DELETE" {
// 		fmt.Fprintf(w, "home!! delete")
// 	}
// 	defer dbClient.Disconnect(ctx)
// }

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
