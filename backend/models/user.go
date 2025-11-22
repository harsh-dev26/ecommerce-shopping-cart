package models

import (
    "time"
    // "gorm.io/gorm"
)

type User struct {
    ID        uint      `gorm:"primaryKey" json:"id"`
    Username  string    `gorm:"unique;not null" json:"username"`
    Password  string    `gorm:"not null" json:"-"`
    Token     string    `json:"token,omitempty"`
    CartID    *uint     `json:"cart_id"`
    CreatedAt time.Time `json:"created_at"`
}