package middleware

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"time"
)

/**
使用社区的https://github.com/gin-contrib/cors 库，一行代码解决前后端分离架构下的跨域问题。
*/

// Cors 跨域中间件
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		cors.New(cors.Config{
			AllowOrigins:     []string{"https://foo.com"},                         // 允许跨域发来请求的网站
			AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}, // 允许的请求方法
			AllowHeaders:     []string{"Origin", "Authorization", "Content-Type"},
			ExposeHeaders:    []string{"Content-Length"},
			AllowCredentials: true,
			AllowOriginFunc: func(origin string) bool { // 自定义过滤源站的方法
				return origin == "https://github.com"
			},
			MaxAge: 12 * time.Hour,
		})
	}
}
