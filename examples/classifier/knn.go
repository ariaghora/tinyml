package main

import (
	"fmt"

	"ghora.net/tinyml/classifier"
	"ghora.net/tinyml/io"
	"ghora.net/tinyml/preprocessing"
)

func main() {
	IrisDataset, err := io.ReadCSV("../../assets/datasets/iris.csv")
	if err != nil {
		panic(err)
	}

	X := IrisDataset.Range(0, 0, IrisDataset.NRows, 4)
	Y := IrisDataset.GetCol(4)
	XTrain, XTest, YTrain, YTest := preprocessing.TrainTestSplit(X, Y, 0.2, true)

	k := 3
	classifier := classifier.NewKNNClassifier(k)
	classifier.Fit(XTrain, YTrain)
	YPred := classifier.Predict(XTest)

	fmt.Println("Predictions:\n", YPred.T())
	fmt.Println("Actual:\n", YTest.T())

	nCorrectlyClassified := 0.0
	for i := 0; i < YPred.NRows; i++ {
		if YPred.Data[i][0] == YTest.Data[i][0] {
			nCorrectlyClassified++
		}
	}
	accuracy := nCorrectlyClassified / float64(YPred.NRows)
	fmt.Println("Accuracy:", accuracy)
}
