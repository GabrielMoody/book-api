package handler

import (
	"net/http"
	"strconv"

	"book-api/app/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// var conn = models.GetConnection()

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
	title := c.PostForm("title")
	desc := c.PostForm("description")
	price, _ := strconv.Atoi(c.PostForm("price"))
	rating, _ := strconv.Atoi(c.PostForm("rating"))

	data := models.Book{
		Title:       title,
		Description: desc,
		Price:       price,
		Rating:      rating,
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
	book := c.Param("book")

	title := c.PostForm("title")
	desc := c.PostForm("description")
	price, _ := strconv.Atoi(c.PostForm("price"))
	rating, _ := strconv.Atoi(c.PostForm("rating"))

	data := models.Book{
		Title:       title,
		Description: desc,
		Price:       price,
		Rating:      rating,
	}

	db.Model(&models.Book{}).Where("title = ?", book).Updates(&data)

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
