package api

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/keyone/go-crypto-gl/models"
)

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
	g := new(models.Globals)
	if err := json.Unmarshal(body, &g); err != nil {
		log.Fatal("Cannot unmarshal JSON")
	}
	return g.ActiveCurrencies
}

// GetAllCoins total
func GetAllCoins(activeCurrencies int64) (coins []models.Coin) {
	resp, err := http.Get("https://api.coinmarketcap.com/v1/ticker/?limit=" + strconv.FormatInt(activeCurrencies, 10))
	if err != nil {
		log.Fatal("Cannot access Api")
	}
	defer resp.Body.Close()
	coinsBody, ioerr := ioutil.ReadAll(resp.Body)
	if ioerr != nil {
		log.Fatal("Cannot read response")
	}
	if err := json.Unmarshal(coinsBody, &coins); err != nil {
		log.Fatal("Cannot unmarshal Coins JSON")
	}
	return
}
