package service

import (
	"GoToDo/internal/repository/dao"
	"GoToDo/internal/repository/model"
	"GoToDo/util"
)

func Register(userRegister model.User) UserDataResponse {
	if dao.CheckUserExist(userRegister.Username) {
		return UserDataResponse{
			Code:    -1,
			Message: "user already exists",
		}
	}
	hashedPassword, err := util.HashPassword(userRegister.Password)
	if err != nil {
		return UserDataResponse{
			Code:    -1,
			Message: "failed to hash password",
			Error:   err.Error(),
		}
	}
	user := model.User{
		Username: userRegister.Username,
		Password: hashedPassword,
	}
	err = dao.InsertOneUser(&user)
	if err != nil {
		return UserDataResponse{
			Code:    -1,
			Message: "failed to create user",
			Error:   err.Error(),
		}
	}
	token, err := util.GenerateToken(user.ID, user.Username)
	if err != nil {
		return UserDataResponse{
			Code:    -1,
			Message: "failed to generate token",
			Error:   err.Error(),
		}
	}

	return UserDataResponse{
		Code:    200,
		Message: "create user success",
		User:    BuildUser(user),
		Token:   token,
	}
}

func Login(userLogin model.User) UserDataResponse {
	user, err := dao.FindOneUserByUsername(userLogin.Username)
	if err != nil {
		return UserDataResponse{
			Code:    -1,
			Message: "user not found",
			Error:   err.Error(),
		}
	}
	err = util.CheckPassword(user.Password, userLogin.Password)
	if err != nil {
		return UserDataResponse{
			Code:    -1,
			Message: "wrong password",
			Error:   err.Error(),
		}
	}

	token, err := util.GenerateToken(user.ID, user.Username)
	if err != nil {
		return UserDataResponse{
			Code:    -1,
			Message: "failed to generate token",
			Error:   err.Error(),
		}
	}
	return UserDataResponse{
		Code:    200,
		Message: "login success",
		User:    BuildUser(user),
		Token:   token,
	}
}
