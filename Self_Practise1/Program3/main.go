package main

import "fmt"

func main() {
	fmt.Println("Even-Odd Filter")
	number := []int{11, 22, 33, 44, 55, 66, 77, 88, 99}
	for _, value := range number {
		if value%2 == 0 {
			fmt.Printf("Even: %d ", value)
		} else {
			fmt.Printf("Odd: %d ", value)
		}
	}
}
