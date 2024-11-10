package database

import (
	"context"
	"log"

	"EnvManager-api/config"
	"EnvManager-api/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type DBClient struct {
	Client   *mongo.Client
	Database string
}

func (db *DBClient) GetCollection(collectionName string) *mongo.Collection {
	return db.Client.Database(db.Database).Collection(collectionName)
}

func (db *DBClient) InsertOne(ctx context.Context, collectionName string, document interface{}) (*mongo.InsertOneResult, error) {
	collection := db.GetCollection(collectionName)
	log.Printf("Inserting document into collection: %s", collectionName)
	result, err := collection.InsertOne(ctx, document)
	if err != nil {
		log.Printf("Error inserting document: %v", err)
	}
	return result, err
}

func (db *DBClient) UpdateOne(ctx context.Context, collectionName string, filter interface{}, update interface{}) (*mongo.UpdateResult, error) {
	collection := db.GetCollection(collectionName)
	log.Printf("Updating document in collection: %s", collectionName)
	result, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		log.Printf("Error updating document: %v", err)
	}
	return result, err
}

func (db *DBClient) EnsureDefaultAdmin(ctx context.Context) {
	filter := bson.M{"username": config.ADMIN_USERNAME}
	var existingUser models.User
	collection := db.GetCollection("users")
	err := collection.FindOne(ctx, filter).Decode(&existingUser)
	if err == mongo.ErrNoDocuments {
		// No user found, create the default admin
		defaultAdmin := models.User{
			Username: config.ADMIN_USERNAME,
			Password: config.ADMIN_PASSWORD,
			Role:     "admin",
		}
		_, err := db.InsertOne(ctx, "users", defaultAdmin)
		if err != nil {
			log.Fatalf("Error creating default admin user: %v", err)
		}
		log.Println("Default admin user created")
	} else if err != nil {
		log.Fatalf("Error checking for default admin user: %v", err)
	} else {
		log.Println("Default admin user already exists")
	}
}
