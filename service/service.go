package service

import (
	"todolist/dao/mysql"
	"todolist/models"
	"todolist/utils"
)

func Register(username, password string) error {
	u := &models.User{
		Username: username,
		Password: password,
	}
	if err := mysql.DB.Create(u).Error; err != nil {
		return err
	}
	return nil
}

func Login(username, password string) (string, error) {
	u := &models.User{}
	if err := mysql.DB.Where("username = ? AND password = ?", username, password).First(u).Error; err != nil {
		return "", err
	}
	// 生成 JWT token 并返回
	token, err := utils.GenerateToken(u.ID)
	if err != nil {
		return "", err
	}
	return token, nil
}

func AddTodo(userID int64, content string) error {
	t := &models.Todo{
		UserId:  userID,
		Content: content,
	}
	return mysql.DB.Create(t).Error
}

func GetTodoList(userID int64) ([]models.Todo, error) {
	var todos []models.Todo
	if err := mysql.DB.Where("user_id = ?", userID).Find(&todos).Error; err != nil {
		return nil, err
	}
	return todos, nil
}

func UpdateTodo(todoID int64, content string) error {
	return mysql.DB.Model(&models.Todo{}).Where("id = ?", todoID).Update("content", content).Error
}

func DeleteTodo(todoID int64) error {
	return mysql.DB.Where("id = ?", todoID).Delete(&models.Todo{}).Error
}
