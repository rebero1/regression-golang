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

func grid(dataDF dataframe.DataFrame) {
	yVals := dataDF.Col("Sales").Float()

	for _, colName := range dataDF.Names() {

		pts := make(plotter.XYs, dataDF.Nrow())

		for i, floatVal := range dataDF.Col(colName).Float() {

			pts[i].X = floatVal
			pts[i].Y = yVals[i]
		}
		p, err := plot.New()

		if err != nil {
			log.Fatal(err)
		}

		p.X.Label.Text = colName
		p.Y.Label.Text = "y"
		p.Add(plotter.NewGrid())
		s, err := plotter.NewScatter(pts)

		if err != nil {
			log.Fatal(err)
		}

		s.GlyphStyle.Radius = vg.Points(3)

		p.Add(s)

		if err := p.Save(4*vg.Inch, 4*vg.Inch, colName+"_scatter.png"); err != nil {
			log.Fatal(err)
		}
	}
}

func histogram(dataDF dataframe.DataFrame) {
	// Use The describe method to summarize
	dataSummary := dataDF.Describe()

	fmt.Println(dataSummary)

	for _, colName := range dataDF.Names() {

		plotVals := make(plotter.Values, dataDF.Nrow())

		for i, floatVal := range dataDF.Col(colName).Float() {
			plotVals[i] = floatVal
		}

		p, err := plot.New()

		if err != nil {
			log.Fatal()
		}

		p.Title.Text = fmt.Sprintf("Histogram of a %s", colName)

		h, err := plotter.NewHist(plotVals, 16)

		if err != nil {
			log.Fatal(err)
		}

		h.Normalize(1)

		p.Add(h)

		if err := p.Save(4*vg.Inch, 4*vg.Inch, colName+"_hist.png"); err != nil {
			log.Fatal(err)
		}

	}
}
func main() {
	// The first step is profiling the dataset

	advertFile, err := os.Open("Advertising.csv")

	if err != nil {
		log.Fatal(err)
	}
	defer advertFile.Close()
	dataDF := dataframe.ReadCSV(advertFile)

	histogram(dataDF)
	grid(dataDF)

}
