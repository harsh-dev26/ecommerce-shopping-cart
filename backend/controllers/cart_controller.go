package controllers

import (
    "ecommerce-backend/config"
    "ecommerce-backend/models"
    "net/http"

    "github.com/gin-gonic/gin"
)

type AddToCartInput struct {
    ItemIDs []uint `json:"item_ids" binding:"required"`
}

func CreateOrUpdateCart(c *gin.Context) {
    userID, _ := c.Get("user_id")
    
    var input AddToCartInput
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // Find or create cart for user
    var cart models.Cart
    result := config.DB.Where("user_id = ? AND status = ?", userID, "active").First(&cart)
    
    if result.Error != nil {
        // Create new cart
        cart = models.Cart{
            UserID: userID.(uint),
            Status: "active",
            Name:   "Shopping Cart",
        }
        config.DB.Create(&cart)
    }

    // Add items to cart
    var items []models.Item
    config.DB.Find(&items, input.ItemIDs)
    
    config.DB.Model(&cart).Association("Items").Append(&items)

    // Update user's cart_id
    config.DB.Model(&models.User{}).Where("id = ?", userID).Update("cart_id", cart.ID)

    config.DB.Preload("Items").First(&cart, cart.ID)
    c.JSON(http.StatusOK, gin.H{"data": cart})
}

func GetCarts(c *gin.Context) {
    var carts []models.Cart
    config.DB.Preload("Items").Find(&carts)
    c.JSON(http.StatusOK, gin.H{"data": carts})
}