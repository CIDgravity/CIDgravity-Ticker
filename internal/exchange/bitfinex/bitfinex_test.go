package bitfinex_test

import (
	"testing"

	"github.com/CIDgravity/Ticker/internal/exchange/bitfinex"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	internalTesting "github.com/CIDgravity/Ticker/pkg/testing"
)

func TestFetch(t *testing.T) {
	pair := "FIL_USD"
	mockResponse := `
	[
		2.4001,
		19325.25889251,
		2.4034,
		12939.53723539,
		0.0751,
		0.03216412,
		2.41,
		35133.36313284,
		2.5301,
		2.2901
	]`

	mockServer := internalTesting.NewMockExchange(t, mockResponse)
	defer mockServer.Close()

	bitfinexExchange := bitfinex.New()
	bitfinexExchange.SetBaseUrl(mockServer.URL)

	resp, err := bitfinexExchange.Fetch(pair)
	require.NoError(t, err)
	assert.Equal(t, 2.41, resp.Price)
	assert.Equal(t, 35133.36313284, resp.Volume)
}
