package main 



import(
	"fmt"
	"os"
	"encoding/csv"
	"github.com/sajari/regression"
)


func main()  {
	f, err:=os.Open("training.csv")

	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	reader := csv.NewReader(f)
	reader.FildPerRecord  =4
	trainingData,err :=reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	var r regression.Regression


}