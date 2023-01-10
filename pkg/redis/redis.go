package redis

import (
	log "github.com/sirupsen/logrus"

	"context"
	"go_learn_web/configs"
	"time"

	"github.com/go-redis/redis/v8"
)

var rdb *redis.Client

func Init() {
	rdb = redis.NewClient(&redis.Options{
		Addr:     configs.RedisHost + ":" + configs.RedisPort,
		Password: configs.RedisPass,
		DB:       0,
		PoolSize: 100,
	})
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		log.Panic(err.Error())
	}
}
