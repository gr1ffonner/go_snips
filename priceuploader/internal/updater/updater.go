package updater

import (
	"log"
	"pricer/internal/db"
	"pricer/internal/getpricer"
	"time"
)

const (
	layout = "2006-01-02T15:04:05"
)

func UpdatePrices() {
	d, err := db.NewDB()
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}
	defer d.Close()
	addrs, err := d.GetTokenAddresses()
	if err != nil {
		log.Fatalf("Error getting token addresses: %v", err)
		return
	}

	url := "https://public-api.birdeye.so/public/multi_price?list_address="
	for _, v := range addrs {
		url += v + "%2C"
	}
	url = url[:len(url)-3]

	responses, err := getpricer.GetPrice(url)
	if err != nil {
		log.Fatal(err)
	}

	for _, resp := range responses {
		for key, item := range resp.Data {
			t, err := time.Parse(layout, item.UpdateHumanTime)
			if err != nil {
				log.Printf("Error parsing time: %v", err)
				continue
			}
			err = d.UpsertTokenPrice(key, item.Value, item.PriceChange24h, t)
			if err != nil {
				log.Fatalf("Error updating token price %s: %v", key, err)
				continue
			}
			log.Printf("Updated token price %s\n", key)
		}
	}
}
