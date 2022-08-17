package classifier

import (
	"fmt"
	"math"

	"ghora.net/tinyml/matrix"
)

type LogisticRegression struct {
	K       int
	W       *matrix.Matrix
	Alpha   float64
	Mu      float64
	Fitted  bool
	MaxIter int
	Verbose bool
}

func NewLogisticRegression(alpha, mu float64, MaxIter int, verbose bool) *LogisticRegression {
	logreg := &LogisticRegression{Alpha: alpha, Mu: mu, MaxIter: MaxIter, Verbose: verbose}
	return logreg
}

func (cls *LogisticRegression) PredictProba(x *matrix.Matrix) *matrix.Matrix {
	if !cls.Fitted {
		panic("LogisticRegression not fitted")
	}
	if x == nil {
		panic("X is nil")
	}
	ones := make([]float64, x.NRows)
	for i := 0; i < x.NRows; i++ {
		ones[i] = 1.0
	}
	x = x.InsertColumnAt(0, ones)
	z := matrix.MatMul(x.Scale(-1.0), cls.W)
	probs := softmax(z)
	return probs
}

func (cls *LogisticRegression) Predict(x *matrix.Matrix) *matrix.Matrix {
	proba := cls.PredictProba(x)
	classMat := matrix.NewMatrix(proba.NRows, 1)
	for i := 0; i < proba.NRows; i++ {
		classMat.Data[i][0] = float64(arrArgMax(proba.Data[i]))
	}
	return classMat
}

func (cls *LogisticRegression) Fit(X *matrix.Matrix, Y *matrix.Matrix) *LogisticRegression {

	// Add intercept term to X
	ones := make([]float64, X.NRows)
	for i := 0; i < X.NRows; i++ {
		ones[i] = 1.0
	}
	XNew := X.InsertColumnAt(0, ones)
	YNew := oneHot(Y)

	cls.W = matrix.NewFullMatrix(X.NCols+1, YNew.NCols, 0.0)

	// Gradient descent
	for i := 0; i < cls.MaxIter; i++ {
		z := matrix.MatMul(XNew.Scale(-1.0), cls.W)
		probs := softmax(z)
		err := YNew.Sub(probs)
		dw := XNew.T().Mul(err).Scale(1.0 / float64(XNew.NRows)).Add(cls.W.Scale(cls.Mu * 2))
		cls.W = cls.W.Sub(dw.Scale(cls.Alpha))

		crossEntropy := probs.Apply(log).ElementwiseMul(YNew).Scale(-1.0)
		meanCrossEntropy := 0.0
		for i := 0; i < crossEntropy.NRows; i++ {
			meanCrossEntropy += crossEntropy.Data[i][0] / float64(XNew.NRows)
		}
		if i%100 == 0 {
			fmt.Println("Loss:", meanCrossEntropy)
		}
	}

	cls.Fitted = true
	return cls
}

func log(x float64) float64 {
	return math.Log(x)
}

func arrMax(arr []float64) float64 {
	max := arr[0]
	for _, v := range arr {
		if v > max {
			max = v
		}
	}
	return max
}

func arrArgMax(arr []float64) int {
	max := arr[0]
	argMax := 0
	for i := 1; i < len(arr); i++ {
		if arr[i] > max {
			max = arr[i]
			argMax = i
		}
	}
	return argMax
}

func arrSum(arr []float64) float64 {
	sum := 0.0
	for _, v := range arr {
		sum += v
	}
	return sum
}

func oneHot(x *matrix.Matrix) *matrix.Matrix {
	if x.NCols != 1 {
		panic("X must have 1 column")
	}
	nClass := int(arrMax(x.AsFlatArray()) + 1)
	Y := matrix.NewMatrix(x.NRows, nClass)
	for i := 0; i < x.NRows; i++ {
		Y.Data[i][int(x.Data[i][0])] = 1.0
	}
	return Y
}

func softmax(x *matrix.Matrix) *matrix.Matrix {
	res := matrix.NewMatrix(x.NRows, x.NCols)
	for i := 0; i < x.NRows; i++ {
		for j := 0; j < x.NCols; j++ {
			res.Data[i][j] = math.Exp(x.Data[i][j])
		}
		rowSum := arrSum(res.Data[i])
		for j := 0; j < x.NCols; j++ {
			res.Data[i][j] /= rowSum
		}
	}
	return res
}
