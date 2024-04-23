package config

import (
	"log"
	"os"
)

// GetAPIKey retrieves the API key from the environment variables
func GetAPIKey() string {
	apiKey, exists := os.LookupEnv("API_KEY")
	if !exists {
		log.Fatal("API_KEY not found in environment variables")
	}
	return apiKey
}

func GetDBCred() [5]string {
	dbHost, exists := os.LookupEnv("DB_HOST")
	if !exists {
		log.Fatal("DB_HOST not found in environment variables")
	}
	dbUser, exists := os.LookupEnv("DB_USER")
	if !exists {
		log.Fatal("DB_USER not found in environment variables")
	}
	dbPass, exists := os.LookupEnv("DB_PASSWORD")
	if !exists {
		log.Fatal("DB_PASS not found in environment variables")
	}
	dbPort, exists := os.LookupEnv("DB_PORT")
	if !exists {
		log.Fatal("DB_PORT not found in environment variables")
	}
	dbName, exists := os.LookupEnv("DB_NAME")
	if !exists {
		log.Fatal("DB_PORT not found in environment variables")
	}
	db := [5]string{dbHost, dbUser, dbPass, dbPort, dbName}
	return db
}
