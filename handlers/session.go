package handlers

import (
	"log"
	"net/http"
	"session-service/services"

	"github.com/gin-gonic/gin"
)

// SessionHandler handles GET requests for /sessions
func SessionHandler(c *gin.Context) {
	machineID := c.Query("machine_id")

	// Fetch filtered sessions from the service layer
	sessions, err := services.GetSessionsByMachineID(machineID)
	if err != nil {
		log.Printf("Error fetching sessions: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error reading session data"})
		return
	}

	// Return the filtered sessions as JSON response
	c.JSON(http.StatusOK, sessions)
}
