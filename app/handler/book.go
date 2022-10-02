package handler

import (
	"errors"
	"net/http"

	"book-api/app/models"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type Book struct {
	Title       string `json:"title" form:"title" binding:"required"`
	Description string `json:"description" form:"description" binding:"required"`
	Price       int    `json:"price" form:"price" binding:"required"`
	Rating      int    `json:"rating" form:"rating" binding:"required"`
}

func GetIndexHandler(db *gorm.DB, c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"name": "gabriel",
	})
}

func GetBooks(db *gorm.DB, c *gin.Context) {
	var book []models.Book
	db.Find(&book)

	c.JSON(http.StatusOK, book)
}

func GetBookByTitle(db *gorm.DB, c *gin.Context) {
	var book []models.Book
	title := c.Param("title")

	db.Where("title = ?", title).Find(&book)

	c.JSON(http.StatusOK, book)
}

func PostBook(db *gorm.DB, c *gin.Context) {
	book := Book{}

	err := c.ShouldBind(&book)

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

	data := models.Book{
		Title:       book.Title,
		Description: book.Description,
		Price:       book.Price,
		Rating:      book.Rating,
	}

	db.Create(&data)

	c.JSON(http.StatusCreated, gin.H{
		"ID":          data.ID,
		"title":       data.Title,
		"description": data.Description,
		"price":       data.Price,
		"rating":      data.Rating,
	})
}

func UpdateBook(db *gorm.DB, c *gin.Context) {
	title := c.Param("book")

	book := Book{}

	err := c.ShouldBind(&book)

	if err != nil {
		var ve validator.ValidationErrors

		if errors.As(err, &ve) {
			errors := make([]ErrorMsg, len(ve))

			for i, e := range ve {
				errors[i] = ErrorMsg{Field: e.Field(), Message: e.Error()}
			}

			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"erorrs": errors,
			})

			c.AbortWithStatus(http.StatusBadRequest)
			return
		}
	}

	data := models.Book{
		Title:       book.Title,
		Description: book.Description,
		Price:       book.Price,
		Rating:      book.Rating,
	}

	db.Model(&models.Book{}).Where("title = ?", title).Updates(&data)

	c.JSON(http.StatusOK, &data)
}

func DeleteBook(db *gorm.DB, c *gin.Context) {
	title := c.Param("title")

	var book models.Book

	db.Where("title = ?", title).Delete(&book)

	c.JSON(http.StatusOK, gin.H{
		"message": "Book has been deleted",
	})
}
