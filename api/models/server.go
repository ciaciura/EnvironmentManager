// models/server.go
package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Server struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	IP          string             `bson:"ip" json:"ip"`
	Name        string             `bson:"name" json:"name"`
	Environment string             `bson:"environment" json:"environment"`
}
