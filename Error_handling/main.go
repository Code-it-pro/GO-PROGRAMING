package main

import (
	"fmt"
)

func divide(a, b float32) (float32, error) {
	if b == 0 {
		// err := "denominator should not be zero"
		// return 0, fmt.Errorf(err)

		return 0, fmt.Errorf("denominator should not be zero")
	}
	return a / b, nil
}
func main() {
	fmt.Println("Starting Error handling")
	var num1, num2 float32
	fmt.Println("Enter a Numerator:")
	fmt.Scan(&num1)
	fmt.Println("Enter Denominator:")
	fmt.Scan(&num2)

	fmt.Printf("Numerator : %f\n", num1)
	fmt.Printf("Denominator : %f\n", num2)

	div, err := divide(num1, num2)

	fmt.Printf("Division is %f", div)

	if err != nil {
		fmt.Println(err)
	}
}
