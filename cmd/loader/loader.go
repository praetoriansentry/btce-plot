package main

import (
	"flag"
	"log"
	"praetoriansentry/btce-plot/pkg/archive"
	"praetoriansentry/btce-plot/pkg/btce"
)

var tradeType string
var out string

func main() {
	flag.Parse()
	trades, err := btce.GetTrades(2000, tradeType)
	if err != nil {
		log.Fatal(err)
	}
	archive.WriteToFile(trades, out)
}

func init() {
	flag.StringVar(&tradeType, "type", "eth_usd", "Trade type")
	flag.StringVar(&out, "o", "data.json", "Output Filename")
}
