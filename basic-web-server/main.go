package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// Create a Gin router instance
	router := gin.Default()

	// Define a route handler for the root path "/"
	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello, Gin!")
	})

	// Define a route handler for the "/hello" path
	router.GET("/hello", func(c *gin.Context) {
		name := c.Query("name") // Get the value of the 'name' query parameter
		c.String(http.StatusOK, "Hello, %s!", name)
	})

	// Run the server on port 8080
	router.Run(":8080")
}
