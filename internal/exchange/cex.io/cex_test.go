package cex_test

import (
	"testing"

	cex "github.com/CIDgravity/Ticker/internal/exchange/cex.io"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	internalTesting "github.com/CIDgravity/Ticker/pkg/testing"
)

func TestFetch(t *testing.T) {
	pair := "FIL_USD"
	mockResponse := `
	{
		"ok":"ok",
		"data":{
			"FIL-USD":{
				"bestBid":"2.408",
				"bestAsk":"2.409",
				"bestBidChange":"0.083",
				"bestBidChangePercentage":"3.56",
				"bestAskChange":"0.086",
				"bestAskChangePercentage":"3.70",
				"low":"2.479",
				"high":"2.479",
				"volume30d":"20607.80841568",
				"lastTradeDateISO":"2025-04-09T23:31:38.437Z",
				"volume":"10.23750000",
				"quoteVolume":"25.38388125",
				"lastTradeVolume":"10.23750000",
				"volumeUSD":"25.38",
				"last":"2.409",
				"lastTradePrice":"2.479",
				"priceChange":"0.086",
				"priceChangePercentage":"3.70"
			}
		}
	}`

	mockServer := internalTesting.NewMockExchange(t, mockResponse)
	defer mockServer.Close()

	cexExchange := cex.New()
	cexExchange.SetBaseUrl(mockServer.URL)

	// Mock the pair resolution
	resp, err := cexExchange.Fetch(pair)
	require.NoError(t, err)
	assert.Equal(t, 2.479, resp.Price)
	assert.Equal(t, 10.2375, resp.Volume)
}
