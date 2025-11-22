package middleware

import (
    "ecommerce-backend/config"
    "ecommerce-backend/models"
    "net/http"
    "strings"

    "github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        token := c.GetHeader("Authorization")
        
        if token == "" {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization token required"})
            c.Abort()
            return
        }

        // Remove "Bearer " prefix if present
        token = strings.TrimPrefix(token, "Bearer ")

        var user models.User
        if err := config.DB.Where("token = ?", token).First(&user).Error; err != nil {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
            c.Abort()
            return
        }

        c.Set("user_id", user.ID)
        c.Next()
    }
}