package main

import (
    "github.com/gin-gonic/gin"
    "github.com/YadavPrasad/FileShip/backend/handlers"
    "github.com/YadavPrasad/FileShip/backend/middleware"
)

func main() {
    r := gin.Default()

    publicRoutes := r.Group("/public")
    {
        publicRoutes.POST("/login", handlers.Login)
        publicRoutes.POST("/register", handlers.Register)
    }

    protectedRoutes := r.Group("/protected")
    protectedRoutes.Use(middleware.AuthenticationMiddleware())
    {

    }

    r.Run(":8080")
}