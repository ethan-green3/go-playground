package main

import (
	"net/http"

	"example.com/crypto-tracker/models"

	"github.com/gin-gonic/gin"
)

func main() {
	InitDatabase()

	router := gin.Default()

	// Basic test route
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})

	// GET /portfolio route
	router.GET("/portfolio", func(c *gin.Context) {
		var coins []models.Coin
		result := DB.Find(&coins)
		if result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
			return
		}

		c.JSON(http.StatusOK, coins)
	})

	router.POST("/addcoin", func(c *gin.Context) {
		var coin models.Coin
		if err := c.ShouldBindJSON(&coin); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		result := DB.Create(&coin)
		if result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
			return
		}

		c.JSON(http.StatusCreated, coin)
	})

	router.DELETE("/portfolio/:symbol", func(c *gin.Context) {
		symbol := c.Param("symbol")
		result := DB.Where("symbol = ?", symbol).Delete(&models.Coin{})

		if result.RowsAffected == 0 {
			c.JSON(http.StatusNotFound, gin.H{"error": "Coin not found"})
			return
		} else if result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": symbol + " removed"})
	})

	router.Run("localhost:8080")
}
