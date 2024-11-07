// models/server.go
package models

type Environment string

const (
	Prod Environment = "PROD"
	Dev  Environment = "DEV"
	Uat  Environment = "UAT"
)

type Server struct {
	ID          string      `json:"id"`
	IP          string      `json:"ip"`
	Name        string      `json:"name"`
	Environment Environment `json:"environment"`
}
