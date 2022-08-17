package matrix

import (
	"math/rand"
	"strconv"
)

type Matrix struct {
	Data  [][]float64
	NRows int
	NCols int
}

func (m *Matrix) Range(Row, Col, NRows, NCols int) *Matrix {
	NewData := [][]float64{}
	for i := Row; i < Row+NRows; i++ {
		NewData = append(NewData, m.Data[i][Col:Col+NCols])
	}
	NewMatrix := new(Matrix)
	NewMatrix.Data = NewData
	NewMatrix.NRows = NRows
	NewMatrix.NCols = NCols
	return NewMatrix
}

func (m *Matrix) GetRow(Row int) *Matrix {
	return m.Range(Row, 0, 1, m.NCols)
}

func (m *Matrix) GetCol(Col int) *Matrix {
	return m.Range(0, Col, m.NRows, 1)
}

func (m *Matrix) T() *Matrix {
	newData := [][]float64{}
	for i := 0; i < m.NCols; i++ {
		newRow := []float64{}
		for j := 0; j < m.NRows; j++ {
			newRow = append(newRow, m.Data[j][i])
		}
		newData = append(newData, newRow)
	}
	newMatrix := new(Matrix)
	newMatrix.Data = newData
	newMatrix.NRows = m.NCols
	newMatrix.NCols = m.NRows
	return newMatrix
}

func (m *Matrix) String() string {
	res := ""
	for i := 0; i < m.NRows; i++ {
		for j := 0; j < m.NCols; j++ {
			res += strconv.FormatFloat(m.Data[i][j], 'f', 2, 64) + ", "
		}
		res += "\n"
	}
	res += "(Matrix of " + strconv.FormatInt(int64(m.NRows), 10) + "x" + strconv.FormatInt(int64(m.NCols), 10) + ")"
	return res
}

func (m *Matrix) InsertColumnAt(colIndex int, data []float64) *Matrix {
	if len(data) != m.NRows {
		panic("len(data) != m.NRows")
	}
	newData := [][]float64{}
	for i := 0; i < m.NRows; i++ {
		newRow := []float64{}
		idx := 0
		for j := 0; j < m.NCols+1; j++ {
			if j == colIndex {
				newRow = append(newRow, data[i])
			} else {
				newRow = append(newRow, m.Data[i][idx])
				idx++
			}
		}
		newData = append(newData, newRow)
	}
	newMatrix := new(Matrix)
	newMatrix.Data = newData
	newMatrix.NRows = m.NRows
	newMatrix.NCols = m.NCols + 1
	return newMatrix
}

func (m *Matrix) InsertRowAt(rowIndex int, data []float64) *Matrix {
	if len(data) != m.NCols {
		panic("len(data) != m.NCols")
	}
	newData := [][]float64{}
	for i := 0; i < m.NRows; i++ {
		if i == rowIndex {
			newData = append(newData, data)
		} else {
			newData = append(newData, m.Data[i])
		}
	}
	newMatrix := new(Matrix)
	newMatrix.Data = newData
	newMatrix.NRows = m.NRows + 1
	newMatrix.NCols = m.NCols
	return newMatrix
}

func (m *Matrix) AsFlatArray() []float64 {
	res := []float64{}
	for i := 0; i < m.NRows; i++ {
		for j := 0; j < m.NCols; j++ {
			res = append(res, m.Data[i][j])
		}
	}
	return res
}

func (m *Matrix) Apply(f func(float64) float64) *Matrix {
	return MatUFunc(m, f)
}

func (m *Matrix) Scale(scalar float64) *Matrix {
	return MatScale(m, scalar)
}

/* Shortcut for binary function as methods */

func (m *Matrix) Add(other *Matrix) *Matrix {
	return MatAdd(m, other)
}

func (m *Matrix) Sub(other *Matrix) *Matrix {
	return MatSub(m, other)
}

func (m *Matrix) ElementwiseMul(other *Matrix) *Matrix {
	return MatElementwiseMul(m, other)
}

func (m *Matrix) Mul(other *Matrix) *Matrix {
	return MatMul(m, other)
}

/* Initialization functions */
func NewMatrix(NRows, NCols int) *Matrix {
	m := new(Matrix)
	m.NRows = NRows
	m.NCols = NCols
	m.Data = make([][]float64, NRows)
	for i := 0; i < NRows; i++ {
		m.Data[i] = make([]float64, NCols)
	}
	return m
}

