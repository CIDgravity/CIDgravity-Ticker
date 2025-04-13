package bitfinex_test

import (
	"net/http"
	"testing"

	"github.com/CIDgravity/Ticker/internal/exchange/bitfinex"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	internalTesting "github.com/CIDgravity/Ticker/pkg/testing"
)

// TestFetchAndUnifiedResponse test the fetch and the conversion to unified response at the same time
func TestFetchAndUnifiedResponse(t *testing.T) {

	t.Run("invalid_response", func(t *testing.T) {
		pair := "FIL_USD"
		mockResponse := `
		[
			2.4001,
			19325.25889251,
			2.4034,
			12939.53723539,
			0.0751,
			0.03216412,
			2.41
		]`

		mockServer := internalTesting.NewMockExchange(t, mockResponse, http.StatusOK)
		defer mockServer.Close()

		bitfinexExchange := bitfinex.New()
		bitfinexExchange.SetBaseUrl(mockServer.URL)

		resp, err := bitfinexExchange.Fetch(pair)
		require.Error(t, err)
		assert.Empty(t, resp)
		assert.ErrorContains(t, err, "invalid response from Bitfinex: expect 10 values")
	})

	t.Run("bad_request", func(t *testing.T) {
		pair := "FIL_USD"
		mockResponse := `[]`

		mockServer := internalTesting.NewMockExchange(t, mockResponse, http.StatusBadRequest)
		defer mockServer.Close()

		bitfinexExchange := bitfinex.New()
		bitfinexExchange.SetBaseUrl(mockServer.URL)

		resp, err := bitfinexExchange.Fetch(pair)
		require.Error(t, err)
		assert.Empty(t, resp)
		assert.ErrorContains(t, err, "invalid HTTP status code")
	})

	t.Run("success", func(t *testing.T) {
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

		mockServer := internalTesting.NewMockExchange(t, mockResponse, http.StatusOK)
		defer mockServer.Close()

		bitfinexExchange := bitfinex.New()
		bitfinexExchange.SetBaseUrl(mockServer.URL)

		resp, err := bitfinexExchange.Fetch(pair)
		require.NoError(t, err)
		assert.Equal(t, 2.40175, resp.Price)
		assert.Equal(t, 16132.39806395, resp.Volume)
	})
}