package weatherapi

import (
	"log"
	"os"
	"strings"
	"testing"
	"time"
)

var apiKey = readApiKeyFile()
func readApiKeyFile() string {
	rawApiKey, err := os.ReadFile("apikey.txt")
	if err != nil {
		log.Fatal(err)
	}

	return strings.TrimSpace(string(rawApiKey))
}

func TestDoHistoryRequest(t *testing.T) {
	request := NewHistoryRequest(40.7790, -73.9692, time.Now().UTC())

	_, err := DoHistoryRequest(apiKey, request)

	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
}
