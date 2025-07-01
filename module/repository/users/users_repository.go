package users

import (
	"backend-go/config"
	"backend-go/module/model/migrate"
)

func CreateUsers(req migrate.User) (*migrate.User, error) {
	if err := config.DbConn.Create(&req).Error; err != nil {
		return nil, err
	}
	return &req, nil
}

func GetUserByEmail(email string) (*migrate.User, error) {
	var user migrate.User
	if err := config.DbConn.Table("users").Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func GetUserById(id string) (*migrate.User, error) {
	var user migrate.User
	if err := config.DbConn.Table("users").Where("id = ?", id).First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}
