package main

import (
	"fmt"
)

func main() {
	fmt.Println("This is if-else statement")

	fmt.Println("Enter a number:")
	var x int
	fmt.Scan(&x)

	// if x > 50 {
	// 	fmt.Println("x is greater than 50")
	// } else if x > 20 {
	// 	fmt.Println("x is greater than 20 but less than or equal to 50")
	// } else if x < 40 {
	// 	fmt.Println("x is less than 40")
	// }

	if x > 50 && x < 100 && (x > 0) {
		fmt.Println("x is greater than 50 and less than 100 and greater than 0")
	} else if x > 100 {
		fmt.Println("x is greater than 100")
	} else {
		fmt.Println("x not be zero")
	}
}
