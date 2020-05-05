package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

//任务的执行时间
type TimePoint struct {
	StartTime int64 `bson:"startTime"`
	EndTime   int64 `bson:"endTime"`
}

//一条日志
type LogRecord struct {
	JobName   string    `bson:"job_name"`  //任务名
	Command   string    `bson:"command"`   //shell命令
	Err       string    `bson:"err"`       //脚本错误
	Content   string    `bson:"content"`   //脚本输出
	TimePoint TimePoint `bson:"timePoint"` //执行时间
}

func main() {
	var (
		client     *mongo.Client
		err        error
		database   *mongo.Database
		collection *mongo.Collection
		recod      *LogRecord
		result     *mongo.InsertOneResult
		docld      primitive.ObjectID
	)

	// 设置客户端连接配置
	// 连接到MongoDB
	if client, err = mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017")); err != nil {
		fmt.Println(err)
	}

	//选择数据库my_db
	database = client.Database("cron")

	// 3.选择表
	collection = database.Collection("log")

	//4.插入记录
	recod = &LogRecord{
		JobName: "job10",
		Command: "echo hello",
		Err:     "",
		Content: "hello",
		TimePoint: TimePoint{
			StartTime: time.Now().Unix(),
			EndTime:   time.Now().Unix() + 10},
	}
	if result, err = collection.InsertOne(context.TODO(), recod); err != nil {
		fmt.Println(err)
		return
	}
	//_id默认生成一个全局ID,Object 12字节的二进制
	docld = result.InsertedID.(primitive.ObjectID)
	fmt.Println("自增ID:", docld.Hex())
}
