package io

import (
	"encoding/csv"
	"errors"
	"io"
	"log"
	"os"
	"strconv"
	"strings"

	"ghora.net/tinyml/matrix"
)

func ReadCSV(path string) (*matrix.Matrix, error) {
	// It is assumed that all values are floats
	f, err := os.Open(path)

	if err != nil {
		return nil, errors.New("could not open file")
	}

	reader := csv.NewReader(f)
	rows := [][]float64{}
	for {
		line, err := reader.Read()
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatal(err)
		}

		row := []float64{}
		for value := range line {
			val, err := strconv.ParseFloat(strings.TrimSpace(line[value]), 64)
			if err != nil {
				log.Fatal(err)
			}
			row = append(row, val)
		}
		rows = append(rows, row)
	}

	res := new(matrix.Matrix)
	res.Data = rows
	res.NRows = len(res.Data)
	res.NCols = len(res.Data[0])
	return res, nil
}
