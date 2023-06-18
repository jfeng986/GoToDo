package router

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"GoToDo/internal/handler"
	"GoToDo/internal/middleware"
)

func NewRouter() *gin.Engine {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST"},
		AllowHeaders:     []string{"Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "http://localhost:5173"
		},
		MaxAge: 12 * time.Hour,
	}))

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
