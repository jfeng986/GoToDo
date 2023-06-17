package middleware

import (
	"time"

	"GoToDo/util"

	"github.com/gin-gonic/gin"
)

func JWT(c *gin.Context) {
	var code int
	token := c.GetHeader("Authorization")
	if token == "" {
		code = -1
	} else {
		claims, err := util.ParseToken(token)
		if err != nil {
			code = -1
		} else if claims.ExpiresAt.Unix() < time.Now().Unix() {
			code = -1
		} else {
			c.Set("claims", claims)
			code = 200
		}
	}
	if code != 200 {
		c.JSON(200, gin.H{
			"code":    code,
			"message": "token error",
		})
		c.Abort()
		return
	}
	c.Next()
}
