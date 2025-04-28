package main

import (
    "be-covidata/config"
    "be-covidata/controller"
    "be-covidata/service"
    "github.com/gin-gonic/gin"
)

func main() {
    config.LoadEnv()

    // Setup database
    service.SetupDB()

    r := gin.Default()

    // Route
    r.POST("/users", controller.CreateUserController)
    r.GET("/users", controller.GetUsersController)
    r.PUT("/users", controller.UpdateUserController)
    r.DELETE("/users/:id", controller.DeleteUserController)

    port := config.GetEnv("PORT")
    r.Run(":" + port) // Start server
}
