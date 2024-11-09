// models/user.go
package models

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Role string

const (
	Admin    Role = "Admin"
	Employee Role = "Employee"
)

type User struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Role     Role   `json:"role"`
}

func (u *User) Save(collection *mongo.Collection) error {
	_, err := collection.InsertOne(context.TODO(), u)
	return err
}

func GetUserByID(collection *mongo.Collection, id string) (*User, error) {
	var user User
	err := collection.FindOne(context.TODO(), bson.M{"id": id}).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
