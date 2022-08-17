package main

import (
	"fmt"

	"ghora.net/tinyml/matrix"
)

func main() {
	mat := matrix.NewRandomMatrix(3, 4)
	fmt.Println(mat)

	oneTwoThree := []float64{1, 2, 3}
	mat = mat.InsertColumnAt(0, oneTwoThree)
	fmt.Println(mat)
}
