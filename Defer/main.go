package main

import "fmt"

func Add(a, b int) int {
	return a + b
}

func main() {
	fmt.Println("Starting 'defer'")

	fmt.Println("1")
	// Defer the printing of 2 until main returns.
	defer fmt.Println("2")
	fmt.Println("3")

	fmt.Println("Example with function")
	fmt.Println("Getting Value.....")
	// These deferred calls execute in reverse order when main ends.
	defer fmt.Println("Ending....")
	defer fmt.Println(Add(5, 5))
	defer fmt.Println(Add(2, 2))
	defer fmt.Println(Add(6, 6))
	fmt.Println("Printing answers....")

}

// defer: a keyword used to delay the execution of a function call until the surrounding function returns.
// What it does: registers a call to be executed later, after the current function finishes.
// Used for: cleanup tasks, releasing resources, closing files, and ensuring actions run even if a function exits early.
// Example:
// defer file.Close()
// defer fmt.Println("done")
