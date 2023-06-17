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
	User  User   `json:"user"`
	Tasks []Task `json:"tasks"`
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

type Task struct {
	ID      uint         `json:"id"`
	UserID  uint         `json:"user_id"`
	Title   string       `json:"title"`
	Content string       `json:"content"`
	Status  model.Status `json:"status"`
}

type TaskDataResponse struct {
	Code     int    `json:"code"`
	Task     Task   `json:"task"`
	TaskList []Task `json:"task_list"`
	Message  string `json:"message"`
	Error    string `json:"error"`
}

func BuildTask(task model.Task) Task {
	return Task{
		ID:      task.ID,
		UserID:  task.UserID,
		Title:   task.Title,
		Content: task.Content,
		Status:  task.Status,
	}
}

func BuildTasks(tasks []model.Task) []Task {
	var tasksData []Task
	for _, task := range tasks {
		tasksData = append(tasksData, BuildTask(task))
	}
	return tasksData
}
