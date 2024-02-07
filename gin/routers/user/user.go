package user

import (
	"github.com/gin-gonic/gin"
	"go_code/Study/Collection/controller"
	"go_code/Study/Collection/gin/middleware"
)

func Routers(r *gin.Engine) {
	r.POST("/login", controller.AuthController)

	userGroup := r.Group("/user", middleware.StatCost(), middleware.JWTAuthMiddleware())
	{
		userGroup.GET("/home", controller.HomeController)
	}
}
