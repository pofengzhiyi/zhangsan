package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	var (
		client     *mongo.Client
		err        error
		database   *mongo.Database
		collection *mongo.Collection
	)
	// 设置客户端连接配置
	// 连接到MongoDB
	if client, err = mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017")); err != nil {
		fmt.Println(err)
	}

	//选择数据库my_db
	database = client.Database("my_db")

	// 3.选择表
	collection = database.Collection("my_collection")

	collection = collection

}
