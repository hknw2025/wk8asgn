package main

import (
	"math"
	"testing"
)

func TestStandardError(t *testing.T) {

	test_data := []float64{179, 160, 136, 227, 123, 23, 45, 67, 1, 234}
	standardErr := StandardError(test_data)
	standardErr = math.Round(standardErr*100) / 100

	if standardErr != 26.20 {
		t.Errorf("Standard Error was incorrect, got: %f, want: %f.", standardErr, 26.20)
	}

}

func BenchmarkMain(b *testing.B) {
	for i := 0; i < b.N; i++ {
		main()
	}
}
