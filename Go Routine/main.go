package main

import (
	"fmt"
	"time"
)

func Hello() {
	fmt.Println("Hello World")
	time.Sleep(3000 * time.Millisecond)
	fmt.Println("after 4 second")
	fmt.Println("\nHello Im Oola")

}

func Hi() {
	fmt.Println("Hello Gagandeep")
	time.Sleep(2000 * time.Millisecond)
	fmt.Println("Hi Oola!")
}
func main() {
	fmt.Println("Learning Goroutine")
	go Hello()
	go Hi()

	time.Sleep(4000 * time.Millisecond)
}
