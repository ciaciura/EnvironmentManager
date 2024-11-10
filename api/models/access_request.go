package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AccessRequest struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	ServerID  primitive.ObjectID `bson:"server_id" json:"server_id"`
	Username  string             `bson:"username" json:"username"`
	Status    string             `bson:"status" json:"status"`
	Requested time.Time          `bson:"requested" json:"requested"`
}
