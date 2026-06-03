package main

import "fmt"

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

	div, err := divide(10, 0)

	fmt.Println("Division is ", div)

	if err != nil {
		fmt.Println(err)
	}
}
