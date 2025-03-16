package main

import (
	"fmt"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.POST("/login", func(c *gin.Context) {
		token := jwt.New(jwt.SigningMethodNone)

		claims := token.Claims.(jwt.MapClaims)
		claims["user"] = "admin"
		claims["exp"] = 15000

		tokenString, err := token.SigningString()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"token": tokenString,
		})
	})

	fmt.Println("Server is running on :8080")
	r.Run(":8080")
}
