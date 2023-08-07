/**
 * https://www.weatherapi.com/docs/
 */

package weatherapi

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

var client = http.Client{
	Timeout: 60 * time.Second,
}

func NewHistoryRequest(lat float64, lng float64, time time.Time) historyRequest {
		return historyRequest{
			q: fmt.Sprintf("%f, %f", lat, lng),
			dt: time.Format("2006-01-02"),
		}
}

func DoHistoryRequest(apiKey string, request historyRequest) (weatherApiResponse, error) {
	req, err := http.NewRequest("GET", "https://api.weatherapi.com/v1/history.json", nil)
	if err != nil {
		return weatherApiResponse{}, err
	}

	query := req.URL.Query()
	query.Add("q", request.q)
	query.Add("dt", request.dt)
	query.Add("key", apiKey)

	req.URL.RawQuery = query.Encode()

	rawResp, err := client.Do(req)
	if err != nil {
		return weatherApiResponse{}, err
	}
	defer rawResp.Body.Close()

	var resp weatherApiResponse
	err = json.NewDecoder(rawResp.Body).Decode(&resp)
	if err != nil {
		return weatherApiResponse{}, err
	}

	return resp, nil
}
