package main

import "fmt"

func divide(a, b float64) (float64, string) {
	if b == 0 {
		err := "denominator should not be zero"
		return 0, ("Error: " + err)
	}
	return a / b, "nil"
}
func main() {
	fmt.Println("Starting Error handling")

	div, err := divide(10, 0)

	fmt.Println("Division is ", div)

	fmt.Println(err)
}
