package main

import (
	"book-api/book"
	"book-api/middleware"
	"book-api/user"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/", middleware.BasicAuth, book.GetIndexHandler)

	authorized := router.Group("/")

	authorized.Use(middleware.BasicAuth)
	{
		authorized.GET("/books", book.GetBooks)
		authorized.POST("/book", book.PostBook)
		authorized.GET("/book/:title", book.GetBookByTitle)
		authorized.PUT("/book/:book", book.UpdateBook)
		authorized.DELETE("/book/:title", book.DeleteBook)
	}

	router.POST("/register", user.CreateUser)
	router.POST("/login", user.Login)

	router.Run("localhost:3000")
}
