package test

import (
	"fmt"
	"testing"

	"ghora.net/tinyml/matrix"
)

func TestColumnInsertion(t *testing.T) {
	mat := matrix.NewRandomMatrix(3, 4)
	ones := []float64{1, 1, 1}

	t.Run("CanInsertColumnAtBeginning", func(t *testing.T) {
		mat := mat.InsertColumnAt(0, ones)
		if mat.NCols != 5 {
			t.Error("Expected 5 columns, got", mat.NCols)
		}

		matArr := mat.AsFlatArray()
		fmt.Println(matArr)
	})

}

func TestMatAdd(t *testing.T) {
	mat1 := matrix.NewMatrixFromArray([][]float64{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	})
	mat2 := matrix.NewMatrixFromArray([][]float64{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	})
	t.Run("AdditionIsCorrect", func(t *testing.T) {
		result := matrix.MatAdd(mat1, mat2).AsFlatArray()
		expected := []float64{2, 4, 6, 8, 10, 12, 14, 16, 18}
		for i := 0; i < len(result); i++ {
			if result[i] != expected[i] {
				t.Error("Expected", expected, "got", result)
			}
		}
	})
}

func TestMatSub(t *testing.T) {
	mat1 := matrix.NewMatrixFromArray([][]float64{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	})
	mat2 := matrix.NewMatrixFromArray([][]float64{
		{1, 2, 3},
		{4, 0, 6},
		{7, 8, 9},
	})
	t.Run("SubtractionIsCorrect", func(t *testing.T) {
		result := matrix.MatSub(mat1, mat2).AsFlatArray()
		expected := []float64{0, 0, 0, 0, 5, 0, 0, 0, 0}
		for i := 0; i < len(result); i++ {
			if result[i] != expected[i] {
				t.Error("Expected", expected, "got", result)
			}
		}
	})
}

func TestMatTranspose(t *testing.T) {
	mat := matrix.NewMatrixFromArray([][]float64{
		{1, 2, 3},
		{4, 5, 6},
	})
	t.Run("TransposeIsCorrect", func(t *testing.T) {
		result := mat.T().AsFlatArray()
		expected := []float64{1, 4, 2, 5, 3, 6}
		for i := 0; i < len(result); i++ {
			if result[i] != expected[i] {
				t.Error("Expected", expected, "got", result)
			}
		}
	})
}

func TestMatScale(t *testing.T) {
	mat := matrix.NewFullMatrix(3, 3, 1)
	t.Run("ScalingIsCorrect", func(t *testing.T) {
		result := mat.Scale(2).AsFlatArray()
		expected := []float64{2, 2, 2, 2, 2, 2, 2, 2, 2}
		for i := 0; i < len(result); i++ {
			if result[i] != expected[i] {
				t.Error("Expected", expected, "got", result)
			}
		}
	})
}

func TestMatMul(t *testing.T) {
	mat := matrix.NewRandomMatrix(3, 4)
	ones := []float64{1, 1, 1}

	t.Run("CanMultiply", func(t *testing.T) {
		mat := mat.InsertColumnAt(0, ones)
		mat2 := matrix.NewRandomMatrix(5, 5)

		if mat.NCols != 5 {
			t.Error("Expected 5 columns, got", mat.NCols)
		}
		if mat2.NCols != 5 {
			t.Error("Expected 5 columns, got", mat2.NCols)
		}
	})

	t.Run("SimpleMultiplicationCorrect", func(t *testing.T) {
		mat1 := matrix.NewMatrixFromArray([][]float64{
			{1, 2},
			{3, 4},
			{5, 6},
		})
		mat2 := matrix.NewMatrixFromArray([][]float64{
			{1, 1},
			{1, 1},
		})

		result := mat1.Mul(mat2).AsFlatArray()
		expected := []float64{3, 3, 7, 7, 11, 11}

		for i := 0; i < len(result); i++ {
			if result[i] != expected[i] {
				t.Error("Expected", expected, "got", result)
			}
		}

	})
}
