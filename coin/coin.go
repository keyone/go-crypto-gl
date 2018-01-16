package coin

import (
	"encoding/json"
	"sort"
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

// multiSorter implements the Sort interface, sorting the changes within.
// code from: https://golang.org/pkg/sort/
type multiSorter struct {
	coins []Coin
	less  []lessFunc
}

// Len is part of sort.Interface.
func (ms *multiSorter) Len() int {
	return len(ms.coins)
}

// Sort sorts the argument slice according to the less functions passed to OrderedBy.
func (ms *multiSorter) Sort(coins []Coin) {
	ms.coins = coins
	sort.Sort(ms)
}

// Swap is part of sort.Interface.
func (ms *multiSorter) Swap(i, j int) {
	ms.coins[i], ms.coins[j] = ms.coins[j], ms.coins[i]
}

// Less is part of sort.Interface. It is implemented by looping along the
// less functions until it finds a comparison that is either Less or
// !Less. Note that it can call the less functions twice per call. We
// could change the functions to return -1, 0, 1 and reduce the
// number of calls for greater efficiency: an exercise for the reader.
func (ms *multiSorter) Less(i, j int) bool {
	p, q := &ms.coins[i], &ms.coins[j]
	// Try all but the last comparison.
	var k int
	for k = 0; k < len(ms.less)-1; k++ {
		less := ms.less[k]
		switch {
		case less(p, q):
			// p < q, so we have a decision.
			return true
		case less(q, p):
			// p > q, so we have a decision.
			return false
		}
		// p == q; try the next comparison.
	}
	// All comparisons to here said "equal", so just return whatever
	// the final comparison reports.
	return ms.less[k](p, q)
}

type lessFunc func(p1, p2 *Coin) bool

// OrderedBy returns a Sorter that sorts using the less functions, in order.
// Call its Sort method to sort the data.
func orderedBy(less ...lessFunc) *multiSorter {
	return &multiSorter{
		less: less,
	}
}

// ByDecreasingPercentChange1H sorts by percent change in 1 hour
func ByDecreasingPercentChange1H(coins []Coin) []Coin {
	decreasingPercentChange1H := func(c1, c2 *Coin) bool {
		return c1.PercentChange1H > c2.PercentChange1H // Note: > orders downwards.
	}
	orderedBy(decreasingPercentChange1H).Sort(coins)
	return coins[0:10]
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

// Globals struct is mapping the /global endpoint
type Globals struct {
	ActiveCurrencies int64 `json:"active_currencies"`
}
