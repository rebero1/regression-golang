package main

import (
	"fmt"
	"log"
	"os"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"

	"github.com/kniren/gota/dataframe"
)

func main() {

	fileData, err := os.Open("/Users/reberoprince/Documents/Golang books/codes/chapter3/regresion/Advertising.csv")

	if err != nil {
		fmt.Println(err)
	}
	defer fileData.Close()
	dataF := dataframe.ReadCSV(fileData)

	fmt.Printf("Description of DatafRAMe\n%v", dataF.Describe())

	for _, colName := range dataF.Names() {

		plotVals := make(plotter.Values, dataF.Nrow())

		for i, floatVal := range dataF.Col(colName).Float() {
			plotVals[i] = floatVal
		}

		p, err := plot.New()

		if err != nil {
			log.Fatal(err)
		}

		p.Title.Text = fmt.Sprintf("Histogram of a %s", colName)

		h, err := plotter.NewHist(plotVals, 16)

		if err != nil {
			fmt.Print(err)
		}

		h.Normalize(1)
		p.Add(h)

		if err := p.Save(4*vg.Inch, 4*vg.Inch, colName+"_hist.png"); err != nil {
			log.Fatal(err)
		}

	}

}
