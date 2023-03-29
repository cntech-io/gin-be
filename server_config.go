package ginbe

type ServerConfig struct {
	port           string
	debugMode      bool
	trustedProxies []string
}
