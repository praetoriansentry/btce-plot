package plot

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"praetoriansentry/btce-plot/pkg/data"
	"text/template"
)

func CreatePlot(indicators []data.Indicator, outputName string, x, y, bucketSeconds int) {
	log.Print("Creating plot")
	datName := createDatFile(indicators)
	t := getTemplate()
	xMin, xMax := getRange(indicators)
	pt := writeTemplateData(t, datName, outputName, x, y, xMin, xMax, bucketSeconds)
	createPng(pt)

	os.Remove(datName)
	os.Remove(pt)
}

func getRange(indicators []data.Indicator) (string, string) {
	l := len(indicators)
	return indicators[0].Date, indicators[l-1].Date
}

func createDatFile(indicators []data.Indicator) string {
	tmpfile, err := ioutil.TempFile("", "dat")
	if err != nil {
		log.Fatal(err)
	}
	for _, i := range indicators {
		fmt.Fprintf(tmpfile, "%s %f %f %f %f %f\r\n", i.Date, i.Open, i.Low, i.High, i.Close, i.Volume)
	}
	tmpfile.Close()
	return tmpfile.Name()
}

func getTemplate() *template.Template {
	data, err := ioutil.ReadFile("plot.gnu")
	if err != nil {
		log.Fatal(err)
	}
	plotTemplate := template.Must(template.New("gnuplot").Parse(string(data)))
	return plotTemplate
}

func writeTemplateData(t *template.Template, fileName, outputName string, x, y int, xMin, xMax string, bucketSeconds int) string {
	templateData := struct {
		DatName string
		OutName string
		X       int
		Y       int
		XMin    string
		XMax    string
		BucketSeconds int
	}{
		fileName,
		outputName,
		x,
		y,
		xMin,
		xMax,
		bucketSeconds * 50 / 100,
	}

	gnuTemplate, err := ioutil.TempFile("", "gnutemplate")
	if err != nil {
		log.Fatal(err)
	}

	defer gnuTemplate.Close()
	err = t.Execute(gnuTemplate, templateData)

	if err != nil {
		log.Fatal(err)
	}
	return gnuTemplate.Name()

}

func createPng(templateName string) {
	c := exec.Command("gnuplot", templateName)

	stdout, err := c.StdoutPipe()
	if err != nil {
		log.Fatal(err)
	}
	stderr, err := c.StderrPipe()
	if err != nil {
		log.Fatal(err)
	}

	err = c.Start()
	if err != nil {
		log.Fatal(err)
	}

	_, err = ioutil.ReadAll(stdout)
	if err != nil {
		log.Fatal(err)
	}

	_, err = ioutil.ReadAll(stderr)
	if err != nil {
		log.Fatal(err)
	}

	err = c.Wait()
	if err != nil {
		log.Fatal(err)
	}

}
