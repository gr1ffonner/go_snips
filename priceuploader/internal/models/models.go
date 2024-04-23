package models

import (
	"encoding/json"
	"pricer/internal/rounding"
	"strconv"
)

type DataItem struct {
	Value           float64 `json:"value"`
	UpdateUnixTime  int64   `json:"updateUnixTime"`
	UpdateHumanTime string  `json:"updateHumanTime"`
	PriceChange24h  float64 `json:"priceChange24h"`
}

type Response struct {
	Data    map[string]DataItem `json:"data"`
	Success bool                `json:"success"`
}

func stringToFloat(str string) (float64, error) {
	f, err := strconv.ParseFloat(str, 64)
	if err != nil {
		return 0, err
	}
	return f, nil
}

func ParseResponse(body []byte) ([]Response, error) {
	var resp Response

	if err := json.Unmarshal(body, &resp); err != nil {
		return nil, err
	}

	result := make(map[string]DataItem)
	// Round the values and price changes
	for key, item := range resp.Data {
		item.Value = rounding.RoundNumber(item.Value)
		item.PriceChange24h = rounding.RoundNumber(item.PriceChange24h)
		result[key] = item
	}

	resp.Data = result
	return []Response{resp}, nil
}
