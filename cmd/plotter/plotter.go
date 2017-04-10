package main

import (
	"log"
	"praetoriansentry/btce-plot/pkg/btce"
	"praetoriansentry/btce-plot/pkg/plot"
	"flag"
)

var limit int
var x int
var y int
var tradeType string
var out string

func main() {
	flag.Parse()
	indicators, err := btce.GetTrades(limit, tradeType)
	if err != nil {
		log.Fatal(err)
	}
	plot.CreatePlot(indicators, out, x, y)
	log.Print("Plot created")
}

func init() {
	flag.IntVar(&limit, "limit", 150, "Trade Limit")
	flag.IntVar(&x, "x", 1920, "X Dimension in Px")
	flag.IntVar(&y, "y", 1080, "Y Dimension in Px")
	flag.StringVar(&tradeType, "type", "eth_usd", "Trade type")
	flag.StringVar(&out, "o", "graph.png", "Output Filename")
	
}
