package controllers

import (
    "backend/models"
    "backend/utils"
    "net/http"

    "github.com/gin-gonic/gin"
)

type RegisterRequest struct {
    Name     string `json:"name" binding:"required"`
    Email    string `json:"email" binding:"required,email"`
    Password string `json:"password" binding:"required,min=6"`
}

type LoginRequest struct {
    Email    string `json:"email" binding:"required,email"`
    Password string `json:"password" binding:"required"`
}

type UserResponse struct {
    ID        int    `json:"id"`
    Name      string `json:"name"`
    Email     string `json:"email"`
    CreatedAt string `json:"created_at"`
}

// Register controller (sudah ada)
func Register(c *gin.Context) {
    var req RegisterRequest
    
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "success": false,
            "message": "Invalid request data",
            "error":   err.Error(),
        })
        return
    }
    
    existingUser, _ := models.GetUserByEmail(req.Email)
    if existingUser != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "success": false,
            "message": "Email already registered",
        })
        return
    }
    
    hashedPassword, err := utils.HashPassword(req.Password)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
            "success": false,
            "message": "Failed to process password",
        })
        return
    }
    
    err = models.CreateUser(req.Name, req.Email, hashedPassword)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
            "success": false,
            "message": "Failed to create user",
            "error":   err.Error(),
        })
        return
    }
    
    user, _ := models.GetUserByEmail(req.Email)
    
    c.JSON(http.StatusCreated, gin.H{
        "success": true,
        "message": "User registered successfully",
        "data": UserResponse{
            ID:        user.ID,
            Name:      user.Name,
            Email:     user.Email,
            CreatedAt: user.CreatedAt.Format("2006-01-02 15:04:05"),
        },
    })
}

// Login controller (BARU)
func Login(c *gin.Context) {
    var req LoginRequest
    
    // Bind JSON request
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "success": false,
            "message": "Invalid request data",
            "error":   err.Error(),
        })
        return
    }
    
    // Cari user by email
    user, err := models.GetUserByEmail(req.Email)
    if err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{
            "success": false,
            "message": "Invalid email or password",
        })
        return
    }
    
    // Verifikasi password
    if !utils.CheckPasswordHash(req.Password, user.Password) {
        c.JSON(http.StatusUnauthorized, gin.H{
            "success": false,
            "message": "Invalid email or password",
        })
        return
    }
    
    // Generate JWT token
    token, err := utils.GenerateToken(user.ID, user.Email)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
            "success": false,
            "message": "Failed to generate token",
        })
        return
    }
    
    c.JSON(http.StatusOK, gin.H{
        "success": true,
        "message": "Login successful",
        "data": gin.H{
            "token": token,
            "user": UserResponse{
                ID:        user.ID,
                Name:      user.Name,
                Email:     user.Email,
                CreatedAt: user.CreatedAt.Format("2006-01-02 15:04:05"),
            },
        },
    })
}