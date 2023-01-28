package configs

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

//
//var (
//	MysqlHost = ""
//	MysqlPort = "3306"
//	MysqlUser = "elysia"
//	MysqlPass = "like.phoenix"
//	MysqlDB   = "learndb"
//
//	RedisHost = ""
//	RedisPort = "6379"
//	RedisPass = "redis.elysia"
//
//	MongoHost = ""
//	MongoPort = "27017"
//	MongoDB   = "learndb"
//
//	LogFilePath = "./logs"
//	LogFileName = "app.log"
//
//	GinRunHost = "127.0.0.1"
//	GinRunPort = "8080"
//)

type configData struct {
	Mysql struct {
		Host     string `yaml:"host"`
		Port     string `yaml:"port"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		Database string `yaml:"database"`
	}
	Redis struct {
		Host     string `yaml:"host"`
		Port     string `yaml:"port"`
		Password string `yaml:"password"`
	}
	Mongodb struct {
		Host     string `yaml:"host"`
		Port     string `yaml:"port"`
		Database string `yaml:"database"`
	}
	Log struct {
		Path string `yaml:"path"`
		Name string `yaml:"name"`
	}
	Gin struct {
		Host string `yaml:"host"`
		Post string `yaml:"post"`
	}
}

var ConfigData *configData

func init() {
	config := new(configData)
	yamlFile, _ := os.ReadFile("./configs/develop_config.yaml")
	err := yaml.Unmarshal(yamlFile, config)
	if err != nil {
		fmt.Println("yaml load has error" + err.Error())
		return
	}
	ConfigData = config
}
