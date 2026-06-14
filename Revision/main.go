package main

import (
	revision "PracticeGo/Revision/test"
	"fmt"
)

func main() {
	fmt.Println("I am a New Main")

	revision.Rev("Hello from Revision")
}

// package main

// import "fmt"

// func main() {
// 	fmt.Println("Testing")

// 	var num []int
// 	num = append(num, 123)
// 	n := []int{}
// 	n = append(n, 23)

// 	fmt.Println(num)
// 	fmt.Println(n)
// 	fmt.Printf("%T\n", num)
// 	fmt.Printf("%T\n", n)
// }
