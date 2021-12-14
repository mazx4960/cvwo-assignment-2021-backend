package api

import (
	"errors"
	"main/database"
	"main/middleware"
	"main/models"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type REGISTER struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func GetUser(c *gin.Context) {
	userId, exist := c.Get("user_id")
	if !exist {
		c.JSON(401, gin.H{"error": "Unable to parse jwt token"})
		return
	}

	var user models.User
	err := database.DB.Where("user_id = ?", userId).First(&user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.JSON(404, gin.H{"error": "User not found"})
		return
	}

	c.JSON(200, gin.H{
		"user": user,
	})
}

func Register(c *gin.Context) {
	var register REGISTER
	
	err := c.ShouldBindJSON(&register)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	var user models.User
	err = database.DB.Where("username = ?", register.Username).First(&user).Error
	if err == nil {
		c.JSON(400, gin.H{
			"error": "Username already exists",
		})
		return
	}

	err = database.DB.Where("email = ?", register.Email).First(&user).Error
	if err == nil {
		c.JSON(400, gin.H{
			"error": "Email already exists",
		})
		return
	}

	password, _ := bcrypt.GenerateFromPassword([]byte(register.Password), 14)

	user = models.User{
		Username: register.Username,
		Password: string(password),
		Email:    register.Email,
		FirstName:     register.FirstName,
		LastName: register.LastName,
	}

	database.DB.Create(&user)

	var token string
	token, err = middleware.CreateToken(user.UserID)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Error creating access token",
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "Registration successful",
		"access_token": token,
		"user": user,
	})
}

type LOGIN struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Login(c *gin.Context) {
	var login LOGIN

	err := c.ShouldBindJSON(&login)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	
	var user models.User
	err = database.DB.Where("username = ?", login.Username).First(&user).Error
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Username does not exist",
		})
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(login.Password))
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Incorrect password",
		})
		return
	}

	var token string
	token, err = middleware.CreateToken(user.UserID)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Error creating access token",
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "Login successful",
		"access_token": token,
		"user": user,
	})
}

func Logout(c *gin.Context) {
	// TODO: revoke access token
	c.JSON(200, gin.H{
		"message": "logout",
	})
}
