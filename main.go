package main

import (
	"fmt"
	"log"
	"session-service/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// Define the route for session handler
	router.GET("/sessions", handlers.SessionHandler)

	port := 8083
	log.Printf("Session Service running at http://localhost:%d\n", port)

	// Start the server
	if err := router.Run(fmt.Sprintf(":%d", port)); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
