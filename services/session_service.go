package services

import (
	"encoding/json"
	"os"
	"session-service/models"
)

// GetSessions fetches all sessions from the JSON file
func GetSessions() ([]models.Session, error) {
	file, err := os.Open("services/sessions.json")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var sessions []models.Session
	if err := json.NewDecoder(file).Decode(&sessions); err != nil {
		return nil, err
	}
	return sessions, nil
}

// GetSessionsByMachineID fetches sessions filtered by machine ID
func GetSessionsByMachineID(machineID string) ([]models.Session, error) {
	sessions, err := GetSessions()
	if err != nil {
		return nil, err
	}

	// Filter sessions based on the machineID
	var filteredSessions []models.Session
	for _, session := range sessions {
		if session.MachineID == machineID {
			filteredSessions = append(filteredSessions, session)
		}
	}
	return filteredSessions, nil
}
