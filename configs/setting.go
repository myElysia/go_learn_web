package configs

const (
	BaseHost = "127.0.0.1"

	MysqlHost = BaseHost
	MysqlPort = "3306"
	MysqlUser = "elysia"
	MysqlPass = "like.phoenix"
	MysqlDB   = "learndb"

	RedisHost = BaseHost
	RedisPort = "6379"
	RedisPass = "redis.elysia"

	MongoHost = BaseHost
	MongoPort = "27017"
	MongoDB   = "learndb"

	LOG_FILE_PATH = "./logs"
	LOG_FILE_NAME = "app.log"
)
