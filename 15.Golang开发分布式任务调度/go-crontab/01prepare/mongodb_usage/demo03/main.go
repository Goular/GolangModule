package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

// 任务的执行时间点
type TimePoint struct {
	StartTime int64 `bson:"startTime"`
	EndTime   int64 `bson:"endTime"`
}

// 一条日志
type LogRecord struct {
	JobName   string    `bson:"jobName"`   // 任务名
	Command   string    `bson:"command"`   // shell命令
	Err       string    `bson:"err"`       // 脚本错误
	Content   string    `bson:"content"`   // 脚本输出
	TimePoint TimePoint `bson:"timePoint"` // 执行时间点
}

func main() {
	var (
		client     *mongo.Client
		err        error
		database   *mongo.Database
		collection *mongo.Collection
		record     *LogRecord
		logArr     []interface{} //  C语言里的addr, type, JAVA Object
		result     *mongo.InsertManyResult
		insertId   interface{} //  objectId
		docId      primitive.ObjectID
	)
	// 1.建立连接
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	if client, err = mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017")); err != nil {
		fmt.Println(err)
		return
	}

	// 2.选择数据库my_db
	database = client.Database("test_db")
	collection = database.Collection("log")

	// 3. 创建object
	record = &LogRecord{
		JobName:   "job10888",
		Command:   "echo hello888",
		Err:       "888",
		Content:   "hello888",
		TimePoint: TimePoint{StartTime: time.Now().Unix(), EndTime: time.Now().Unix() + 10},
	}

	// 4. 批量插入记录(document)
	logArr = []interface{}{record, record, record}

	//发起插入
	if result, err = collection.InsertMany(context.TODO(), logArr); err != nil {
		fmt.Println(err)
		return
	}
	// 遍历读取
	for _, insertId = range result.InsertedIDs {
		docId = insertId.(primitive.ObjectID)
		fmt.Println("自增ID:", docId.Hex())
	}
}
