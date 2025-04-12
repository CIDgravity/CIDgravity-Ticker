package crypto_test

import (
	"testing"

	crypto "github.com/CIDgravity/Ticker/internal/exchange/crypto.com"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	internalTesting "github.com/CIDgravity/Ticker/pkg/testing"
)

func TestFetch(t *testing.T) {
	pair := "FIL_USD"
	mockResponse := `
	{
		"id" : -1,
		"method" : "public/get-tickers",
		"code" : 0,
		"result" : {
			"data" : [ 
				{
					"i" : "FIL_USD",
					"h" : "2.5363",
					"l" : "2.2837",
					"a" : "2.4142",
					"v" : "24190.15",
					"vv" : "58892.84",
					"c" : "0.0373",
					"b" : "2.4074",
					"k" : "2.4120",
					"oi" : "0",
					"t" : 1744272739898
				}
			]
		}
	}`

	mockServer := internalTesting.NewMockExchange(t, mockResponse)
	defer mockServer.Close()

	cryptoExchange := crypto.New()
	cryptoExchange.SetBaseUrl(mockServer.URL)

	// Mock the pair resolution
	resp, err := cryptoExchange.Fetch(pair)
	require.NoError(t, err)
	assert.Equal(t, 2.4142, resp.Price)
	assert.Equal(t, 24190.15, resp.Volume)
}
