package book

import (
	"net/http"
	"strconv"

	"book-api/models"

	"github.com/gin-gonic/gin"
)

var conn = models.GetConnection()

func GetIndexHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"name": "gabriel",
	})
}

func GetBooks(c *gin.Context) {
	var book []models.Book
	conn.Find(&book)

	c.JSON(http.StatusOK, book)
}

func GetBookByTitle(c *gin.Context) {
	var book []models.Book
	title := c.Param("title")

	conn.Where("title = ?", title).Find(&book)

	c.JSON(http.StatusOK, book)
}

func PostBook(c *gin.Context) {
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

	conn.Create(&data)

	c.JSON(http.StatusCreated, gin.H{
		"ID":          data.ID,
		"title":       data.Title,
		"description": data.Description,
		"price":       data.Price,
		"rating":      data.Rating,
	})
}

func UpdateBook(c *gin.Context) {
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

	conn.Model(&models.Book{}).Where("title = ?", book).Updates(&data)

	c.JSON(http.StatusOK, &data)
}

func DeleteBook(c *gin.Context) {
	title := c.Param("title")

	var book models.Book

	conn.Where("title = ?", title).Delete(&book)

	c.JSON(http.StatusOK, gin.H{
		"message": "Book has been deleted",
	})
}
