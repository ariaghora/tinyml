package preprocessing

import (
	"math/rand"

	"ghora.net/tinyml/matrix"
)

func TrainTestSplit(X, Y *matrix.Matrix, testSize float64, shuffle bool) (XTrain, XTest, YTrain, YTest *matrix.Matrix) {
	if testSize < 0 || testSize > 1 {
		panic("testSize must be between 0 and 1")
	}
	n := X.NRows
	nTest := int(float64(n) * testSize)
	nTrain := n - nTest

	var indices []int
	if shuffle {
		indices = rand.Perm(n)
	} else {
		indices = make([]int, n)
		for i := 0; i < n; i++ {
			indices[i] = i
		}
	}

	XTrainData := [][]float64{}
	YTrainData := [][]float64{}
	XTestData := [][]float64{}
	YTestData := [][]float64{}
	for i := 0; i < n; i++ {
		if i < nTrain {
			XTrainData = append(XTrainData, X.Data[indices[i]])
			YTrainData = append(YTrainData, Y.Data[indices[i]])
		} else {
			XTestData = append(XTestData, X.Data[indices[i]])
			YTestData = append(YTestData, Y.Data[indices[i]])
		}
	}
	XTrain = &matrix.Matrix{}
	XTrain.Data = XTrainData
	XTrain.NRows = nTrain
	XTrain.NCols = X.NCols
	YTrain = &matrix.Matrix{}
	YTrain.Data = YTrainData
	YTrain.NRows = nTrain
	YTrain.NCols = Y.NCols
	XTest = &matrix.Matrix{}
	XTest.Data = XTestData
	XTest.NRows = nTest
	XTest.NCols = X.NCols
	YTest = &matrix.Matrix{}
	YTest.Data = YTestData
	YTest.NRows = nTest
	YTest.NCols = Y.NCols

	return XTrain, XTest, YTrain, YTest
}
