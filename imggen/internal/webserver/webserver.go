package webserver

import (
	"encoding/json"
	"fmt"
	"imggen/internal/config"
	"imggen/internal/db"
	"imggen/internal/images/getimage"
	"imggen/internal/upload"
	"log"
	"net/http"
	"os"
	"strconv"
	"sync"

	imgutil "imggen/internal/images/imagegen"
	compressor "imggen/internal/images/press"
	models "imggen/internal/nft"
)

func stringToInt(s string) (int, error) {
	// Convert the string to an integer
	i, err := strconv.Atoi(s)
	if err != nil {
		return 0, err // Return 0 and the error if conversion fails
	}
	return i, nil // Return the integer value and no error if conversion succeeds
}

func floatToString(f float64) string {
	// Convert float to string with 1 decimal place
	return strconv.FormatFloat(f, 'f', 1, 64)
}

func StartServer() {
	http.HandleFunc("/api/post", handlePost)
	port := ":8080"
	log.Printf("Server listening on localhost%s\n", port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

func handlePost(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	log.Println("Received POST request")

	var tokenInfo models.TokenInfo

	// Decode JSON request body into the TokenInfo struct
	err := json.NewDecoder(r.Body).Decode(&tokenInfo)
	if err != nil {
		http.Error(w, "Error decoding JSON request body", http.StatusBadRequest)
		return
	}
	// Send 200 response status code immediately
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Request processed successfully json")
	go func() {
		defer func() {
			if r := recover(); r != nil {
				log.Printf("Recovered from panic: %v", r)
			}
		}()

		wd, err := os.Getwd()
		if err != nil {
			log.Fatalf("Error getting current directory: %v", err)
		}

		dir := wd + "/images/"

		tokenName := tokenInfo.TokenName
		logoImagePath, err := getimage.DownloadImage(tokenInfo.TokenImg, tokenName, dir)
		if err != nil {
			log.Printf("Error downloading logo: %v", err)
			logoImagePath = wd + "/bg/" + "default_logo.png"
			log.Println("Image not found, using default logo")
		} else {
			log.Println("Image downloaded")
		}

		// Print information about tariffs
		// Create a wait group to synchronize the goroutines
		var wg sync.WaitGroup

		var tariffsData models.TariffsData
		mu := &sync.Mutex{}
		// Iterate over tariffs and generate images concurrently

		for _, tariff := range tokenInfo.Tariffs {
			wg.Add(1) // Increment wait group counter
			go func(tariff models.Tariff) {
				defer wg.Done() // Decrement wait group counter when done
				resultImageName := "nft" + "_" + tariff.LockPeriodDays + ".png"
				// Generate the image and get the image content
				imgContent, err := imgutil.CreateImage(logoImagePath, wd+"/bg/"+tokenInfo.NFTBackground+".png", tokenName, floatToString(tariff.APY), tariff.LockPeriodDays, wd+"/PtSerif_Regular.ttf")
				if err != nil {
					log.Printf("Error generating image content: %v\nRetrying...", err)
					imgContent, _ = imgutil.CreateImage(wd+"/bg/"+"/default_logo.png", wd+"/bg/"+tokenInfo.NFTBackground+".png", tokenName, floatToString(tariff.APY), tariff.LockPeriodDays, wd+"/PtSerif_Regular.ttf")
				}
				log.Println("Image generated")
				compressedImageURL, err := compressor.Compress(imgContent, resultImageName)
				if err != nil {
					log.Printf("Error compressing image: %v", err)
					return
				}

				compressedImagePath, err := getimage.DownloadImage(compressedImageURL, resultImageName, dir)
				if err != nil {
					log.Fatalf("Error downloading compressed image: %v", err)
				}

				compressedImageFile, err := os.Open(compressedImagePath)
				if err != nil {
					log.Fatalf("Error opening image file: %v", err)
				}
				log.Println("Image compressed")

				// Upload the generated image content
				apiKey := config.GetAPIKeyNFT()
				imageURL, err := upload.UploadImage(apiKey, resultImageName, compressedImageFile)
				if err != nil {
					log.Printf("Error uploading image: %v", err)
				}
				log.Println("Image uploaded into nft storage")
				nft := models.NewNFT(tokenInfo.TokenName, tokenInfo.TokenSymbol, imageURL)
				jsonURL, err := upload.UploadJSON(apiKey, *nft)
				if err != nil {
					log.Printf("Error uploading JSON: %v", err)
					return
				}
				dur, err := stringToInt(tariff.LockPeriodDays)
				if err != nil {
					log.Printf("Error converting duration: %v", err)
				}
				tariffData := models.Tariffs{
					Duration: dur,
					URL:      jsonURL,
				}

				// Use a mutex to protect access to tariffsData
				mu.Lock()
				tariffsData.Tariffs = append(tariffsData.Tariffs, tariffData)
				log.Printf("Tariff data added: %v", tariffData)
				mu.Unlock()

				// Update the database with the tariff data
			}(tariff)
		}

		// Wait for all goroutines to finish
		wg.Wait()

		db, err := db.NewDB()
		if err != nil {
			log.Fatalf("Error creating DB instance: %v", err)
		}
		defer db.Close()

		err = db.UpdateTariffData(tokenInfo.TokenAddr, tariffsData)
		if err != nil {
			log.Printf("Error updating tariff data: %v", err)
		}

		log.Println("Tariffs data updated in the database")
		fmt.Fprintln(w, "Request processed successfully")
	}()
}
