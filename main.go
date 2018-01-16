package main

import (
	"fmt"
	"time"

	"github.com/keyone/go-crypto-gl/coin"

	"github.com/keyone/go-crypto-gl/api"
)

func main() {
	// Timer to evaluate how fast the application is
	start := time.Now()
	fmt.Println("*****************************************")
	fmt.Println("*** Top 10 Gainers in the past 1 hour ***")
	fmt.Println("*****************************************")

	fmt.Printf("%s\t%s\t%s\t\t%s\n", "Rank", "Symbol", "Price", "PercentChange1H")

	for _, coin := range coin.ByDecreasingPercentChange1H(api.GetAllCoins(1500)) {
		fmt.Printf("%d\t%s\t$%.5f\t%.2f%%\n", coin.Rank, coin.Symbol, coin.PriceUSD, coin.PercentChange1H)
	}

	// elapsed timer
	elapsed := time.Since(start)
	fmt.Printf("\nExecution time: %s\n", elapsed)
}
