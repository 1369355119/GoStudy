package hello

import (
	"github.com/gin-gonic/gin"
	"go_code/Study/Collection/gin/middleware"
	"net/http"
	"time"
)

func Routers(r *gin.Engine) {
	r.GET("/", middleware.RateLimitMiddleware(1*time.Second, 1), func(c *gin.Context) {
		c.String(http.StatusOK, "Welcome Gin Server")
	})
}
