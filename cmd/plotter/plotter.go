package main

import (
	"fmt"
	"praetoriansentry/btce-plot/pkg/btce"
	"flag"
)

var limit int
var tradeType string

func main() {
	flag.Parse()
	btce.GetTrades(limit, tradeType)
	fmt.Println("Beep")
}

func init() {
	flag.IntVar(&limit, "limit", 150, "Trade Limit")
	flag.StringVar(&tradeType, "type", "eth_usd", "Trade type")
	
}
