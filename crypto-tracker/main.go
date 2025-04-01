package main

import (
	"net/http"

	"example.com/crypto-tracker/models"
	"example.com/crypto-tracker/utils"
	"golang.org/x/crypto/bcrypt"

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

	router.POST("/register", func(c *gin.Context) {
		var req struct {
			Username string `json:"username"`
			Password string `json:"password"`
		}

		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
		if err != nil {
			c.JSON(500, gin.H{"error": "Failed to hash password"})
			return
		}

		user := models.User{Username: req.Username, Password: string(hashedPassword)}
		if err := DB.Create(&user).Error; err != nil {
			c.JSON(400, gin.H{"error": "Username already exists"})
			return
		}

		c.JSON(201, gin.H{"message": "User created"})
	})

	router.POST("/login", func(c *gin.Context) {
		var req struct {
			Username string `json:"username"`
			Password string `json:"password"`
		}

		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		var user models.User
		if err := DB.Where("username = ?", req.Username).First(&user).Error; err != nil {
			c.JSON(401, gin.H{"error": "Invalid username or password"})
			return
		}

		if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
			c.JSON(401, gin.H{"error": "Invalid username or password"})
			return
		}

		token, err := utils.GenerateToken(user.ID)
		if err != nil {
			c.JSON(500, gin.H{"error": "Failed to generate token"})
			return
		}

		c.JSON(200, gin.H{"token": token})
	})

	router.Run("localhost:8080")
}
