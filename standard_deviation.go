package main

import (
	"fmt"
	"math"
)

// standardDeviation calculates the standard deviation of a slice of float64 numbers.
func standardDeviation(nums []float64) float64 {
	// Step 1: Calculate the mean
	var sum float64
	for _, num := range nums {
		sum += num
	}
	mean := sum / float64(len(nums))

	// Step 2 & 3: Calculate the mean of the squared differences from the Mean
	var sumOfSquares float64
	for _, num := range nums {
		sumOfSquares += math.Pow(num-mean, 2)
	}
	meanOfSquares := sumOfSquares / float64(len(nums))

	// Step 4: Take the square root
	return math.Sqrt(meanOfSquares)
}

func main() {
	nums := []float64{1.0, 2.0, 3.0, 4.0, 5.0}
	fmt.Println("Standard Deviation:", standardDeviation(nums))
}
