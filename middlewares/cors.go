package middlewares

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func CORSMiddleware(config *cors.Config) gin.HandlerFunc {
	if config == nil {
		return cors.Default()
	} else {
		return cors.New(*config)
	}
}
