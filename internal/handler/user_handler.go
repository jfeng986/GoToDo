package handler

import (
	"log"
	"net/http"

	"GoToDo/internal/repository/model"
	"GoToDo/internal/service"

	"github.com/gin-gonic/gin"
)

func UserRegisterHandler(c *gin.Context) {
	var userRegister model.User
	if err := c.ShouldBind(&userRegister); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	} else {
		res := service.Register(userRegister)
		c.JSON(http.StatusOK, res)
		return
	}
}

func UserLoginHandler(c *gin.Context) {
	var userLogin model.User

	if err := c.ShouldBind(&userLogin); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	} else {
		res := service.Login(userLogin)
		c.JSON(http.StatusOK, res)
		return
	}
}
