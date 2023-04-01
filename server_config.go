package ginbe

import "github.com/gin-contrib/cors"

type ServerConfig struct {
	Port             string
	DebugMode        bool
	TrustedProxies   []string
	EnablePrometheus bool
	CorsConfig       cors.Config
}
