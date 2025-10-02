package dto

type Response struct {
	Status  string      `json:"status"`           // "success" or "error"
	Message string      `json:"message"`          // human-readable message
	Data    interface{} `json:"data,omitempty"`   // response payload (optional)
	Meta    *Meta       `json:"meta,omitempty"`   // extra metadata (pagination, etc.)
	Errors  interface{} `json:"errors,omitempty"` // validation or detailed errors
}

type Meta struct {
	Page       int `json:"page,omitempty"`
	Limit      int `json:"limit,omitempty"`
	Total      int `json:"total,omitempty"`
	TotalPages int `json:"total_pages,omitempty"`
}
