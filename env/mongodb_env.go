package env

import (
	"github.com/cntech-io/gin-be/constant"
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
		Username:         utils.GetStringEnv(constant.MONGODB_USERNAME, false),
		Password:         utils.GetStringEnv(constant.MONGODB_PASSWORD, false),
		Database:         utils.GetStringEnv(constant.MONGODB_DATABASE, false),
		ConnectionString: utils.GetStringEnv(constant.MONGODB_CONNECTION_STRING, false),
	}
}
