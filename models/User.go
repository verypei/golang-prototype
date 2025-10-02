package models

import "time"

type User struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Username  string    `json:"username"`
	Email     string    `gorm:"unique" json:"email"`
	Password  string    `json:"-" gorm:"not null"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
