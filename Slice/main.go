package main

import "fmt"

func main() {
	fmt.Println("Starting with Slice in Go Programing")

	// initializing a slice
	// First way

	var numbers []int = []int{1, 2, 3, 4, 5}

	fmt.Println("Numbers:", numbers)
	fmt.Printf("Length of numbers: %d\n", len(numbers))
	fmt.Printf("Capacity of numbers: %d\n", cap(numbers))

	num1 := make([]int, 5)

	num1[2] = 12
	num1[3] = 13
	num1[4] = 14

	fmt.Println("Num1:", num1)
	fmt.Printf("Length of num1: %d\n", len(num1))
	fmt.Printf("Capacity of num1: %d\n", cap(num1))

	num1 = append(num1, 13, 14, 15)
	fmt.Println("Num1:", num1)
	fmt.Printf("Length of num1: %d\n", len(num1))
	fmt.Printf("Capacity of num1: %d\n", cap(num1))

	numbers = append(numbers, 6, 7, 8)
	fmt.Println("Numbers:", numbers)
	fmt.Printf("Length of numbers: %d\n", len(numbers))
	fmt.Printf("Capacity of numbers: %d\n", cap(numbers))

	//if array comparision

	arr1 := [3]int{1, 2, 3}

	fmt.Println("Array 1:", arr1)
	fmt.Printf("Length of arr1: %d\n", len(arr1))
	fmt.Printf("Capacity of arr1: %d\n", cap(arr1))

	arrSlice := arr1[:]
	arrSlice = append(arrSlice, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15)
	fmt.Println("Array 1:", arrSlice)
	fmt.Printf("Length of arr1: %d\n", len(arrSlice))
	fmt.Printf("Capacity of arr1: %d\n", cap(arrSlice))

}
