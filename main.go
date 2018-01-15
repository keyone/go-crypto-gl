package main

import (
	"fmt"
	"time"

	"github.com/keyone/go-crypto-gl/api"
)

func main() {
	// Timer to evaluate how fast the application is
	start := time.Now()

	// We are getting the total number of active currencies from the Global endpoint.
	activeCurrencies := api.GetActiveCurrencies()

	// TODO: We get an array of all currencies with their percent changes
	// from this endpoing: https://api.coinmarketcap.com/v1/ticker/?limit=902
	coins := api.GetAllCoins(activeCurrencies)
	fmt.Printf("%#v\n", coins)

	// elapsed timer
	elapsed := time.Since(start)
	fmt.Printf("\nExecution time: %s\n", elapsed)
}
