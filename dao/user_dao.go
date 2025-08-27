package dao

import (
	"hulio-user-service/config"
	"hulio-user-service/dao/model"
)

// 新建用户
func CreateUser(user *model.User) error {
	return config.DB.Create(user).Error
}

// 根据 ID 查找用户
func GetUserByID(id uint) (*model.User, error) {
	var user model.User
	err := config.DB.First(&user, id).Error
	return &user, err
}

// 根据 username 查找用户
func GetUserByUsername(username string) (*model.User, error) {
	var user model.User
	err := config.DB.Where("username = ?", username).First(&user).Error
	return &user, err
}

// 更新用户
func UpdateUser(user *model.User) error {
	return config.DB.Save(user).Error
}

// 删除用户
func DeleteUser(id uint) error {
	return config.DB.Delete(&model.User{}, id).Error
}
