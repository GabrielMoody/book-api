package handler

import (
	"book-api/app/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// var conn = models.GetConnection()

func Login(db *gorm.DB, c *gin.Context) {
	var user models.User
	username := c.PostForm("username")
	password := c.PostForm("password")
	result := db.Where("username = ?", username).Find(&user)

	if result.RowsAffected < 1 {
		c.JSON(http.StatusBadRequest, gin.H{
			"Status":  "Failed",
			"Message": "Username not found",
		})

		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "Failed",
			"message": "Password is wrong",
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "Success",
		"message": "Login is success",
	})
}

func CreateUser(db *gorm.DB, c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	hashed, err := bcrypt.GenerateFromPassword([]byte(password), 4)

	if err != nil {
		panic(err)
	}

	data := models.User{
		Username: username,
		Password: string(hashed),
	}

	db.Create(&data)

	c.JSON(http.StatusCreated, gin.H{
		"status":  "Succes",
		"message": "User has been registered",
	})
}
