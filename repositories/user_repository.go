package repositories

import (
	"golang-prototype/database"
	"golang-prototype/models"
)

func GetAllUsers() ([]models.User, error) {
	var users []models.User
	result := database.DB.Find(&users)
	return users, result.Error
}
