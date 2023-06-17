package handler

import (
	"log"
	"net/http"

	"GoToDo/internal/repository/model"
	"GoToDo/internal/service"
	"GoToDo/util"

	"github.com/gin-gonic/gin"
)

func CreateTaskHandler(c *gin.Context) {
	var createTask model.Task
	if err := c.ShouldBind(&createTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	} else {
		log.Printf("Task data: %+v\n", createTask)
		claims, exists := c.Get("claims")
		if !exists {
			c.JSON(http.StatusBadRequest, gin.H{"error": "no claims"})
			return
		}
		userClaims, ok := claims.(*util.UserClaims)
		if !ok {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "claims cast error"})
			return
		}
		res := service.CreateTask(userClaims.ID, createTask)
		c.JSON(http.StatusOK, res)
		return
	}
}

func GetTaskHandler(c *gin.Context) {
	claims, exists := c.Get("claims")
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "no claims"})
		return
	}
	userClaims, ok := claims.(*util.UserClaims)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "claims cast error"})
		return
	}
	res := service.GetOneTask(userClaims.ID, c.Param("id"))
	c.JSON(http.StatusOK, res)
}

func GetAllTasksHandler(c *gin.Context) {
	claims, exists := c.Get("claims")
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "no claims"})
		return
	}
	userClaims, ok := claims.(*util.UserClaims)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "claims cast error"})
		return
	}
	res := service.GetAllTasks(userClaims.ID)
	c.JSON(http.StatusOK, res)
}

func UpdateTaskHandler(c *gin.Context) {
	var updateTask model.Task
	log.Println("updateTask: ", updateTask)
	if err := c.ShouldBind(&updateTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	} else {
		claims, exists := c.Get("claims")
		if !exists {
			c.JSON(http.StatusBadRequest, gin.H{"error": "no claims"})
			return
		}
		userClaims, ok := claims.(*util.UserClaims)
		if !ok {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "claims cast error"})
			return
		}
		res := service.UpdateTask(userClaims.ID, c.Param("id"), updateTask)
		c.JSON(http.StatusOK, res)
		return
	}
}

func DeleteTaskHandler(c *gin.Context) {
	claims, exists := c.Get("claims")
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "no claims"})
		return
	}
	userClaims, ok := claims.(*util.UserClaims)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "claims cast error"})
		return
	}
	res := service.DeleteTask(userClaims.ID, c.Param("id"))
	c.JSON(http.StatusOK, res)
}
