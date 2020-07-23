package main

import "go.mongodb.org/mongo-driver/bson/primitive"

type blogText struct {
	Title      string             `bson:"title,omitempty"`
	UpdateTime int                `bson:"updateTime,omitempty"`
	CreateTime int                `bson:"createTime,omitempty"`
	Content    string             `bson:"content,omitempty"`
	ViewCount  int                `bson:"viewCount,omitempty"`
	Category   []string           `bson:"category,omitempty"`
	Author     string             `bson:"author,omitempty"`
	ID         primitive.ObjectID `bson:"_id,omitempty"`
}
