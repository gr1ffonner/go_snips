package getpricer

import (
	"io"
	"log"
	"net/http"
	"pricer/internal/config"
	"pricer/internal/models"
)

func GetPrice(url string) ([]models.Response, error) {
	req, _ := http.NewRequest("GET", url, nil)
	apiKey := config.GetAPIKey()

	req.Header.Add("X-API-KEY", apiKey)
	req.Header.Add("x-chain", "solana")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	responses, err := models.ParseResponse(body)
	if err != nil {
		return nil, err
	}

	return responses, nil
}
