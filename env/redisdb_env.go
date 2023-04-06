package env

import (
	"github.com/cntech-io/gin-be/utils"
)

type RedisDBEnv struct {
	Host     string
	Port     string
	Password string
}

func NewRedisDBEnv() *RedisDBEnv {
	return &RedisDBEnv{
		Host:     utils.GetStringEnv("REDISDB_HOST", false),
		Port:     utils.GetStringEnv("REDISDB_PORT", false),
		Password: utils.GetStringEnv("REDISDB_PASSWORD", false),
	}
}
