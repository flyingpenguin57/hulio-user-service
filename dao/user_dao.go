package dao

import (
	"errors"
	"hulio-user-service/config"
	"hulio-user-service/dao/model"

	"gorm.io/gorm"
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
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil // 没找到，但不是错误
	}	
	return &user, err
}

// 更新用户, 只更新非0字段
func UpdateUser(user *model.User) error {
	result := config.DB.Model(&model.User{}).Where("id = ?", user.ID).Updates(user)
	return result.Error
}

// 删除用户
func DeleteUser(id uint) error {
	return config.DB.Delete(&model.User{}, id).Error
}
