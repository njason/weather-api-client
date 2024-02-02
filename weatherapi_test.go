package weatherapi

import (
	"log"
	"os"
	"testing"
	"time"

	"gopkg.in/yaml.v2"
)

type TestConfig struct {
	WeatherApiKey string `yaml:"weatherApiKey"`
}

func loadTestConfig() TestConfig {
	f, err := os.Open("config.yaml")
	if err != nil {
		log.Fatalln(err.Error())
	}
	defer f.Close()

	var config TestConfig
	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&config)
	if err != nil {
		log.Fatalln(err.Error())
	}

	return config
}
var testConfig = loadTestConfig()

func TestDoHistoryRequest(t *testing.T) {
	request := NewHistoryRequest(40.7790, -73.9692, time.Now().UTC())
	
	_, err := DoHistoryRequest(testConfig.WeatherApiKey, request)

	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
}
