package middleware

import (
	"net/http"

	"book-api/models"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func BasicAuth(c *gin.Context) {
	username, password, ok := c.Request.BasicAuth()
	user := models.GetUserByUsername(username)

	if !ok || user == nil {
		c.Abort()
		c.Writer.Header().Set("WWW-Authecticate", "Basic realm=Restricted")
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "Failed",
			"message": "Authentication is required",
		})
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))

	if err != nil {
		c.Abort()
		c.Writer.Header().Set("WWW-Authecticate", "Basic realm=Restricted")
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "Failed",
			"message": "Authentication is required",
		})
	}
}
