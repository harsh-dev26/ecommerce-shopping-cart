package config

import (
    "log"

    "ecommerce-backend/models"

    // PURE-Go SQLite driver (NO CGO REQUIRED)
    "github.com/glebarez/sqlite"

    "gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
    // Use the pure-Go SQLite driver
    database, err := gorm.Open(sqlite.Open("ecommerce.db"), &gorm.Config{})
    if err != nil {
        log.Fatal("Failed to connect to database:", err)
    }

    // Auto migrate the schemas
    err = database.AutoMigrate(
        &models.User{},
        &models.Cart{},
        &models.Item{},
        &models.Order{},
    )
    if err != nil {
        log.Fatal("Failed to migrate database:", err)
    }

    DB = database
    log.Println("Database connected successfully!")
}
