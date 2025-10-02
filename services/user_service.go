package services

import (
	"errors"
	"golang-prototype/database"
	"golang-prototype/dto"
	"golang-prototype/models"

	"golang.org/x/crypto/bcrypt"
)

// func GetUsers() (string, error) {
// 	log.Println("Service: Getting all users-----------")
// 	return "ok", nil
// }

func UserRegister(req dto.UserRegisterRequestDTO) (*models.User, error) {
	// Check if email already exists
	var existing models.User
	if err := database.DB.Where("email = ?", req.Email).First(&existing).Error; err == nil {
		return nil, errors.New("email already registered")
	}

	// Hash password
	hashed, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	// Create user
	user := models.User{
		Username: req.Username,
		Email:    req.Email,
		Password: string(hashed),
	}

	if err := database.DB.Create(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}
