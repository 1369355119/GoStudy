package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func HomeController(c *gin.Context) {
	username := c.MustGet("username").(string)
	c.JSON(http.StatusOK, gin.H{
		"code": 2000,
		"msg":  "success",
		"data": gin.H{"username": username},
	})
}
