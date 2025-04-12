package kraken

import (
	"fmt"
	"strconv"

	"github.com/CIDgravity/Ticker/config"
	"github.com/CIDgravity/Ticker/internal/exchange"
	"github.com/CIDgravity/Ticker/pkg/http"
)

type Kraken struct {
	config exchange.ExchangeConfig
}

type KrakenResponse struct {
	Error  []string                     `json:"error"`
	Result map[string]KrakenPairDetails `json:"result"`
}

// KrakenPairDetails represents the details of a trading pair (e.g., FILUSD)
type KrakenPairDetails struct {
	Ask                        []string `json:"a"` // Ask [<price>, <whole lot volume>, <lot volume>]
	Bid                        []string `json:"b"` // Bid [<price>, <whole lot volume>, <lot volume>]
	LastTradeClosed            []string `json:"c"` // Last trade closed [<price>, <lot volume>]
	Volume                     []string `json:"v"` // Volume [<today>, <last 24 hours>]
	VolumeWeightedAveragePrice []string `json:"p"` // Volume weighted average price [<today>, <last 24 hours>]
	NumberOfTrades             []int    `json:"t"` // Number of trades [<today>, <last 24 hours>]
	Low                        []string `json:"l"` // Low [<today>, <last 24 hours>]
	High                       []string `json:"h"` // High [<today>, <last 24 hours>]
	TodaysOpeningPrice         string   `json:"o"` // Today's opening price
}

func (r KrakenResponse) ToUnifiedResponse(exchangeName string, configPair, exchangePair string) (exchange.ExchangeFetchResponseForPair, error) {
	result, ok := r.Result[exchangePair]

	if !ok {
		return exchange.ExchangeFetchResponseForPair{}, fmt.Errorf("pair not found in response")
	}

	// parse string values to float for calculation
	askPrice, err := strconv.ParseFloat(result.Ask[0], 64)
	if err != nil {
		return exchange.ExchangeFetchResponseForPair{}, err
	}

	bidPrice, err := strconv.ParseFloat(result.Bid[0], 64)
	if err != nil {
		return exchange.ExchangeFetchResponseForPair{}, err
	}

	askVolume, err := strconv.ParseFloat(result.Ask[2], 64)
	if err != nil {
		return exchange.ExchangeFetchResponseForPair{}, err
	}

	bidVolume, err := strconv.ParseFloat(result.Bid[2], 64)
	if err != nil {
		return exchange.ExchangeFetchResponseForPair{}, err
	}

	return exchange.ExchangeFetchResponseForPair{
		Price:  (bidPrice + askPrice) / 2,
		Volume: (bidVolume + askVolume) / 2,
	}, nil
}

func New() *Kraken {
	// Parse each config pair to the pairs related to the platform

	return &Kraken{
		config: exchange.ExchangeConfig{
			Name:     "Kraken",
			Endpoint: "https://api.kraken.com/0/public/Ticker",
			Timeout:  "15s",
		},
	}
}

// SetEndpoint update the endpoint (used for testing purposes)
// Must not contains the ending slash
func (x *Kraken) SetBaseUrl(baseUrl string) {
	x.config.Endpoint = baseUrl + "/0/public/Ticker"
}

// GetName return exchange name
func (x *Kraken) GetName() string {
	return x.config.Name
}

// Fetch return response for current pair with an unified response format
func (x *Kraken) Fetch(pair string) (exchange.ExchangeFetchResponseForPair, error) {
	exchangePair, err := config.GetPairForPlatform(x.config.Name, pair)
	if err != nil {
		return exchange.ExchangeFetchResponseForPair{}, err
	}

	url := fmt.Sprintf("%s?pair=%s", x.config.Endpoint, exchangePair)
	resp, err := http.ExecuteRequest[KrakenResponse](url, x.config.Timeout)
	if err != nil {
		return exchange.ExchangeFetchResponseForPair{}, err
	}

	return resp.ToUnifiedResponse(x.GetName(), pair, exchangePair)
}
