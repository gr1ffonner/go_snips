package config

import (
	"log"
	"os"
)

// GetAPIKey retrieves the API key from the environment variables
func GetAPIKeyNFT() string {
	apiKey, exists := os.LookupEnv("API_KEY_NFT")
	if !exists {
		log.Fatal("API_KEY not found in environment variables")
	}
	return apiKey
}

func GetHostName() string {
	hostname, exists := os.LookupEnv("DB_HOST")
	if !exists {
		log.Fatal("API_KEY not found in environment variables")
	}
	return hostname
}

func GetAPIKeyKraken() string {
	apiKey, exists := os.LookupEnv("API_KEY_KRAKEN")
	if !exists {
		log.Fatal("API_KEY_KRAKEN not found in environment variables")
	}
	return apiKey
}

func GetAPISecretKraken() string {
	apiSecret, exists := os.LookupEnv("API_SECRET_KRAKEN")
	if !exists {
		log.Fatal("API_SECRET_KRAKEN not found in environment variables")
	}
	return apiSecret
}
