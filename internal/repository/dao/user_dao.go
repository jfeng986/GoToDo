package dao

import (
	"GoToDo/internal/repository/database"
	"GoToDo/internal/repository/model"
)

func CheckUserExist(username string) bool {
	var count int64
	database.DB.Model(&model.User{}).Where("username = ?", username).Count(&count)
	return count != 0
}

func FindOneUserByUsername(username string) (model.User, error) {
	var user model.User
	return user, database.DB.Where("username = ?", username).First(&user).Error
}

func FindOneUserById(id uint) (model.User, error) {
	var user model.User
	return user, database.DB.Where("id = ?", id).First(&user).Error
}

func InsertOneUser(user *model.User) error {
	return database.DB.Create(&user).Error
}

/*
func InsertOneUser(user *model.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	time.Sleep(time.Second * 10)
	db := database.DB.WithContext(ctx)
	err := db.Create(&user).Error
	return err
}
*/
