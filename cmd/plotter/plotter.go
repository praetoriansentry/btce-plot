package main

import (
	"log"
	"praetoriansentry/btce-plot/pkg/btce"
	"flag"
)

var limit int
var tradeType string

func main() {
	flag.Parse()
	indicators, err := btce.GetTrades(limit, tradeType)
	if err != nil {
		log.Fatal(err)
	}
	log.Print(indicators)
}

func init() {
	flag.IntVar(&limit, "limit", 150, "Trade Limit")
	flag.StringVar(&tradeType, "type", "eth_usd", "Trade type")
	
}
