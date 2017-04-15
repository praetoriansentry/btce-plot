package main

import (
	"flag"
	"log"
	"praetoriansentry/btce-plot/pkg/btce"
	"praetoriansentry/btce-plot/pkg/plot"
)

var limit int
var x int
var y int
var tradeType string
var out string
var bucketSize int
var terminal string

func main() {
	flag.Parse()
	indicators, err := btce.GetIndicators(limit, tradeType, bucketSize)
	if err != nil {
		log.Fatal(err)
	}
	plot.CreatePlot(indicators, out, terminal, x, y, bucketSize)
	log.Print("Plot created")
}

func init() {
	flag.IntVar(&limit, "limit", 150, "Trade Limit")
	flag.IntVar(&bucketSize, "bucket", 60, "The number of seconds that we'll use for bucketing")
	flag.IntVar(&x, "x", 1920, "X Dimension in Px")
	flag.IntVar(&y, "y", 1080, "Y Dimension in Px")
	flag.StringVar(&tradeType, "type", "eth_usd", "Trade type")
	flag.StringVar(&out, "o", "graph.png", "Output Filename")
	flag.StringVar(&terminal, "t", "pngcairo", "The gnuplot terminal that we'll use")

}
