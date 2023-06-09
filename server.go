package ginbe

import (
	"fmt"

	r "github.com/cntech-io/gin-be/response"
	"github.com/gin-gonic/gin"
	ginprometheus "github.com/zsais/go-gin-prometheus"
)

type Server struct {
	config *ServerConfig
	router *gin.Engine
}

func NewServer(conf ServerConfig) *Server {
	server := &Server{}
	server.config = &conf
	if conf.DebugModeFlag {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	server.router = gin.New()
	server.router.SetTrustedProxies(conf.TrustedProxies)

	if conf.PrometheusFlag {
		prom := ginprometheus.NewPrometheus("gin")
		prom.Use(server.router)
	}

	return server
}

func (s *Server) AddController(c *Controller) *Server {
	if c.path == "" {
		panic("controller path missing")
	}
	if c.version == "" {
		panic("controller version missing")
	}
	group := s.router.Group(fmt.Sprintf("%v%v", c.version, c.path))
	for _, resource := range c.resources {
		group.Handle(string(resource.method), resource.path, append(resource.middlewares, resource.handler)...)
	}
	return s
}

func (s *Server) AttachMiddleWare(middleware gin.HandlerFunc) *Server {
	s.router.Use(middleware)
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
	if s.config.AppPort == "" {
		s.router.Run(":8080")
	} else {
		s.router.Run(s.config.AppPort)
	}
}

func (s *Server) Router() *gin.Engine {
	return s.router
}
