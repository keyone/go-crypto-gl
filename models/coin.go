package models

import (
	"encoding/json"
	"time"
)

// Coin struct
type Coin struct {
	Name             string
	Symbol           string
	Rank             int64
	PriceUSD         float64
	PriceBTC         float64
	DayVolumeUSD     float64
	MarketCapUSD     float64
	AvailableSupply  float64
	TotalSupply      float64
	MaxSupply        float64
	PercentChange1H  float64
	PercentChange24H float64
	PercentChange7D  float64
	LastUpdated      time.Time
}

// JSONCoin struct
// This idea comes from:  https://blog.gopheracademy.com/advent-2016/advanced-encoding-decoding/
type JSONCoin struct {
	Name             string  `json:"name"`
	Symbol           string  `json:"symbol"`
	Rank             int64   `json:"rank,string"`
	PriceUSD         float64 `json:"price_usd,string"`
	PriceBTC         float64 `json:"price_btc,string"`
	DayVolumeUSD     float64 `json:"24h_volume_usd,string"`
	MarketCapUSD     float64 `json:"market_cap_usd,string"`
	AvailableSupply  float64 `json:"total_market_cap_usd,string"`
	TotalSupply      float64 `json:"available_supply,string"`
	MaxSupply        float64 `json:"max_supply,string"`
	PercentChange1H  float64 `json:"percent_change_1h,string"`
	PercentChange24H float64 `json:"percent_change_24h,string"`
	PercentChange7D  float64 `json:"percent_change_7d,string"`
	LastUpdated      int64   `json:"last_updated,string"`
}

// Coin method that translates a JSONCoin into a Coin struct
func (jc JSONCoin) Coin() Coin {
	return Coin{
		jc.Name,
		jc.Symbol,
		jc.Rank,
		jc.PriceUSD,
		jc.PriceBTC,
		jc.DayVolumeUSD,
		jc.MarketCapUSD,
		jc.AvailableSupply,
		jc.TotalSupply,
		jc.MaxSupply,
		jc.PercentChange1H,
		jc.PercentChange24H,
		jc.PercentChange7D,
		time.Unix(jc.LastUpdated, 0),
	}
}

// UnmarshalJSON method implements Unmarshaler interface
func (c *Coin) UnmarshalJSON(b []byte) error {
	var jc JSONCoin
	if err := json.Unmarshal(b, &jc); err != nil {
		return err
	}
	*c = jc.Coin()
	return nil
}
