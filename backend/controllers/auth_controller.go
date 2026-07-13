package controllers

import(
	"backend/models"
	"backend/utils"

	"net/http"
	"github.com/gin-gonic/gin"
)

type RegisterRequest struct{
	Name string `json:"name" binding:"required"`
	Email string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

type UserResponse struct{
	ID int `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
	CreatedAt string `json:"created_at"`
}

func Register(c *gin.Context){
	var req RegisterRequest

	if err := c.ShouldBindJSON(&req); err != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"success" : false,
			"message" : "Invalid request data",
			"error" : err.Error(),
		})
		return 
	}

	existingUser, _ := models.GetUserByEmail(req.Email)
	if existingUser != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"success" : false,
			"message" : "User already registered",
		})
		return
	}

	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil{
		c.JSON(http.StatusInternalServerError, gin.H{
			"success" : false,
			"message" : "Failed to process password",
		})
		return
	}

	err = models.CreateUser(req.Name, req.Email, hashedPassword)
	if err != nil{
		c.JSON(http.StatusInternalServerError, gin.H{
			"success" : false,
			"message" : "Failed to create user",
			"error" : err.Error(),
		})
		return
	}

	user, _ := models.GetUserByEmail(req.Email)

	c.JSON(http.StatusCreated, gin.H{
		"success" : true,
		"message" : "User registered",
		"data" : UserResponse{
			ID:			user.ID,
			Name:		user.Name,
			Email:		user.Email,
			CreatedAt: 	user.CreatedAt.Format("2006-04-08 16:03:04"),
		},

	})
}

