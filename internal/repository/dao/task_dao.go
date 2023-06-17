package dao

import (
	"GoToDo/internal/repository/database"
	"GoToDo/internal/repository/model"
)

func InsertOneTask(task *model.Task) error {
	return database.DB.Create(&task).Error
}

func FindOneTaskById(id string) (model.Task, error) {
	var task model.Task
	return task, database.DB.Where("id = ?", id).First(&task).Error
}

func FindAllTasksByUserId(userId uint) ([]model.Task, error) {
	var tasks []model.Task
	return tasks, database.DB.Where("user_id = ?", userId).Find(&tasks).Error
}

func UpdateOneTask(task *model.Task) error {
	return database.DB.Save(&task).Error
}

func DeleteOneTask(task *model.Task) error {
	return database.DB.Delete(&task).Error
	// return database.DB.Unscoped().Delete(&task).Error
}
