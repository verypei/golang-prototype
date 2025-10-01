package controllers

import (
	"encoding/json"
	"golang-prototype/services"
	"log"
	"net/http"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
	log.Println("Fetching all users----------------")
	users, err := services.GetUsers()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func UserRegister(w http.ResponseWriter, r *http.Request) {
	log.Println("Fetching all users----------------")
	users, err := services.UserRegister()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}
