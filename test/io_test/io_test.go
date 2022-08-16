package test

import (
	"testing"

	tio "ghora.net/tinyml/io"
)

func TestValidShape(t *testing.T) {
	mat, err := tio.ReadCSV("../../assets/datasets/dummy.csv")

	t.Run("CanOpenFile", func(t *testing.T) {
		if err != nil {
			t.Error(err)
		}
	})

	// test that the array has 3 rows
	t.Run("Has3Rows", func(t *testing.T) {
		if mat.NRows != 3 {
			t.Error("Expected 3 rows, got", mat.NRows)
		}
	})

	if mat.NCols != 4 {
		t.Error("Expected 4 columns, got", mat.NCols)
	}

}
