package main

import (
	"net/http"

	"example.com/crypto-tracker/models"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// Basic test route
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})

	// GET /portfolio route
	router.GET("/portfolio", func(c *gin.Context) {
		c.JSON(http.StatusOK, models.Portfolio)
	})

	router.POST("/addcoin", func(c *gin.Context) {
		var newCoin models.Coin

		if err := c.ShouldBindJSON(&newCoin); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		models.Portfolio = append(models.Portfolio, newCoin)
		c.JSON(http.StatusCreated, newCoin)
	})

	router.Run(":8080")
}
