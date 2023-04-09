package main

import "fmt"

func main() {
  // Two arrays with two different channels
  a := []int{1, 2, 3}
  b := []float64{4.5, 5.5, 6.5}

  // Create a new array to store the result
  c := make([]float64, len(a))

  // Add the corresponding elements of a and b
  for i := 0; i < len(a); i++ {
    c[i] = float64(a[i]) + b[i]
  }

  // Print the result
  fmt.Println(c)
}
