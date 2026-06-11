package main

import "fmt"

func main() {
	fmt.Println("Starting With Pointer")

	// var num int = 10
	// var ptr *int
	// fmt.Println("Value in Num: ", num)
	// fmt.Println("Pointer value: ", ptr)

	// ptr = &num
	// fmt.Println("Pointer value: ", ptr)
	// fmt.Println("Data at Pointer: ", *ptr)

	number := 20
	pointer := &number

	Addvalue(pointer)

}

func Addvalue(point *int) {
	fmt.Println("Imported pointer address: ", point)
	fmt.Println("Imported data: ", *point)

	*point *= 2

	fmt.Println("Imported data: ", point)
	fmt.Println("Imported data: ", *point)

}
