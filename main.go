package main

import (
	"go_learn_web/logs"
	"go_learn_web/pkg/gin"
	"go_learn_web/pkg/mongo"
	"go_learn_web/pkg/mysql"
	"go_learn_web/pkg/redis"
)

func main() {
	// 日志初始化
	logs.Init()
	// mysql连接，产生sqlconn, gormdb
	mysql.Init()
	// 连接mongodb
	mongo.Init()
	// 连接redis
	redis.Init()
	gin.Init()
}
