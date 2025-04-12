package gemini_test

import (
	"testing"

	"github.com/CIDgravity/Ticker/internal/exchange/gemini"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	internalTesting "github.com/CIDgravity/Ticker/pkg/testing"
)

func TestFetch(t *testing.T) {
	pair := "FIL_USD"
	mockResponse := `
	{
		"symbol": "FILUSD",
		"open": "21.5",
		"high": "23.0",
		"low": "20.5",
		"close": "22.0",
		"changes": ["0.5"],
		"bid": "21.9",
		"ask": "22.1"
	}`

	mockServer := internalTesting.NewMockExchange(t, mockResponse)
	defer mockServer.Close()

	geminiExchange := gemini.New()
	geminiExchange.SetBaseUrl(mockServer.URL)

	// Mock the pair resolution
	resp, err := geminiExchange.Fetch(pair)
	require.NoError(t, err)
	assert.Equal(t, 20.5, resp.Price)
	assert.Equal(t, 0.0, resp.Volume)
}
