package users

import (
	"backend-go/internal/model/migrate"

	"gorm.io/gorm"
)

func CreateUsers(db *gorm.DB, req migrate.User) (*migrate.User, error) {
	if err := db.Create(&req).Error; err != nil {
		return nil, err
	}
	return &req, nil
}

func GetUserByEmail(db *gorm.DB, email string) (*migrate.User, error) {
	var user migrate.User
	if err := db.Table("users").Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func GetUserById(db *gorm.DB, id string) (*migrate.User, error) {
	var user migrate.User
	if err := db.Table("users").Where("id = ?", id).First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}
