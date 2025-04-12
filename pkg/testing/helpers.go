package testing

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"runtime"
	"testing"
)

func NewMockExchange(t *testing.T, jsonResponse string) *httptest.Server {
	t.Helper()

	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, jsonResponse)
	}))
}

func GetTestPathForData(t *testing.T, folder string) string {
	_, testFile, _, ok := runtime.Caller(1)
	if !ok {
		t.Fatal("Could not determine caller location")
	}

	baseDir := filepath.Dir(testFile)
	return filepath.Join(baseDir, folder)
}
