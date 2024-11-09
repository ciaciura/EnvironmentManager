// handlers/server.go
package handlers

import (
	"EnvManager-api/models"
	"encoding/json"
	"net/http"

	"go.mongodb.org/mongo-driver/mongo"
)

// Assume you have a global variable for the MongoDB collection
var serverCollection *mongo.Collection

// CreateServer godoc
// @Summary Create a new server
// @Description Create a new server with the input payload
// @Tags servers
// @Accept  json
// @Produce  json
// @Param server body models.Server true "Server to create"
// @Success 201 {object} models.Server
// @Failure 400 {object} ErrorResponse
// @Router /servers [post]
func CreateServer(w http.ResponseWriter, r *http.Request) {
	// Only allow if role is Admin
	var server models.Server
	if err := json.NewDecoder(r.Body).Decode(&server); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := server.Save(serverCollection); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(server)
}

// GetServers godoc
// @Summary List all servers
// @Description Get a list of all servers
// @Tags servers
// @Produce  json
// @Success 200 {array} models.Server
// @Router /servers [get]
func GetServers(w http.ResponseWriter, r *http.Request) {
	servers, err := models.GetAllServers(serverCollection)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(servers)
}

// UpdateServer godoc
// @Summary Update server info
// @Description Update the server with the given ID
// @Tags servers
// @Accept  json
// @Produce  json
// @Param id path string true "Server ID"
// @Param server body models.Server true "Server to update"
// @Success 200 {object} models.Server
// @Failure 400 {object} ErrorResponse
// @Router /servers/{id} [put]
func UpdateServer(w http.ResponseWriter, r *http.Request) {
	// Update server info
}

// DeleteServer godoc
// @Summary Delete a server
// @Description Delete the server with the given ID
// @Tags servers
// @Param id path string true "Server ID"
// @Success 204
// @Failure 400 {object} ErrorResponse
// @Router /servers/{id} [delete]
func DeleteServer(w http.ResponseWriter, r *http.Request) {
	// Only allow if role is Admin
	// Delete server info
}
