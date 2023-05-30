package env

import (
	"github.com/cntech-io/gin-be/constant"

	"github.com/cntech-io/gin-be/utils"
)

type PostgresDBEnv struct {
	Host          string
	Port          string
	Username      string
	Password      string
	MigrationFlag bool
}

func NewPostgressDBEnv() *PostgresDBEnv {
	return &PostgresDBEnv{
		Host:          utils.GetStringEnv(constant.POSTGRESDB_HOST, false),
		Port:          utils.GetStringEnv(constant.POSTGRESDB_PORT, false),
		Username:      utils.GetStringEnv(constant.POSTGRESDB_USERNAME, false),
		Password:      utils.GetStringEnv(constant.POSTGRESDB_PASSWORD, false),
		MigrationFlag: utils.GetBooleanEnv(constant.POSTGRESDB_MIGRATION_FLAG, false),
	}
}
