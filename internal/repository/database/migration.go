package database

import "GoToDo/internal/repository/model"

func migration() {
	err := DB.Set("gorm:table_options", "charset=utf8mb4").
		AutoMigrate(&model.User{}, &model.Task{})
	if err != nil {
		return
	}
}

func Test() {
	return
}
