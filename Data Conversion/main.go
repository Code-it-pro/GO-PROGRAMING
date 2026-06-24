package main

import (
	"fmt"
	"strconv"
)

func main() {
	// Define an integer value.
	num := 3245

	// Print the integer value itself.
	fmt.Println(num)
	// Print the type of the integer value.
	fmt.Printf("Type of Num : %T\n", num)

	// Convert the integer to a float32.
	converted_float := float32(num)
	// Print the converted float value.
	fmt.Println(converted_float)
	// Print the type of the converted float value.
	fmt.Printf("Type of converted_float : %T\n", converted_float)

	// Convert the integer to a string.
	converted_string := strconv.Itoa(num)
	// Print the converted string value.
	fmt.Println(converted_string)
	// Print the type of the converted string value.
	fmt.Printf("Type of converted_string : %T\n", converted_string)

	// Define a float value.
	new_num := 2.14
	fmt.Println(new_num)
	// Print the type of the float value.
	fmt.Printf("Type of new_num : %T\n", new_num)
	// Add the integer and float values by converting the integer to float64.
	var sum float64 = float64(num) + new_num
	fmt.Printf("Type of sum : %T\n", sum)

	// Convert a string to an integer.
	str := "1234"
	connum, _ := strconv.Atoi(str)

	// Add two integers.
	new_sum := num + connum
	fmt.Println(new_sum)
	// Print the type of the integer sum.
	fmt.Printf("Type of new_sum : %T\n", new_sum)
}

// strconv is a standard Go package for converting strings to basic
// data types and basic data types to strings.
// Examples used here: strconv.Itoa converts int to string and strconv.Atoi converts string to int.
