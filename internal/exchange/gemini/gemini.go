package gemini

import (
	"fmt"
	"strconv"

	"github.com/CIDgravity/Ticker/config"
	"github.com/CIDgravity/Ticker/internal/exchange"
	"github.com/CIDgravity/Ticker/pkg/http"
)

type Gemini struct {
	config exchange.ExchangeConfig
}

// MarketData represents the JSON structure
type GeminiResponse struct {
	Symbol  string   `json:"symbol"`
	Open    string   `json:"open"`
	High    string   `json:"high"`
	Low     string   `json:"low"`
	Close   string   `json:"close"`
	Changes []string `json:"changes"`
	Bid     string   `json:"bid"`
	Ask     string   `json:"ask"`
}

func (r GeminiResponse) ToUnifiedResponse(exchangeName string, configPair string) (exchange.ExchangeFetchResponseForPair, error) {
	bidPrice, err := strconv.ParseFloat(r.Bid, 32)
	if err != nil {
		return exchange.ExchangeFetchResponseForPair{}, err
	}

	askPrice, err := strconv.ParseFloat(r.Ask, 32)
	if err != nil {
		return exchange.ExchangeFetchResponseForPair{}, err
	}

	return exchange.ExchangeFetchResponseForPair{
		Price:  (bidPrice + askPrice) / 2,
		Volume: 0.0,
	}, nil
}

func New() *Gemini {
	return &Gemini{
		config: exchange.ExchangeConfig{
			Name:     "Gemini",
			Endpoint: "https://api.gemini.com/v2/ticker",
			Timeout:  "15s",
		},
	}
}

// SetEndpoint update the endpoint (used for testing purposes)
// Must not contains the ending slash
func (x *Gemini) SetBaseUrl(baseUrl string) {
	x.config.Endpoint = baseUrl + "/v2/ticker"
}

// GetName return exchange name
func (x *Gemini) GetName() string {
	return x.config.Name
}

func (x *Gemini) Fetch(pair string) (exchange.ExchangeFetchResponseForPair, error) {
	exchangePair, err := config.GetPairForPlatform(x.config.Name, pair)
	if err != nil {
		return exchange.ExchangeFetchResponseForPair{}, err
	}

	url := fmt.Sprintf("%s/%s", x.config.Endpoint, exchangePair)
	resp, err := http.ExecuteRequest[GeminiResponse](url, x.config.Timeout)

	if err != nil {
		return exchange.ExchangeFetchResponseForPair{}, err
	}

	return resp.ToUnifiedResponse(x.GetName(), pair)
}
