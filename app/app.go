package app

import (
	"book-api/config"
	"fmt"
	"log"

	"book-api/app/handler"
	"book-api/app/middleware"
	"book-api/app/models"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type App struct {
	Router *gin.Engine
	DB     *gorm.DB
}

func (a *App) Initialize(config *config.Config) {
	dbURI := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local",
		config.DB.Username,
		config.DB.Password,
		config.DB.Host,
		config.DB.Port,
		config.DB.Name,
		config.DB.Charset)

	db, err := gorm.Open(mysql.Open(dbURI), &gorm.Config{})

	if err != nil {
		log.Fatal("Could no connect to database")
	}

	a.DB = models.DBMigrate(db)
	a.Router = gin.Default()
	a.SetRouters()
}

func (a *App) SetRouters() {
	a.Router.GET("/", middleware.BasicAuth, a.handleRequest(handler.GetIndexHandler))

	authorized := a.Router.Group("/")

	authorized.Use(middleware.BasicAuth)
	{
		authorized.GET("/books", a.handleRequest(handler.GetBooks))
		authorized.POST("/book", a.handleRequest(handler.PostBook))
		authorized.GET("/book/:title", a.handleRequest(handler.GetBookByTitle))
		authorized.PUT("/book/:book", a.handleRequest(handler.UpdateBook))
		authorized.DELETE("/book/:title", a.handleRequest(handler.DeleteBook))
	}

	a.Router.POST("/register", a.handleRequest(handler.CreateUser))
	a.Router.POST("/login", a.handleRequest(handler.Login))
}

type RequestHandlerFunction func(db *gorm.DB, c *gin.Context)

func (a *App) handleRequest(handler RequestHandlerFunction) gin.HandlerFunc {
	return func(c *gin.Context) {
		handler(a.DB, c)
	}
}
