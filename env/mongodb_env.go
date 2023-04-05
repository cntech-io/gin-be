package env

import (
	"github.com/cntech-io/gin-be/utils"
)

type MongoDBEnv struct {
	Username         string
	Password         string
	Database         string
	ConnectionString string
}

func NewMongoDBEnv() *MongoDBEnv {
	return &MongoDBEnv{
		Username:         utils.GetStringEnv("MONGODB_USERNAME", false),
		Password:         utils.GetStringEnv("MONGODB_PASSWORD", false),
		Database:         utils.GetStringEnv("MONGODB_DATABASE", false),
		ConnectionString: utils.GetStringEnv("MONGODB_CONNECTION_STRING", false),
	}
}
