package controllers

import (
	"encoding/json"
	"golang-prototype/dto"
	"golang-prototype/services"
	"log"
	"net/http"
)

// func GetUsers(w http.ResponseWriter, r *http.Request) {
// 	log.Println("Fetching all users----------------")
// 	users, err := services.GetUsers()
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}

// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(users)
// }

func UserRegister(w http.ResponseWriter, r *http.Request) {
	// Decode request body into DTO
	var req dto.UserRegisterRequestDTO
	log.Println("Decoding request body----:", json.NewDecoder(r.Body))
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Println("❌ Failed to decode request:", err)
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Call service layer to create user
	user, err := services.UserRegister(req)
	if err != nil {
		log.Println("❌ Service error:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Map user model to response DTO
	resp := dto.UserResponseDTO{
		ID:        user.ID,
		Username:  user.Username,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}

	// Return JSON response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
