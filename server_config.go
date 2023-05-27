package ginbe

import "github.com/gin-contrib/cors"

type ServerConfig struct {
	AppPort        string
	DebugModeFlag  bool
	TrustedProxies []string
	PrometheusFlag bool
	CorsConfig     cors.Config
}
