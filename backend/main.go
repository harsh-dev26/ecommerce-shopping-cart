package main

import (
    "ecommerce-backend/config"
    "ecommerce-backend/routes"
    "os"

    "github.com/gin-contrib/cors"
    "github.com/gin-gonic/gin"
)

func main() {
    r := gin.Default()

    // CORS middleware
    r.Use(cors.New(cors.Config{
        AllowOrigins:     []string{"*"}, // allow frontend later
        AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
        AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
        ExposeHeaders:    []string{"Content-Length"},
        AllowCredentials: true,
    }))

    // Connect to database
    config.ConnectDatabase()

    // Setup routes
    routes.SetupRoutes(r)

    // ---- IMPORTANT FOR RENDER ----
    port := os.Getenv("PORT")
    if port == "" {
        port = "8080" // local fallback
    }

    r.Run(":" + port)
}
