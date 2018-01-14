package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {

	changelap := flag.String("changelap", "1h", "The change lap: It could be 1h, 24h or 7d")
	top := flag.String("top", "5", "total number of gainers to consider")

	fmt.Println("Welcome to Top Gainers App")

	// We are getting the total number of active currencies from the Global endpoint.
	activeCurrencies := GetActiveCurrencies()
	fmt.Println(activeCurrencies)

	// TODO: We get an array of all currencies with their percent changes
	// from this endpoing: https://api.coinmarketcap.com/v1/ticker/?limit=902
	ShowTopGainers(*changelap, *top)

}

// Globals struct is mapping the /global endpoint
type Globals struct {
	TotalMarketCapUSD      float64 `json:"total_market_cap_usd"`
	Total24VolumeUSD       float64 `json:"total_24h_volume_usd"`
	BTCPercentageMarketCap float64 `json:"bitcoin_percentage_of_market_cap"`
	ActiveCurrencies       int64   `json:"active_currencies"`
	ActiveAssets           int64   `json:"active_assets"`
	ActiveMarkets          int64   `json:"active_markets"`
	LastUpdated            int64   `json:"last_updated"`
}

// Coin struct
type Coin struct {
	Name             string  `json:"name"`
	Symbol           string  `json:"symbol"`
	Rank             int64   `json:"rank"`
	PriceUSD         float64 `json:"price_usd"`
	PriceBTC         float64 `json:"price_btc"`
	DayVolumeUSD     float64 `json:"24h_volume_usd"`
	MarketCapUSD     float64 `json:"market_cap_usd"`
	AvailableSupply  float64 `json:"total_market_cap_usd"`
	TotalSupply      float64 `json:"available_supply"`
	MaxSupply        float64 `json:"max_supply"`
	PercentChange1H  float64 `json:"percent_change_1h"`
	PercentChange24H float64 `json:"percent_change_24h"`
	PercentChange7D  float64 `json:"percent_change_7d"`
	LastUpdated      int64   `json:"last_updated"`
}

// GetActiveCurrencies function returns total nomber of active currencies
func GetActiveCurrencies() int64 {
	resp, err := http.Get("https://api.coinmarketcap.com/v1/global/")
	if err != nil {
		log.Fatal("Cannot access API")
	}
	defer resp.Body.Close()

	body, ioerr := ioutil.ReadAll(resp.Body)
	if ioerr != nil {
		log.Fatal("Cannot read response")
	}

	var g = new(Globals)
	if err := json.Unmarshal(body, &g); err != nil {
		log.Fatal("Cannot unmarshal JSON")
	}
	return g.ActiveCurrencies
}

// ShowTopGainers total
func ShowTopGainers(changelap string, top string) {

}
