package models

import (
    "time"
)

type Cart struct {
    ID        uint      `gorm:"primaryKey" json:"id"`
    UserID    uint      `gorm:"not null" json:"user_id"`
    Name      string    `json:"name"`
    Status    string    `json:"status"` // active, ordered
    CreatedAt time.Time `json:"created_at"`
    Items     []Item    `gorm:"many2many:cart_items;" json:"items,omitempty"`
}