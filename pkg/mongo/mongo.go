package mongo

import (
	"context"
	"go_learn_web/configs"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	FACTORY_INFO_COLL       string = "factoryInfo"
	FACTORY_DATA_COLL       string = "factoryData"
	FACTORY_DATA_COUNT_COLL string = "factoryDataCount"
)

var mgo *mongo.Client
var FactoryInfoColl *mongo.Collection
var FactoryDataColl *mongo.Collection
var FactoryDataCountColl *mongo.Collection

func init() {
	// 设置客户端连接配置
	clientOptions := options.Client().ApplyURI("mongodb://" + configs.MongoHost + ":" + configs.MongoPort)

	// 连接到MongoDB
	var err error
	mgo, err = mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		panic(err.Error())
	}

	// 检查连接
	if err := mgo.Ping(context.Background(), nil); err != nil {
		panic(err.Error())
	}

	FactoryInfoColl = mgo.Database(configs.MongoDB).Collection(FACTORY_INFO_COLL)
	FactoryDataColl = mgo.Database(configs.MongoDB).Collection(FACTORY_DATA_COLL)
	FactoryDataCountColl = mgo.Database(configs.MongoDB).Collection(FACTORY_DATA_COUNT_COLL)
}
