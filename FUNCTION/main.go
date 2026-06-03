package main

import (
	"fmt"
)

// func SimpleFunction() {
// 	fmt.Println("I am Simple function")
// }

func add(a, b int) int {
	return a + b
}

func multiply(a, b int) int {
	result := a * b
	return result
}

func main() {
	fmt.Println("This is a function program")
	// SimpleFunction()
	sum := add(10, 20)
	mult := multiply(10, 2)
	fmt.Println(" Sum =", sum)
	fmt.Println(" Mult =", mult)
}
