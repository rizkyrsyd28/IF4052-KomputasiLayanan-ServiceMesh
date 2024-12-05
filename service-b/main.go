package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// TicketOption represents a ticket's data
type TicketOption struct {
	ID    int    `json:"id"`
	Class string `json:"class"`
	Price int    `json:"price"`
}

// ResponseModel represents the API response structure
type ResponseModel struct {
	Success bool           `json:"success"`
	Data    []TicketOption `json:"data"`
}

// Sample ticket options
var ticketOptions = map[string][]TicketOption{
	"JAKARTA": {
		{ID: 1, Class: "Premium", Price: 1900000},
		{ID: 2, Class: "VIP", Price: 12300000},
	},
	"SURABAYA": {
		{ID: 1, Class: "Premium", Price: 6400000},
		{ID: 2, Class: "VIP", Price: 19300000},
	},
}

func main() {
	// Create a new Gin router
	router := gin.Default()

	// Middleware to log incoming requests
	router.Use(func(c *gin.Context) {
		log.Printf("Incoming request: method=%s, path=%s, query=%s",
			c.Request.Method, c.Request.URL.Path, c.Request.URL.RawQuery)
		c.Next()
	})


	// Route to handle ticket requests
	router.POST("/api/v-1/ticket", func(c *gin.Context) {

		destinationHeader := c.GetHeader("x-destination")

		if destinationHeader == "" {
			log.Printf("Destination header is missing")
			c.JSON(http.StatusBadRequest, gin.H{"message": "Destination header is missing"})
			return
		}

		options, exists := ticketOptions[destinationHeader]
		if !exists {
			log.Printf("Destination not found: %s", destinationHeader)
			c.JSON(http.StatusNotFound, gin.H{"message": "Destination not found"})
			return
		}

		response := ResponseModel{
			Success: true,
			Data:    options,
		}
		c.JSON(http.StatusOK, response)
	})

	// Custom handler for 404 errors
	router.NoRoute(func(c *gin.Context) {
		log.Printf("Endpoint not found: method=%s, path=%s", c.Request.Method, c.Request.URL.Path)
		c.JSON(http.StatusNotFound, gin.H{"message": "Endpoint not found"})
	})

	// Run the server
	log.Println("Server is running on http://0.0.0.0:4001")
	if err := router.Run(":4001"); err != nil {
		log.Fatalf("Error starting server: %s", err)
	}
}
