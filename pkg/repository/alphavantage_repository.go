package repository

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sort"
	"strconv"
)

type AlphaVantageRepository struct {
	APIKey string
}

func (r *AlphaVantageRepository) GetStockPrice(symbol string) (float64, string, error) {
	url := fmt.Sprintf("https://www.alphavantage.co/query?function=TIME_SERIES_INTRADAY&symbol=%s&interval=1min&apikey=%s", symbol, r.APIKey)

	resp, err := http.Get(url)
	if err != nil {
		return 0, "", fmt.Errorf("failed to fetch stock price: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return 0, "", fmt.Errorf("failed to fetch stock price: %s", resp.Status)
	}

	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return 0, "", err
	}

	timeseries, ok := result["Time Series (1min)"].(map[string]interface{})
	if !ok {
		return 0, "", fmt.Errorf("error extracting time series data")
	}

	var timestamps []string
	for ts := range timeseries {
		timestamps = append(timestamps, ts)
	}
	sort.Sort(sort.Reverse(sort.StringSlice(timestamps)))

	latestTime := timestamps[0]
	latestData := timeseries[latestTime].(map[string]interface{})
	price := latestData["4. close"].(string)

	priceFloat, err := strconv.ParseFloat(price, 64)
	if err != nil {
		return 0, "", err
	}

	return priceFloat, latestTime, nil
}
