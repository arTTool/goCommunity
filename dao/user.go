package dao

import (
	"errors"
	"goCommunity/model"
	"goCommunity/util"
)

func UserRegister(email, password string) error {
	pwd := util.GetMd5(password)
	user := model.User{
		Email:    email,
		Password: pwd,
	}
	err := DB.Create(&user).Error
	if err != nil {
		return err
	}
	return nil
}

func GetUser(email, password string) (*model.User, error) {
	var user model.User
	err := DB.Where("email = ? and password=?", email, password).First(&user).Error
	if err != nil {
		return nil, errors.New("用户未找到")
	}
	return &user, nil
}
