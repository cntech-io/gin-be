package ginbe

import (
	"fmt"

	r "github.com/cntech-io/gin-be/response"
	"github.com/gin-gonic/gin"
)

type server struct {
	config *ServerConfig
	router *gin.Engine
}

func NewServer(conf ServerConfig) *server {
	server := &server{}
	if conf.debugMode {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	server.router = gin.New()
	server.router.SetTrustedProxies(conf.trustedProxies)

	return server
}

func (s *server) AddController(c *controller) *server {
	if c.path == "" {
		panic("controller path missing")
	}
	if c.version == "" {
		panic("controller version missing")
	}
	group := s.router.Group(fmt.Sprintf("%v%v", c.version, c.path))
	for _, api := range c.apis {
		group.Handle(string(api.method), api.path, append(api.middlewares, api.handler)...)
	}
	return s
}

func (s *server) AttachMiddleWare(middleware gin.HandlerFunc) *server {
	s.router.Use((middleware))
	return s
}

func (s *server) AttachHealth() *server {
	group := s.router.Group("/health")
	group.GET("/", func(ctx *gin.Context) {
		ctx.JSON(r.OKResponse("Healthy"))
	})
	return s
}

func (s *server) Run() {
	if s.config.port == "" {
		s.router.Run(":8080")
	} else {
		s.router.Run(s.config.port)
	}
}
