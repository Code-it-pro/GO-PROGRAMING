package main

import "fmt"

func main() {
	fmt.Println("This is a program to demonstrate the use of for loops in Go.")

	// Basic for loop
	// for i := 0; i < 5; i++ {
	// 	fmt.Printf("Value of i: %d\n", i)
	// 	fmt.Println("Value of i:", i)
	// }

	// For loop an while
	// counter := 0
	// for {
	// 	fmt.Println("This is an infinite loop. Count:", counter)
	// 	counter++
	// 	if counter == 1000 {
	// 		break
	// 	}
	// }

	// For loop with range
	num := []int{12, 25, 346, 347, 345}
	for i, value := range num {
		fmt.Printf("Index : %d Value: %d \n", i, value)
	}
}
