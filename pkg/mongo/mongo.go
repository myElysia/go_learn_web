package mongo

import (
	log "github.com/sirupsen/logrus"

	"context"
	"go_learn_web/configs"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	FACTORY_INFO_COLL       string = "bookInfo"
	FACTORY_DATA_COLL       string = "factoryData"
	FACTORY_DATA_COUNT_COLL string = "factoryDataCount"
)

var mgo *mongo.Client
var FactoryInfoColl *mongo.Collection
var FactoryDataColl *mongo.Collection
var FactoryDataCountColl *mongo.Collection

func Init() {
	// 设置客户端连接配置
	clientOptions := options.Client().ApplyURI("mongodb://" + configs.MongoHost + ":" + configs.MongoPort)

	// 连接到MongoDB
	var err error
	mgo, err = mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Panic(err.Error())
	}

	// 检查连接
	if err := mgo.Ping(context.Background(), nil); err != nil {
		log.Panic(err.Error())
	}

	FactoryInfoColl = mgo.Database(configs.MongoDB).Collection(FACTORY_INFO_COLL)
	FactoryDataColl = mgo.Database(configs.MongoDB).Collection(FACTORY_DATA_COLL)
	FactoryDataCountColl = mgo.Database(configs.MongoDB).Collection(FACTORY_DATA_COUNT_COLL)
}
