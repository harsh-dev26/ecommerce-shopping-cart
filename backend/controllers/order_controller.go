package controllers

import (
    "ecommerce-backend/config"
    "ecommerce-backend/models"
    "net/http"

    "github.com/gin-gonic/gin"
)

type CreateOrderInput struct {
    CartID uint `json:"cart_id" binding:"required"`
}

func CreateOrder(c *gin.Context) {
    userID, _ := c.Get("user_id")
    
    var input CreateOrderInput
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // Verify cart belongs to user
    var cart models.Cart
    if err := config.DB.Where("id = ? AND user_id = ?", input.CartID, userID).First(&cart).Error; err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Cart not found"})
        return
    }

    // Create order
    order := models.Order{
        CartID: input.CartID,
        UserID: userID.(uint),
    }
    config.DB.Create(&order)

    // Update cart status
    config.DB.Model(&cart).Update("status", "ordered")

    c.JSON(http.StatusCreated, gin.H{"data": order})
}

func GetOrders(c *gin.Context) {
    var orders []models.Order
    config.DB.Find(&orders)
    c.JSON(http.StatusOK, gin.H{"data": orders})
}