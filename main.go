package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"math"
	"math/rand"
	"os"
	"sort"
	"strconv"
)

// bootstrapSample generates a bootstrap sample from the input data
func bootstrapSample(data []float64, sampleSize int) []float64 {
	sample := make([]float64, sampleSize)

	for i := 0; i < sampleSize; i++ {
		randomIndex := rand.Intn(len(data)) // Random index in the range of data
		sample[i] = data[randomIndex]
	}

	return sample
}

// median calculates the median of a dataset
func median(data []float64) float64 {
	sort.Float64s(data) // Sort the data

	n := len(data)
	if n%2 == 0 {
		return float64(data[n/2-1]+data[n/2]) / 2.0 // Average of middle two
	}
	return float64(data[n/2]) // Middle element
}

// standardError calculates the standard error of a dataset
func StandardError(data []float64) float64 {
	// Calculate mean
	sum := 0.0
	for _, x := range data {
		sum += x
	}
	mean := sum / float64(len(data))

	// Calculate standard deviation
	sumSqDiff := 0.0
	for _, x := range data {
		sumSqDiff += math.Pow(x-mean, 2)
	}
	stdDev := math.Sqrt(sumSqDiff / float64(len(data)-1))

	// Calculate standard error
	return stdDev / math.Sqrt(float64(len(data)))
}

// flatten takes a slice of slices of integers and returns a single slice
func Flatten(nestedSlices [][]string) []string {
	var flatSlice []string
	for _, subSlice := range nestedSlices {
		flatSlice = append(flatSlice, subSlice...) // Append all elements of subSlice
	}
	return flatSlice
}

func stringsToFloats(strings []string) ([]float64, error) {
	var numbers []float64
	for _, str := range strings {
		num, err := strconv.ParseFloat(str, 64) // Convert string to float
		if err != nil {
			return nil, fmt.Errorf("failed to convert '%s' to float: %v", str, err)
		}
		numbers = append(numbers, num)
	}
	return numbers, nil
}

// main demonstrates bootstrap sampling with medians and standard error
func main() {

	// open file from the first user argument
	f, err := os.Open("data/co2_data.csv")
	if err != nil {
		log.Fatal(err)
	}

	// remember to close the file at the end of the program
	defer f.Close()

	// Read CSV file using csv.Reader
	csvReader := csv.NewReader(f)
	data, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	// remove header
	data = data[1:]

	// Flatten the slice
	flat_data := Flatten(data)

	final_data, err := stringsToFloats(flat_data)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Number of bootstrap samples and sample size
	numSamples := 1000
	sampleSize := len(final_data)
	medians := make([]float64, numSamples)
	// Generate bootstrap samples and calculate medians
	for i := 0; i < numSamples; i++ {
		bootstrap := bootstrapSample(final_data, sampleSize)
		medians[i] = median(bootstrap)
	}
	// Calculate the standard error of the medians
	fmt.Println("Bootstrapped Medians:", medians)
	standardErr := StandardError(medians)
	medianOfMedians := median(medians)
	fmt.Printf("Median of Bootstrapped Medians  %.4f\n", medianOfMedians)
	fmt.Printf("Standard Error of Bootstrapped Medians: %.4f\n", standardErr)

}
