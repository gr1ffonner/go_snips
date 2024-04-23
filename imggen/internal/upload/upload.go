package upload

import (
	"bytes"
	"encoding/json"
	"fmt"
	models "imggen/internal/nft"
	"io"
	"mime/multipart"
	"net/http"
)

func UploadImage(apiKey, filename string, file io.Reader) (string, error) {
	url := "https://api.nft.storage/upload"

	client := &http.Client{}

	// Create a new multipart writer
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	// Write the file to the multipart writer
	part, err := writer.CreateFormFile("file", filename)
	if err != nil {
		return "", err
	}
	_, err = io.Copy(part, file)
	if err != nil {
		return "", err
	}

	// Close the multipart writer and get the boundary string
	err = writer.Close()
	if err != nil {
		return "", err
	}

	// Create a new request
	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		return "", err
	}

	// Set the Authorization header with your API key
	req.Header.Set("Authorization", "Bearer "+apiKey)

	// Set the Content-Type header to multipart/form-data with the boundary string
	req.Header.Set("Content-Type", writer.FormDataContentType())

	// Send the request
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// Read the response body
	var response models.Response
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return "", err
	}

	cidUrl := fmt.Sprintf("https://%s.ipfs.nftstorage.link/%s", response.Value.Pin.Cid, filename)

	return cidUrl, nil
}

func UploadJSON(apiKey string, data models.NFT) (string, error) {
	url := "https://api.nft.storage/upload"

	jsonData, err := json.Marshal(data)
	if err != nil {
		return "", err
	}

	client := &http.Client{}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return "", err
	}

	req.Header.Set("Authorization", "Bearer "+apiKey)
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var response models.Response
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return "", err
	}

	cidUrl := fmt.Sprintf("https://%s.ipfs.nftstorage.link/", response.Value.Pin.Cid)

	return cidUrl, nil
}
