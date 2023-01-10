package main

import (
	log "github.com/sirupsen/logrus"

	"go_learn_web/logs"
	"go_learn_web/pkg/mongo"
	"go_learn_web/pkg/mysql"
	"go_learn_web/pkg/redis"
)

func main() {
	logs.Init()
	mysql.Init()
	mongo.Init()
	redis.Init()
	log.Info("server has started.")
}
