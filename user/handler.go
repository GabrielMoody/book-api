package user

import (
	"book-api/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

var conn = models.GetConnection()

func Login(c *gin.Context) {
	var user models.User
	username := c.PostForm("username")
	password := c.PostForm("password")
	result := conn.Where("username = ?", username).Find(&user)

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

func CreateUser(c *gin.Context) {
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

	conn.Create(&data)

	c.JSON(http.StatusCreated, gin.H{
		"status":  "Succes",
		"message": "User has been registered",
	})
}
