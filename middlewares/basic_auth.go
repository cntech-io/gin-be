package middlewares

import (
	r "github.com/cntech-io/gin-be/response"
	"github.com/gin-gonic/gin"
)

func BasicAuthMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		username, password, ok := context.Request.BasicAuth()
		if !ok || username == "" || password == "" {
			context.JSON(r.BadRequestResponse(
				"Invalid request",
				"Provided authentication informations are invalid or have invalid format",
			))
			context.Abort()
		}

		context.Next()
	}
}
