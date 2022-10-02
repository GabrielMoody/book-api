package handler

import (
	"book-api/app/models"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Register struct {
	Email                 string `json:"email" form:"email" binding:"required,email"`
	Username              string `json:"username" form:"username" binding:"required"`
	Password              string `json:"password" form:"password" binding:"required,min=8"`
	Password_Confirmation string `json:"password_confirmation" form:"password_confirmation" binding:"required,eqfield=Password"`
}

type Login struct {
	Username string `json:"username" form:"username" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
}

type ErrorMsg struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func UserLogin(db *gorm.DB, c *gin.Context) {
	user := Login{}

	err := c.ShouldBind(&user)

	if err != nil {
		var ve validator.ValidationErrors

		if errors.As(err, &ve) {
			out := make([]ErrorMsg, len(ve))
			for i, e := range ve {
				out[i] = ErrorMsg{Field: e.Field(), Message: e.Error()}
			}

			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"errora": out,
			})
		}

		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	result := models.GetUserByUsername(db, user.Username)

	if result == nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Status": "Failed",
			"Error":  "Username not found",
		})

		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(result.Password), []byte(user.Password))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": "Password is wrong",
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": user.Username,
	})
}

func CreateUser(db *gorm.DB, c *gin.Context) {
	user := Register{}

	err := c.ShouldBind(&user)

	if err != nil {
		var ve validator.ValidationErrors

		if errors.As(err, &ve) {
			errors := make([]ErrorMsg, len(ve))

			for i, e := range ve {
				errors[i] = ErrorMsg{Field: e.Field(), Message: e.Error()}
			}

			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"errors": errors,
			})
		}

		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	hashed, err := bcrypt.GenerateFromPassword([]byte(user.Password), 4)

	if err != nil {
		panic(err)
	}

	data := models.User{
		Username: user.Username,
		Email:    user.Email,
		Password: string(hashed),
	}

	db.Create(&data)

	c.JSON(http.StatusCreated, gin.H{
		"message": "User has been registered",
	})
}
