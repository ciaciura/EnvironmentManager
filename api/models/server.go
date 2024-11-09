// models/server.go
package models

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

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

func (s *Server) Save(collection *mongo.Collection) error {
	_, err := collection.InsertOne(context.TODO(), s)
	return err
}

func GetAllServers(collection *mongo.Collection) ([]Server, error) {
	var servers []Server
	cursor, err := collection.Find(context.TODO(), bson.D{})
	if err != nil {
		return nil, err
	}
	if err = cursor.All(context.TODO(), &servers); err != nil {
		return nil, err
	}
	return servers, nil
}
