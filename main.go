package main

import (
	"os"
	"github.com/joho/godotenv"
	"net/http"
	"github.com/gin-gonic/gin"
	routes "github.com/jamyMarkos/gin-auth/routes"
)

func main()  {
	// Load environment variables from .env file
	err := godotenv.Load(".env")

	if err != nil {
		panic("Error loading .env file")
	}
	
	port := os.Getenv("PORT")
	
	if port == "" {
		port = "8080"
	}
	router := gin.New()
	router.Use(gin.Logger())

	// router.GET("/ping", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{
	// 		"message": "pong",
	// 	})
	// })


	routes.AuthRoutes(router)
	routes.UserRoutes(router)

	router.GET("/api-1", func (c *gin.Context) {
		c.JSON(200, gin.H {
			"success": "Access granted for api-1",
		})
	})

	router.GET("/api-2", func (c *gin.Context) {
		c.JSON(200, gin.H {
			"success": "Access granted for api-2",
		})
	})

	router.Run(":" + port)
}