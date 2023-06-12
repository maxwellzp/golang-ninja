package coincap

import "fmt"

type Asset struct {
	ID           string `json:"id"`
	Rank         string `json:"rank"`
	Symbol       string `json:"symbol"`
	Name         string `json:"name"`
	Supply       string `json:"supply"`
	MaxSupply    string `json:"maxSupply"`
	MarketCapUSD string `json:"marketCapUsd"`
	VolumeUsd24h string `json:"volumeUsd24hr"`
	PriceUSD     string `json:"priceUsd"`
}

type assetsResponse struct {
	Assets    []Asset `json:"data"`
	Timestamp int64   `json:"timestamp"`
}

type assetResponse struct {
	Asset     Asset `json:"data"`
	Timestamp int64 `json:"timestamp"`
}

func (d Asset) Info() string {
	return fmt.Sprintf("[ID] %s | [RANK] %s | [SYMBOL] %s | [NAME] %s | [PRICE] %s",
		d.ID, d.Rank, d.Symbol, d.Name, d.PriceUSD)
}
