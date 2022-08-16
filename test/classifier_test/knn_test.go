package classifier_test

import (
	"testing"

	"ghora.net/tinyml/classifier"
	"ghora.net/tinyml/io"
)

func TestFit(t *testing.T) {

	data, _ := io.ReadCSV("../../assets/datasets/dummy.csv")
	knn := classifier.NewKNNClassifier(1).Fit(data, data)

	t.Run("FitSuccess", func(t *testing.T) {
		if !knn.Fitted {
			t.Error("Expected knn to be fitted")
		}
	})

	t.Run("MemorizedDataNotNil", func(t *testing.T) {
		if knn.X == nil {
			t.Error("Expected knn.Data to be non-nil, got", knn.X)
		}
	})
}
