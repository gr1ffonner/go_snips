package main

import (
	"pricer/internal/updater"
	"time"
)

const interval = 5 * time.Minute

func main() {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	// Run the initial update
	updater.UpdatePrices()

	// Start a goroutine to handle updates triggered by the ticker
	go func() {
		for range ticker.C {
			updater.UpdatePrices()
		}
	}()

	// Block main goroutine to keep the program running
	select {}
}
