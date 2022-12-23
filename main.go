package main

import (
	"go_learn_web/logs"
	"go_learn_web/pkg/mongo"
	"go_learn_web/pkg/mysql"
	"go_learn_web/pkg/redis"
)

func main() {
	logs.InitLog()
	mysql.InitMysqlDB()
	mongo.InitMongo()
	redis.InitRedis()
	logs.Log.Info("server has started.")
}
