package btce

import (
	"net/http"
	"log"
	"fmt"
	"encoding/json"
	"io/ioutil"
	"praetoriansentry/btce-plot/pkg/data"
)

func GetTrades(limit int, tradeType string) {
	url := fmt.Sprintf("https://btc-e.com/api/3/trades/%s?limit=%d", tradeType, limit)
	log.Printf("Fetching data from BTC-E url: %s", url)
	resp, err := http.Get(url)
	if err != nil {
		log.Print("There was an issue connecting to btce")
		log.Print(err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Print("There was an issue reading response body")
		log.Print(err)
		return
	}

	var responseData data.TradeResponse
	err = json.Unmarshal(body, &responseData)
	
	if err != nil {
		log.Print("There was an issue reading the json data")
		log.Print(err)
		return
	}
	

	log.Print(responseData)
	
}
