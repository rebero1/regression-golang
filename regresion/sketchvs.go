package main

import (
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
		log.Fatal(err)
	}

	defer fileData.Close()

	advertDF := dataframe.ReadCSV(fileData)
	yVals := advertDF.Col("Sales").Float()

	for _, colName := range advertDF.Names() {
		pts := make(plotter.XYs, advertDF.Nrow())

		for i, floatVal := range advertDF.Col(colName).Float() {

			pts[i].X = floatVal
			pts[i].Y = yVals[i]
		}

		p, err := plot.New()

		if err != nil {
			log.Fatal(err)
		}

		p.X.Label.Text = colName
		p.Y.Label.Text = "y"

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
