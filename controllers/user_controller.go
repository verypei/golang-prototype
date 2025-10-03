package controllers

import (
	"encoding/json"
	"golang-prototype/dto"
	"golang-prototype/services"
	"golang-prototype/utils"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
	log.Println("Fetching all users with meta ----------------")

	// pagination params from query (default page=1, limit=10)
	page, _ := strconv.Atoi(r.URL.Query().Get("page"))
	if page < 1 {
		page = 1
	}
	limit, _ := strconv.Atoi(r.URL.Query().Get("limit"))
	if limit < 1 {
		limit = 10
	}
	offset := (page - 1) * limit

	// fetch users & total count
	users, total, err := services.GetUsersWithPagination(limit, offset)
	if err != nil {
		utils.Error(w, http.StatusInternalServerError, "Failed to fetch users")
		return
	}

	// build meta info
	meta := &dto.Meta{
		Page:       page,
		Limit:      limit,
		Total:      total,
		TotalPages: (total + limit - 1) / limit, // ceiling division
	}

	utils.SuccessWithMeta(w, "Fetched users successfully", users, meta)
}

func GetUserById(w http.ResponseWriter, r *http.Request) {
	// ✅ Get param from Chi
	idStr := chi.URLParam(r, "id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		utils.Error(w, http.StatusBadRequest, "Invalid user ID")
		return
	}

	user, err := services.GetUserByID(uint(id))
	if err != nil {
		utils.Error(w, http.StatusNotFound, "User not found")
		return
	}

	utils.Success(w, "User fetched successfully", user)
}

func UserRegister(w http.ResponseWriter, r *http.Request) {
	// Decode request body into DTO
	var req dto.UserRegisterRequestDTO
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.Error(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	user, err := services.UserRegister(req)
	if err != nil {
		utils.Error(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.Success(w, "User registered successfully", user)
}
