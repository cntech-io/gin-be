package ginbe

type ServerConfig struct {
	Port             string
	DebugMode        bool
	TrustedProxies   []string
	EnablePrometheus bool
}
