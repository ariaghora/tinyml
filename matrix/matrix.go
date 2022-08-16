package matrix

import "strconv"

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
