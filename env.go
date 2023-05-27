package ginbe

import (
	"fmt"

	"github.com/cntech-io/gin-be/constant"
	"github.com/cntech-io/gin-be/utils"
	"github.com/joho/godotenv"
)

type ServerEnv struct {
	DebugModeFlag  bool
	AppPort        string
	TrustedProxies []string
	PrometheusFlag bool
}

func NewServerEnv() *ServerEnv {
	if err := godotenv.Load(); err != nil {
		fmt.Println(".env file not found")
	}
	return &ServerEnv{
		DebugModeFlag:  utils.GetBooleanEnv(constant.DEBUG_MODE_FLAG, false),
		AppPort:        utils.GetStringEnv(constant.APP_PORT, false),
		TrustedProxies: utils.GetStringArrayEnv(constant.TRUSTED_PROXIES, ",", false),
		PrometheusFlag: utils.GetBooleanEnv(constant.PROMETHEUS_FLAG, false),
	}
}
