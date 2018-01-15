package main

import (
	"fmt"
	"time"

	"github.com/keyone/go-crypto-gl/api"
)

func main() {
	// Timer to evaluate how fast the application is
	start := time.Now()

	coins := api.GetAllCoins(2000)
	fmt.Printf("%#v\n", coins)

	// elapsed timer
	elapsed := time.Since(start)
	fmt.Printf("\nExecution time: %s\n", elapsed)
}
