package fixerClient;

import (
	"fmt"
	"net/http"
	"time"
	"io/ioutil"
	"encoding/json"
)

var baseURL = "https://api.fixer.io"

// FixerRates represents a set of exchange rates, i.g. "USD" -> 1.2222
type FixerRates map[string]float32

// FixerResponse represents fixer API response structure
type FixerResponse struct {
	Date string `json:"date"`
	Base string `json:"base"`
	Rates FixerRates `json:"rates"`
}

var client = &http.Client{
	Timeout: time.Second * 10,
}

func requestAndParse(url string, base string) (FixerResponse, error) {
	resp, reqErr := client.Get(fmt.Sprintf("%s/%s?base=%s", baseURL, url, base))

	defer resp.Body.Close()

	if reqErr != nil {
		// TODO: append error here
		return FixerResponse{}, reqErr
	}

	respBody, bodyErr := ioutil.ReadAll(resp.Body)

	if bodyErr != nil {
		return FixerResponse{}, bodyErr
	}

	var parsedBody FixerResponse

	parseErr := json.Unmarshal(respBody, &parsedBody)

	if parseErr != nil {
		return FixerResponse{}, parseErr
	}

	return parsedBody, nil
}

// Latest returns fixer API response for the given base
// and for the latest date available
func Latest(base string) (FixerResponse, error) {
	return requestAndParse("/latest", base)
}

// ForDate returns fixer API response for the given base and
// for the specified date
func ForDate(date string, base string) (FixerResponse, error) {
	return requestAndParse(fmt.Sprintf("/%s", date), base)
}
