package env

import (
	"github.com/cntech-io/gin-be/constant"

	"github.com/cntech-io/gin-be/utils"
)

type RedisDBEnv struct {
	Host     string
	Port     string
	Password string
}

func NewRedisDBEnv() *RedisDBEnv {
	return &RedisDBEnv{
		Host:     utils.GetStringEnv(constant.REDISDB_HOST, false),
		Port:     utils.GetStringEnv(constant.REDISDB_PORT, false),
		Password: utils.GetStringEnv(constant.REDISDB_PASSWORD, false),
	}
}
