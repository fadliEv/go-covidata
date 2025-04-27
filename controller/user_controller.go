package controller

import (
    "be-covidata/entity"
    "be-covidata/service"
    "github.com/gin-gonic/gin"
    "net/http"
)

func CreateUserController(c *gin.Context) {
    var user entity.User
    if err := c.ShouldBindJSON(&user); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    if err := service.CreateUser(user); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusCreated, gin.H{"message": "User created successfully"})
}

func GetUsersController(c *gin.Context) {
    users, err := service.GetUsers()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, users)
}
