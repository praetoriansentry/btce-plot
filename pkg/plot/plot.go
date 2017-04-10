package plot

import (
	"log"
	"io/ioutil"
	"fmt"
	"praetoriansentry/btce-plot/pkg/data"
)

func CreatePlot(indicators []data.Indicator, x, y int) {
	log.Print("Creating plot")
	log.Print(createDatFile(indicators))
}

func createDatFile(indicators []data.Indicator) string{
	tmpfile, err := ioutil.TempFile("", "dat")
	if err != nil {
		log.Fatal(err)
	}
	for _ , i:= range indicators {
		fmt.Fprintf(tmpfile, "%s %f %f %f %f %f\n", i.Date, i.Open, i.High, i.Low, i.Close, i.Volume)
	}
	tmpfile.Close()
	return tmpfile.Name()
}
