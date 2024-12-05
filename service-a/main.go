package main

import (
	"log"
	"net/http"
	"strings"

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
		{ID: 1, Class: "Economy", Price: 500000},
		{ID: 2, Class: "Business", Price: 1500000},
	},
	"SURABAYA": {
		{ID: 1, Class: "Economy", Price: 400000},
		{ID: 2, Class: "Business", Price: 1200000},
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
	router.GET("/api/v1/ticket/:destination", func(c *gin.Context) {
		destination := strings.ToUpper(c.Param("destination"))

		options, exists := ticketOptions[destination]
		if !exists {
			log.Printf("Destination not found: %s", destination)
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
	log.Println("Server is running on http://0.0.0.0:4000")
	if err := router.Run(":4000"); err != nil {
		log.Fatalf("Error starting server: %s", err)
	}
}
