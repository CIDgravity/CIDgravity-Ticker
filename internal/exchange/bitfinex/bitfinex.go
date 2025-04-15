package bitfinex

import (
	"fmt"

	"github.com/CIDgravity/Ticker/config"
	"github.com/CIDgravity/Ticker/internal/exchange"
	"github.com/CIDgravity/Ticker/pkg/http"
)

type Bitfinex struct {
	config exchange.ExchangeConfig
}

type BitfinexResponse struct {
	Bid                 float64 `json:"bid"`                   // Price of last highest bid
	BidSize             float64 `json:"bid_size"`              // Sum of the 25 highest bid sizes
	Ask                 float64 `json:"ask"`                   // Price of last lowest ask
	AskSize             float64 `json:"ask_size"`              // Sum of the 25 lowest ask sizes
	DailyChange         float64 `json:"daily_change"`          // Amount that the last price has changed since yesterday
	DailyChangeRelative float64 `json:"daily_change_relative"` // Relative price change since yesterday (*100 for percentage change)
	LastPrice           float64 `json:"last_price"`            // Price of the last trade
	Volume              float64 `json:"volume"`                // Daily volume
	High                float64 `json:"high"`                  // Daily high
	Low                 float64 `json:"low"`                   // Daily low
}

func (r BitfinexResponse) ToUnifiedResponse(exchangeName string, configPair string) exchange.ExchangeFetchResponseForPair {
	return exchange.ExchangeFetchResponseForPair{
		Price:  (r.Bid + r.Ask) / 2,
		Volume: (r.BidSize + r.AskSize) / 2,
	}
}

func New() *Bitfinex {
	return &Bitfinex{
		config: exchange.ExchangeConfig{
			Name:     "Bitfinex",
			Endpoint: "https://api-pub.bitfinex.com/v2/ticker",
			Timeout:  "15s",
		},
	}
}

// SetBaseURL update the endpoint (used for testing purposes)
// Must not contains the ending slash
func (x *Bitfinex) SetBaseURL(baseURL string) {
	x.config.Endpoint = baseURL + "/v2/ticker"
}

// GetName return exchange name
func (x *Bitfinex) GetName() string {
	return x.config.Name
}

// Fetch return a map with each pair with an unified response format
func (x *Bitfinex) Fetch(pair string) (exchange.ExchangeFetchResponseForPair, error) {
	exchangePair, err := config.GetPairForPlatform(x.config.Name, pair)
	if err != nil {
		return exchange.ExchangeFetchResponseForPair{}, err
	}

	url := fmt.Sprintf("%s/%s", x.config.Endpoint, exchangePair)
	resp, err := http.ExecuteRequest[[]float64](url, x.config.Timeout)
	if err != nil {
		return exchange.ExchangeFetchResponseForPair{}, err
	}

	// Bitfinex doesn't store in JSON struct, but only using array indexes
	// We need to remap to struct before converted to unified response
	if resp == nil || len(*resp) < 10 {
		return exchange.ExchangeFetchResponseForPair{}, fmt.Errorf("invalid response from Bitfinex: expect 10 values")
	}

	bitfinexResponse := *resp
	bitfinexParsedResponse := BitfinexResponse{
		Bid:                 bitfinexResponse[0],
		BidSize:             bitfinexResponse[1],
		Ask:                 bitfinexResponse[2],
		AskSize:             bitfinexResponse[3],
		DailyChange:         bitfinexResponse[4],
		DailyChangeRelative: bitfinexResponse[5],
		LastPrice:           bitfinexResponse[6],
		Volume:              bitfinexResponse[7],
		High:                bitfinexResponse[8],
		Low:                 bitfinexResponse[9],
	}

	return bitfinexParsedResponse.ToUnifiedResponse(x.GetName(), pair), nil
}
