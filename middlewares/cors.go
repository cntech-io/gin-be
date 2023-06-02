package middlewares

import (
	"fmt"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func CORSMiddleware(config *cors.Config) gin.HandlerFunc {
	if config == nil {
		fmt.Println("default cors config applied")
		return cors.Default()
	} else {
		fmt.Println("new cors config applied")
		return cors.New(*config)
	}
}
