package main

import (
	"context"
	"fmt"
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

//jobName过滤条件
type FindByJobName struct {
	JobName string `bson:"job_name"` //	JobName赋值为job10
}

//startTime小于某时间
//{"$lt":timestamp}
type TimeBeforeCond struct {
	Before int64 `bson:"$lt"`
}

//{"timePoint.startTime":{"$lt“:timestamp}
type DeleteCond struct {
	beforeCond TimeBeforeCond `bson:"timePoint.startTime"`
}

func main() {

	//mongodb读取回来时是bson，需要反序列化为LogRecord对象
	var (
		client     *mongo.Client
		err        error
		database   *mongo.Database
		collection *mongo.Collection
		delCond    *DeleteCond
		delResult  *mongo.DeleteResult
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

	//4.要删除开始时间早于当前时间的所有日志($lt是less than）
	//delete（{timePoint.startTime":{"$lt":当前时间}})
	delCond = &DeleteCond{beforeCond: TimeBeforeCond{Before: time.Now().Unix()}}

	//执行
	if delResult, err = collection.DeleteMany(context.TODO(), delCond); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("删除的行数", delResult.DeletedCount)
}
