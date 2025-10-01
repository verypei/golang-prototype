package services

import (
	"log"
)

// func GetUsers() ([]models.User, error) {
// 	log.Println("Service: Getting all users-----------")
// 	return repositories.GetAllUsers()
// }

func GetUsers() (string, error) {
	log.Println("Service: Getting all users-----------")
	return "ok", nil
}

func UserRegister() (string, error) {
	log.Println("Service: Register user-----------")
	return "ok from register---------", nil
}
