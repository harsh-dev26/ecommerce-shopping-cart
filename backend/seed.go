package main

import (
    "ecommerce-backend/config"
    "ecommerce-backend/models"
    "log"
)

func main() {
    config.ConnectDatabase()

    // Create sample items
    items := []models.Item{
        {Name: "Laptop", Status: "available"},
        {Name: "Mouse", Status: "available"},
        {Name: "Keyboard", Status: "available"},
        {Name: "Monitor", Status: "available"},
        {Name: "Headphones", Status: "available"},
    }

    for _, item := range items {
        config.DB.Create(&item)
    }

    log.Println("Sample data created successfully!")
}