func NewMatrixFromArray(data [][]float64) *Matrix {
	m := new(Matrix)
	m.NRows = len(data)
	m.NCols = len(data[0])
	m.Data = data
	return m
}

func NewMatrixFromFlatArray(data []float64) *Matrix {
	m := new(Matrix)
	m.NRows = len(data)
	m.NCols = 1
	m.Data = [][]float64{}
	for i := 0; i < m.NRows; i++ {
		m.Data = append(m.Data, []float64{data[i]})
	}
	return m
}

func NewFullMatrix(NRows, NCols int, value float64) *Matrix {
	m := NewMatrix(NRows, NCols)
	for i := 0; i < NRows; i++ {
		for j := 0; j < NCols; j++ {
			m.Data[i][j] = value
		}
	}

	return m
}

func NewRandomMatrix(NRows, NCols int) *Matrix {
	m := NewMatrix(NRows, NCols)
	for i := 0; i < NRows; i++ {
		for j := 0; j < NCols; j++ {
			m.Data[i][j] = rand.Float64()
		}
	}

	return m
}

/* Arithmetic operations */

func MatMul(A, B *Matrix) *Matrix {
	if A.NCols != B.NRows {
		panic("A.NCols != B.NRows: " + strconv.Itoa(A.NCols) + " != " + strconv.Itoa(B.NRows))
	}
	newData := [][]float64{}
	for i := 0; i < A.NRows; i++ {
		newRow := []float64{}
		for j := 0; j < B.NCols; j++ {
			newRow = append(newRow, 0.0)
			for k := 0; k < A.NCols; k++ {
				newRow[j] += A.Data[i][k] * B.Data[k][j]
			}
		}
		newData = append(newData, newRow)
	}
	newMatrix := new(Matrix)
	newMatrix.Data = newData
	newMatrix.NRows = A.NRows
	newMatrix.NCols = B.NCols
	return newMatrix
}

func MatBFunc(A, B *Matrix, f func(float64, float64) float64) *Matrix {
	if A.NCols != B.NCols || A.NRows != B.NRows {
		panic("A.NCols != B.NCols || A.NRows != B.NRows:" + strconv.Itoa(A.NCols) + " != " + strconv.Itoa(B.NCols) + " || " + strconv.Itoa(A.NRows) + " != " + strconv.Itoa(B.NRows))
	}
	newData := [][]float64{}
	for i := 0; i < A.NRows; i++ {
		newRow := []float64{}
		for j := 0; j < B.NCols; j++ {
			newRow = append(newRow, f(A.Data[i][j], B.Data[i][j]))
		}
		newData = append(newData, newRow)
	}
	newMatrix := new(Matrix)
	newMatrix.Data = newData
	newMatrix.NRows = A.NRows
	newMatrix.NCols = B.NCols
	return newMatrix
}

func MatUFunc(A *Matrix, f func(float64) float64) *Matrix {
	newData := [][]float64{}
	for i := 0; i < A.NRows; i++ {
		newRow := []float64{}
		for j := 0; j < A.NCols; j++ {
			newRow = append(newRow, 0.0)
			newRow[j] = f(A.Data[i][j])
		}
		newData = append(newData, newRow)
	}
	newMatrix := new(Matrix)
	newMatrix.Data = newData
	newMatrix.NRows = A.NRows
	newMatrix.NCols = A.NCols
	return newMatrix
}

func MatAdd(A, B *Matrix) *Matrix {
	return MatBFunc(A, B, func(a, b float64) float64 {
		return a + b
	})
}

func MatSub(A, B *Matrix) *Matrix {
	return MatBFunc(A, B, func(a, b float64) float64 {
		return a - b
	})
}

func MatElementwiseMul(A, B *Matrix) *Matrix {
	return MatBFunc(A, B, func(a, b float64) float64 {
		return a * b
	})
}

func MatElementwiseDiv(A, B *Matrix) *Matrix {
	return MatBFunc(A, B, func(a, b float64) float64 {
		return a / b
	})
}

func MatScale(A *Matrix, scalar float64) *Matrix {
	newData := [][]float64{}
	for i := 0; i < A.NRows; i++ {
		newRow := []float64{}
		for j := 0; j < A.NCols; j++ {
			newRow = append(newRow, scalar*A.Data[i][j])
		}
		newData = append(newData, newRow)
	}
	newMatrix := new(Matrix)
	newMatrix.Data = newData
	newMatrix.NRows = A.NRows
	newMatrix.NCols = A.NCols
	return newMatrix
}
