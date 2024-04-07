package config

import (
	"os"

	"github.com/joho/godotenv"
)

// Config func to get env value
func Config(key string) string {
	// load .env file
	err := godotenv.Load(".env.sample")
	if err != nil {
		err = godotenv.Load("../.env.sample")
		if err != nil {
			err = godotenv.Load("../../.env.sample")
			if err != nil {
				err = godotenv.Load("./.env.sample")
			}
		}
	}

	return os.Getenv(key)
}
