package fmfw_test

import (
	"testing"

	"github.com/CIDgravity/Ticker/internal/exchange/fmfw"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	internalTesting "github.com/CIDgravity/Ticker/pkg/testing"
)

func TestFetch(t *testing.T) {
	pair := "FIL_USD"
	mockResponse := `
	{
		"FILUSDT": {
			"ask": "2.4128",
			"bid": "2.4087",
			"last": "2.4126",
			"low": "2.2796",
			"high": "2.5321",
			"open": "2.3252",
			"volume": "262379.462",
			"volume_quote": "632768.1805284",
			"timestamp": "2025-04-10T08:15:20.314Z"
		}
	}`

	mockServer := internalTesting.NewMockExchange(t, mockResponse)
	defer mockServer.Close()

	fmfwExchange := fmfw.New()
	fmfwExchange.SetBaseUrl(mockServer.URL)

	// Mock the pair resolution
	resp, err := fmfwExchange.Fetch(pair)
	require.NoError(t, err)
	assert.Equal(t, 2.4126, resp.Price)
	assert.Equal(t, 262379.462, resp.Volume)
}
