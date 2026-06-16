package main

import "fmt"

func main() {
	fmt.Println("Problem 3")

	prices := []float64{10.0, 20.0, 99.9, 120.0}

	// var val *int
	fmt.Println("Original prices:", prices)
	for i, _ := range prices {
		tax := 1.1
		ScaleValue(&prices[i], tax)
	}

	fmt.Println("Modified prices:", prices)

}

func ScaleValue(val *float64, factor float64) {
	*val = *val * factor
}
