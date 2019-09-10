package main

import (
	"bufio"
	"log"
	"os"

	"github.com/kniren/gota/dataframe"
)

func main() {
	fileData, err := os.Open("/Users/reberoprince/Documents/Golang books/codes/chapter3/regresion/Advertising.csv")

	if err != nil {
		log.Fatal(err)
	}
	defer fileData.Close()

	advertDF := dataframe.ReadCSV(fileData)

	trainingNum := (4 * advertDF.Nrow()) / 5

	testNum := advertDF.Nrow() / 5

	if trainingNum+testNum < advertDF.Nrow() {
		trainingNum++
	}
	trainingNumIdx := make([]int, trainingNum)
	testIdx := make([]int, testNum)

	for index := 0; index < trainingNum; index++ {
		trainingNumIdx[index] = index
	}

	for index := 0; index < testNum; index++ {
		testIdx[index] = trainingNum + index
	}

	trainingDF := advertDF.Subset(trainingNumIdx)
	testDF := advertDF.Subset(testIdx)

	setMap := map[int]dataframe.DataFrame{
		0: trainingDF,
		1: testDF,
	}

	for idx, setName := range []string{"training.csv", "test.csv"} {

		f, err := os.Create(setName)
		if err != nil {
			log.Fatal(err)
		}

		w := bufio.NewWriter(f)

		if err := setMap[idx].WriteCSV(w); err != nil {
			log.Fatal(err)
		}

	}

}
