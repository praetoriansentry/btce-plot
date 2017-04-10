package plot

import (
	"log"
	"io/ioutil"
	"fmt"
	"os/exec"
	"praetoriansentry/btce-plot/pkg/data"
	"text/template"
)

func CreatePlot(indicators []data.Indicator, outputName string, x, y int) {
	log.Print("Creating plot")
	datName := createDatFile(indicators)
	t := getTemplate()
	pt := writeTemplateData(t, datName, outputName, x, y)
	createPng(pt)
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

func getTemplate() *template.Template {
	data, err := ioutil.ReadFile("plot.gnu")
	if err != nil {
		log.Fatal(err)
	}
	plotTemplate := template.Must(template.New("gnuplot").Parse(string(data)))
	return plotTemplate
}

func writeTemplateData(t *template.Template, fileName, outputName string, x, y int) string {
	templateData := struct {
		DatName    string
		OutName string
		X int
		Y int
	}{
		fileName,
		outputName,
		x,
		y,
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

func createPng( templateName string) {
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


	gnuplotOutput, err := ioutil.ReadAll(stdout)
	if err != nil {
		log.Fatal(err)
	}
	gnuplotOutputerr, err := ioutil.ReadAll(stderr)
	if err != nil {
		log.Fatal(err)
	}

	err = c.Wait()
	if err != nil {
		log.Fatal(err)
	}

	log.Print(gnuplotOutput)
	log.Print(gnuplotOutputerr)

}
