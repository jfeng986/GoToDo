package service

import (
	"log"

	"GoToDo/internal/repository/dao"
	"GoToDo/internal/repository/model"
)

func CreateTask(userId uint, createTask model.Task) TaskDataResponse {
	user, err := dao.FindOneUserById(userId)
	if err != nil {
		return TaskDataResponse{
			Code:    -1,
			Message: "user not found",
			Error:   err.Error(),
		}
	}
	task := model.Task{
		UserID:  user.ID,
		Title:   createTask.Title,
		Content: createTask.Content,
		Status:  createTask.Status,
		User:    user,
	}
	err = dao.InsertOneTask(&task)
	if err != nil {
		return TaskDataResponse{
			Code:    -1,
			Message: "failed to create task",
			Error:   err.Error(),
		}
	}
	return TaskDataResponse{
		Code:    200,
		Message: "create task success",
		Task:    BuildTask(task),
	}
}

func GetOneTask(userId uint, taskId string) TaskDataResponse {
	task, err := dao.FindOneTaskById(taskId)
	if err != nil {
		return TaskDataResponse{
			Code:    -1,
			Message: "task not found",
			Error:   err.Error(),
		}
	}
	if task.UserID != userId {
		return TaskDataResponse{
			Code:    -1,
			Message: "forbidden",
		}
	}
	return TaskDataResponse{
		Code:    200,
		Message: "get task success",
		Task:    BuildTask(task),
	}
}

func GetAllTasks(userId uint) TaskDataResponse {
	tasks, err := dao.FindAllTasksByUserId(userId)
	if err != nil {
		return TaskDataResponse{
			Code:    -1,
			Message: "tasks not found",
			Error:   err.Error(),
		}
	}
	return TaskDataResponse{
		Code:     200,
		Message:  "get tasks success",
		TaskList: BuildTasks(tasks),
	}
}

func UpdateTask(userId uint, taskId string, updateTask model.Task) TaskDataResponse {
	task, err := dao.FindOneTaskById(taskId)
	if err != nil {
		return TaskDataResponse{
			Code:    -1,
			Message: "task not found",
			Error:   err.Error(),
		}
	}
	if task.UserID != userId {
		return TaskDataResponse{
			Code:    -1,
			Message: "forbidden",
		}
	}
	log.Println("update task:", updateTask)
	task.Title = updateTask.Title
	task.Content = updateTask.Content
	task.Status = updateTask.Status
	err = dao.UpdateOneTask(&task)
	if err != nil {
		return TaskDataResponse{
			Code:    -1,
			Message: "failed to update task",
			Error:   err.Error(),
		}
	}
	return TaskDataResponse{
		Code:    200,
		Message: "update task success",
		Task:    BuildTask(task),
	}
}

func DeleteTask(userId uint, taskId string) TaskDataResponse {
	task, err := dao.FindOneTaskById(taskId)
	if err != nil {
		return TaskDataResponse{
			Code:    -1,
			Message: "task not found",
			Error:   err.Error(),
		}
	}
	if task.UserID != userId {
		return TaskDataResponse{
			Code:    -1,
			Message: "forbidden",
		}
	}
	err = dao.DeleteOneTask(&task)
	if err != nil {
		return TaskDataResponse{
			Code:    -1,
			Message: "failed to delete task",
			Error:   err.Error(),
		}
	}
	return TaskDataResponse{
		Code:    200,
		Message: "delete task success",
	}
}
