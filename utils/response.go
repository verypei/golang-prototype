package utils

import (
	"encoding/json"
	"net/http"

	"golang-prototype/dto"
)

func JSONResponse(w http.ResponseWriter, statusCode int, res dto.Response) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(res)
}

// Success for simple response
func Success(w http.ResponseWriter, message string, data interface{}) {
	JSONResponse(w, http.StatusOK, dto.Response{
		Status:  "success",
		Message: message,
		Data:    data,
	})
}

// SuccessWithMeta for responses with pagination
func SuccessWithMeta(w http.ResponseWriter, message string, data interface{}, meta *dto.Meta) {
	JSONResponse(w, http.StatusOK, dto.Response{
		Status:  "success",
		Message: message,
		Data:    data,
		Meta:    meta,
	})
}

// Error for general error
func Error(w http.ResponseWriter, statusCode int, message string) {
	JSONResponse(w, statusCode, dto.Response{
		Status:  "error",
		Message: message,
	})
}

// ValidationError for field-level validation errors
func ValidationError(w http.ResponseWriter, errors map[string]string) {
	JSONResponse(w, http.StatusBadRequest, dto.Response{
		Status:  "error",
		Message: "Validation failed",
		Errors:  errors,
	})
}
