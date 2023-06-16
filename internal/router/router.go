package router

import (
	"GoToDo/internal/handler"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()

	v1 := r.Group("api/v1")
	{
		v1.GET("ping", func(c *gin.Context) {
			c.JSON(200, "success")
		})
		v1.POST("user/register", handler.UserRegisterHandler)
		v1.POST("user/login", handler.UserLoginHandler)
	}
	return r
}
