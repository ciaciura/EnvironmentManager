package config

import (
	"os"

	"github.com/joho/godotenv"
)

var (
	JWT_SECRET       = os.Getenv("JWT_SECRET")
	MONGODB_URI      = os.Getenv("MONGODB_URI")
	MONGODB_DATABASE = os.Getenv("MONGODB_DATABASE")
	ADMIN_USERNAME   = os.Getenv("ADMIN_USERNAME")
	ADMIN_PASSWORD   = os.Getenv("ADMIN_PASSWORD")
	APP_PORT         = ":" + os.Getenv("APP_PORT")
	Env              = os.Getenv("Env")
	IsDocker         = os.Getenv("DOCKER_ENV")
)

func init() {
	if IsDocker == "false" {
		// Load environment variables from .env-debug file
		err := godotenv.Overload(".env-debug")
		if err != nil {
			panic(err)
		} else {
			JWT_SECRET = os.Getenv("JWT_SECRET")
			MONGODB_URI = os.Getenv("MONGODB_URI")
			MONGODB_DATABASE = os.Getenv("MONGODB_DATABASE")
			ADMIN_USERNAME = os.Getenv("ADMIN_USERNAME")
			ADMIN_PASSWORD = os.Getenv("ADMIN_PASSWORD")
			APP_PORT = ":" + os.Getenv("APP_PORT")
			Env = os.Getenv("Env")
		}
	}
}
