package middlewares

import (
	"fmt"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func CORSMiddleware(config cors.Config) gin.HandlerFunc {
	fmt.Println("new cors config applied\n", config)
	return cors.New(config)
}
