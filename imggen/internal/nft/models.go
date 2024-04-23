package models

import (
	"image/color"
	"strings"
)

type NFT struct {
	Name        string `json:"name"`
	Symbol      string `json:"symbol"`
	Description string `json:"description"`
	Image       string `json:"image"`
	ExternalURL string `json:"external_url"`
	Collection  struct {
		Name   string `json:"name"`
		Family string `json:"family"`
	} `json:"collection"`
}

type Pin struct {
	Cid string `json:"cid"`
}

type Response struct {
	Value struct {
		Cid string `json:"cid"`
		Pin Pin    `json:"pin"`
	} `json:"value"`
}

type Tariff struct {
	LockPeriodDays string  `json:"lock_period_days"`
	APY            float64 `json:"apy"`
}

type Tariffs struct {
	Duration int    `json:"duration"`
	URL      string `json:"url"`
}

type TariffsData struct {
	Tariffs []Tariffs `json:"tariffs"`
}

type TokenInfo struct {
	TokenName     string   `json:"tokenName"`
	TokenSymbol   string   `json:"tokenSymbol"`
	TokenAddr     string   `json:"tokenAddr"`
	TokenImg      string   `json:"tokenImg"`
	NFTBackground string   `json:"NFTBackground"`
	Tariffs       []Tariff `json:"tariffs"`
}

var ColorMap = map[string]color.RGBA{
	"background1":  {255, 255, 255, 255},
	"background2":  {255, 255, 255, 255},
	"background3":  {255, 255, 255, 255},
	"background4":  {255, 255, 255, 255},
	"background5":  {255, 255, 255, 255},
	"background6":  {255, 255, 255, 255},
	"background7":  {255, 255, 255, 255},
	"background8":  {255, 255, 255, 255},
	"background9":  {255, 255, 255, 255},
	"background10": {255, 255, 255, 255},
	"background11": {255, 255, 255, 255},
	"background12": {255, 255, 255, 255},
}

// NewNFT creates a new NFT with the provided parameters
func NewNFT(name, symbol, image string) *NFT {
	nameExt := strings.Replace(strings.ToLower(name), " ", "-", -1)
	nft := &NFT{
		Name:        name + " Staking",
		Symbol:      symbol,
		Description: name + " Staking NFT. Do not burn!",
		Image:       image,
		ExternalURL: "https://staking.solplutus.com/" + nameExt,
	}
	nft.Collection.Name = nft.Name
	nft.Collection.Family = nft.Name
	return nft
}
