package service

import (
	"log"
	"time"

	"GoToDo/internal/repository/model"
)

type User struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	CreateAt string `json:"create_at"`
}

type ProfileData struct {
	User User `json:"user"`
}

type UserDataResponse struct {
	Code    int    `json:"code"`
	User    User   `json:"user"`
	Message string `json:"message"`
	Error   string `json:"error"`
	Token   string `json:"token"`
}

type ProfileDataResponse struct {
	Code    int         `json:"code"`
	Profile ProfileData `json:"profile"`
	Message string      `json:"message"`
	Error   string      `json:"error"`
}

func BuildUser(user model.User) User {
	loc, err := time.LoadLocation("America/Los_Angeles")
	if err != nil {
		log.Println("failed to load location:", err)
		return User{}
	} else {
		log.Println("load location success")
	}
	return User{
		ID:       user.ID,
		Username: user.Username,
		CreateAt: user.CreatedAt.In(loc).Format("2006-01-02 15:04:05"),
	}
}
