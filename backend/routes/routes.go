package routes

import (
    "ecommerce-backend/controllers"
    "ecommerce-backend/middleware"

    "github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
    // User routes
    r.POST("/users", controllers.CreateUser)
    r.POST("/users/login", controllers.Login)
    r.GET("/users", controllers.GetUsers)

    // Item routes
    r.POST("/items", controllers.CreateItem)
    r.GET("/items", controllers.GetItems)

    // Protected routes
    authorized := r.Group("/")
    authorized.Use(middleware.AuthMiddleware())
    {
        // Cart routes
        authorized.POST("/carts", controllers.CreateOrUpdateCart)
        authorized.GET("/carts", controllers.GetCarts)

        // Order routes
        authorized.POST("/orders", controllers.CreateOrder)
        authorized.GET("/orders", controllers.GetOrders)
    }
}