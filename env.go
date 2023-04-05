package ginbe

import (
	"fmt"

	"github.com/cntech-io/gin-be/utils"
	"github.com/joho/godotenv"
)

type ServerEnv struct {
	DebugMode        bool
	Port             string
	TrustedProxies   []string
	EnablePrometheus bool
}

func NewServerEnv() *ServerEnv {
	if err := godotenv.Load(); err != nil {
		fmt.Println(".env file not found")
	}
	return &ServerEnv{
		DebugMode:        utils.GetBooleanEnv("DEBUG_MODE", false),
		Port:             utils.GetStringEnv("PORT", false),
		TrustedProxies:   utils.GetStringArrayEnv("TRUSTED_PROXIES", ",", false),
		EnablePrometheus: utils.GetBooleanEnv("PROMETHEUS_FLAG", false),
	}
}
