package config

import "os"

var JWTSecret = os.Getenv("JWT_SECRET")
var MongoDBURI = os.Getenv("MONGODB_URI")
var MongoDBDatabase = os.Getenv("MONGODB_DATABASE")
var AdminUsername = os.Getenv("ADMIN_USERNAME")
var AdminPassword = os.Getenv("ADMIN_PASSWORD")
var AppPort = ":" + os.Getenv("APP_PORT")
var Env = os.Getenv("Env")
