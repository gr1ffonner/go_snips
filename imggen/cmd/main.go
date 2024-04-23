package main

import (
	"imggen/internal/webserver"
)

func main() {
	webserver.StartServer()
}

// package main
//
// import (
// 	// "imggen/internal/images/getimage"
// 	imgutil "imggen/internal/images/imagegen"
// 	"log"
// 	"os"
// 	// "path/filepath"
// )
//
// func main() {
// 	wd, err := os.Getwd()
// 	if err != nil {
// 		log.Fatalf("Error getting current directory: %v", err)
// 	}
//
// 	// dir := wd + "/images/"
// 	// img_url := "https://www.freepnglogos.com/uploads/lion-logo-png/commercial-real-estate-black-lion-investment-group-los-0.png"
// 	// img_name := filepath.Base(img_url)
// 	// logoImagePath, err := getimage.DownloadImage(img_url, img_name, dir)
// 	bgPath := wd + "/bg" + "/background1.png"
// 	outputPath := wd + "/result.png"
// 	tokenName := "Literally a motherfking NFT PAAP"
// 	tokenAPY := "11.5"
// 	tokenDays := "250"
// 	fontPath := wd + "/PtSerif_Regular.ttf"
// 	logoImagePath := wd + "/bg" + "/default_logo.png"
// 	_, err = imgutil.CreateImageFile(logoImagePath, bgPath, outputPath, tokenName, tokenAPY, tokenDays, fontPath)
// 	if err != nil {
// 		log.Printf("Error generating image content: %v\nRetrying...", err)
// 		_, _ = imgutil.CreateImageFile(wd+"/bg"+"/default_logo.png", bgPath, outputPath, tokenName, tokenAPY, tokenDays, fontPath)
// 	}
//
// 	log.Println("Image generated")
// }
