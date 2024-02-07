package controller

import (
	"github.com/gin-gonic/gin"
	"go_code/Study/Collection/jwt"
	"go_code/Study/Collection/model"
	"net/http"
)

func AuthController(c *gin.Context) {
	// 用户发送用户名和密码过来
	var user model.UserInfo
	if err := c.ShouldBind(&user); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 2001,
			"msg":  "无效的参数",
		})
		return
	}

	// 校验用户名和密码是否正确
	if user.Username == "1" && user.Password == "1" {
		// 生成Token
		token, err := jwt.GenToken(user.Username)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code": 2002,
				"msg":  "生成token失败" + err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"code": 2000,
			"msg":  "success",
			"data": gin.H{"token": token},
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 2003,
		"msg":  "鉴权失败",
	})
	return
}
