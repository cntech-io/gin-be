package ginbe

import (
	"fmt"

	r "github.com/cntech-io/gin-be/response"
	"github.com/gin-gonic/gin"
)

type Server struct {
	config *ServerConfig
	router *gin.Engine
}

func NewServer(conf ServerConfig) *Server {
	server := &Server{}
	if conf.DebugMode {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	server.router = gin.New()
	server.router.SetTrustedProxies(conf.TrustedProxies)

	return server
}

func (s *Server) AddController(c *controller) *Server {
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

func (s *Server) AttachMiddleWare(middleware gin.HandlerFunc) *Server {
	s.router.Use((middleware))
	return s
}

func (s *Server) AttachHealth() *Server {
	group := s.router.Group("/health")
	group.GET("/", func(ctx *gin.Context) {
		ctx.JSON(r.OKResponse("Healthy"))
	})
	return s
}

func (s *Server) Run() {
	if s.config.Port == "" {
		s.router.Run(":8080")
	} else {
		s.router.Run(s.config.Port)
	}
}
