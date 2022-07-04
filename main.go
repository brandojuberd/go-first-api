package main

import (
	"fmt"
	"go-first-api/models"
	"go-first-api/models/users"
	"go-first-api/models/weapons"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var port = "3000"
var ginMode = "development"

func main() {

	godotenv.Load("development.env")
	port = os.Getenv("PORT")
	ginMode = os.Getenv("GIN_MODE")
	fmt.Println(port, ginMode)
	gin.SetMode(ginMode)

	fmt.Println("Welcome to Go First API")

	// Initiate database
	models.InitDatabaseConnection()

	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		firstName := c.DefaultQuery("firstName", "Guest")
		c.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("Welcome, %v", firstName),
		})
	})

	weapons.Init(&router.RouterGroup)
	users.Init(&router.RouterGroup)

	router.Run("localhost:" + port)
}


