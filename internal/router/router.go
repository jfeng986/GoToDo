package router

import (
	"github.com/gin-gonic/gin"

	"GoToDo/internal/handler"
	"GoToDo/internal/middleware"
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

		authed := v1.Group("")
		authed.Use(middleware.JWT)
		{
			authed.GET("user/profile", handler.GetProfileHandler)
			authed.POST("create_task", handler.CreateTaskHandler)
			authed.GET("get_one_task/:id", handler.GetTaskHandler)
			authed.GET("get_all_tasks", handler.GetAllTasksHandler)
			authed.POST("update_task/:id", handler.UpdateTaskHandler)
			authed.POST("delete_task/:id", handler.DeleteTaskHandler)
		}
	}
	return r
}
