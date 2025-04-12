package kraken_test

import (
	"testing"

	"github.com/CIDgravity/Ticker/internal/exchange/kraken"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	internalTesting "github.com/CIDgravity/Ticker/pkg/testing"
)

func TestFetch(t *testing.T) {
	pair := "FIL_USD"
	mockResponse := `
	{
		"error": [],
		"result": {
			"FILUSD": {
				"a": ["2.40600","1693","1693.000"],
				"b": ["2.40400","2976","2976.000"],
				"c": ["2.40900","120.13494810"],
				"v": ["20266.86995345","191581.17323023"],
				"p": ["2.42500","2.46280"],
				"t": [119, 953],
				"l": ["2.38800", "2.28000"],
				"h": ["2.47300","2.53100"],
				"o": "2.47300"
			}
		}
	}`

	mockServer := internalTesting.NewMockExchange(t, mockResponse)
	defer mockServer.Close()

	krakenExchange := kraken.New()
	krakenExchange.SetBaseUrl(mockServer.URL)

	// Mock the pair resolution
	resp, err := krakenExchange.Fetch(pair)
	require.NoError(t, err)
	assert.Equal(t, 2.42500, resp.Price)
	assert.Equal(t, 20266.86995345, resp.Volume)
}
