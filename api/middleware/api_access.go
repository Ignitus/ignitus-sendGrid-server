package middleware

import (
	"net/http"

	"github.com/Ignitus/ignitus-mailing-microservice/structure"
	"github.com/gin-gonic/gin"
)

// APIAccessTo Middleware only accept requests originated from host.
func APIAccessTo(host string) gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.Host != host {
			c.JSON(http.StatusUnauthorized, structure.ErrorResponse("Unauthorized."))
			c.Abort()
			return
		}
		c.Next()
	}
}
