package controllers

import (
    "ecommerce-backend/config"
    "ecommerce-backend/models"
    "net/http"

    "github.com/gin-gonic/gin"
)

type CreateItemInput struct {
    Name   string `json:"name" binding:"required"`
    Status string `json:"status"`
}

func CreateItem(c *gin.Context) {
    var input CreateItemInput
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    item := models.Item{
        Name:   input.Name,
        Status: input.Status,
    }

    config.DB.Create(&item)
    c.JSON(http.StatusCreated, gin.H{"data": item})
}

func GetItems(c *gin.Context) {
    var items []models.Item
    config.DB.Find(&items)
    c.JSON(http.StatusOK, gin.H{"data": items})
}