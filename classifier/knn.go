package classifier

import (
	"fmt"
	"sort"

	"ghora.net/tinyml/matrix"
)

type KNNClassifier struct {
	K      int
	X      *matrix.Matrix
	Y      *matrix.Matrix
	Fitted bool
}

func NewKNNClassifier(K int) *KNNClassifier {
	return &KNNClassifier{K: K}
}

func (cls *KNNClassifier) Fit(X *matrix.Matrix, Y *matrix.Matrix) *KNNClassifier {
	cls.X = X
	cls.Y = Y
	cls.Fitted = true
	return cls
}

func euclideanDistance(x, y []float64) float64 {
	if len(x) != len(y) {
		panic("x and y must have the same length")
	}
	sum := 0.0
	for i := 0; i < len(x); i++ {
		sum += (x[i] - y[i]) * (x[i] - y[i])
	}
	return sum
}

func (cls *KNNClassifier) Predict(X *matrix.Matrix) *matrix.Matrix {
	if !cls.Fitted {
		panic("KNNClassifier not fitted")
	}
	if X == nil {
		panic("X is nil")
	}

	closestIndices := [][]int{}
	for i := 0; i < X.NRows; i++ {
		distancesFromIthSample := []float64{}
		for j := 0; j < cls.X.NRows; j++ {
			dist := euclideanDistance(X.Data[i], cls.X.Data[j])
			distancesFromIthSample = append(distancesFromIthSample, dist)
		}
		closestIndicesFromIthSample := argsort(distancesFromIthSample)
		closestIndices = append(closestIndices, closestIndicesFromIthSample)
	}
	YPredData := [][]float64{}
	for i := 0; i < X.NRows; i++ {
		KNearestIndices := closestIndices[i][:cls.K]
		KNearestClasses := []float64{}
		for _, j := range KNearestIndices {
			KNearestClasses = append(KNearestClasses, cls.Y.Data[j][0])
		}
		// mode
		modeVal := []float64{}
		modeVal = append(modeVal, float64(mode(KNearestClasses)))

		YPredData = append(YPredData, modeVal)
	}
	YPred := new(matrix.Matrix)
	YPred.Data = YPredData
	YPred.NRows = X.NRows
	YPred.NCols = 1

	return YPred
}

func (cls *KNNClassifier) String() string {
	return fmt.Sprintf("KNNClassifier{K: %d}", cls.K)
}

type argsortStruct struct {
	s    []float64 // Points to orignal array
	inds []int     // Indices to be returned.
}

func (a argsortStruct) Len() int {
	return len(a.s)
}

func (a argsortStruct) Less(i, j int) bool {
	return a.s[a.inds[i]] < a.s[a.inds[j]]
}

func (a argsortStruct) Swap(i, j int) {
	a.inds[i], a.inds[j] = a.inds[j], a.inds[i]
}

func argsort(src []float64) []int {
	inds := make([]int, len(src))
	for i := range src {
		inds[i] = i
	}
	_argsort(src, inds)
	return inds
}

func _argsort(src []float64, inds []int) {
	if len(src) != len(inds) {
		panic("floats: length of inds does not match length of slice")
	}
	a := argsortStruct{s: src, inds: inds}
	sort.Sort(a)
}

func mode(testArray []float64) (mode float64) {
	countMap := make(map[float64]int)
	for _, value := range testArray {
		countMap[value] += 1
	}

	max := 0
	for _, key := range testArray {
		freq := countMap[key]
		if freq > max {
			mode = key
			max = freq
		}
	}
	return
}
