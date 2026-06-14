package main

import (
	"fmt"
)

func main() {
	fmt.Println("The Calculator using Switch cases")
	var a, b int
	fmt.Println("enter a number")
	fmt.Scan(&a)
	fmt.Println("enter a number")
	fmt.Scan(&b)

	var op string
	fmt.Println("Enter the operator '+ , - , * , /' ")
	fmt.Scan(&op)
	result := calc(a, b, op)
	fmt.Println("Result:", result)
}

func calc(a, b int, op string) int {
	switch op {
	case "+":
		return a + b
	case "-":
		return a - b
	case "*":
		return a * b
	case "/":
		if b == 0 {
			fmt.Println("divided by zero")
			return 0
		}
		return a / b
	default:
		fmt.Println("Unknown operator")
	}
	return 0
}
