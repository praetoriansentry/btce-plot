package btce

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"praetoriansentry/btce-plot/pkg/analysis"
	"praetoriansentry/btce-plot/pkg/data"
)

func GetTrades(limit int, tradeType string) ([]data.Indicator, error) {
	url := fmt.Sprintf("https://btc-e.com/api/3/trades/%s?limit=%d", tradeType, limit)
	log.Printf("Fetching data from BTC-E url: %s", url)
	resp, err := http.Get(url)
	if err != nil {
		log.Print("There was an issue connecting to btce")
		log.Print(err)
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Print("There was an issue reading response body")
		log.Print(err)
		return nil, err
	}

	var responseData data.TradeResponse
	err = json.Unmarshal(body, &responseData)

	if err != nil {
		log.Print("There was an issue reading the json data")
		log.Print(err)
		return nil, err
	}

	ts, ok := responseData[tradeType]
	if !ok {
		log.Print("Data didn't contain a valid trade type")
		return nil, errors.New("Mismatched trade type")
	}
	indicators := analysis.BuildIndicators(ts, 5*60)
	return indicators, nil

}
