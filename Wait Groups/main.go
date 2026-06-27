package main

import (
	"fmt"
	"sync"
)

func Task(i int, Queue *sync.WaitGroup) {
	defer Queue.Done()
	fmt.Printf("Task %d Started\n", i)
	fmt.Printf("Task %d Ended\n", i)
}

func main() {
	var Queue sync.WaitGroup
	fmt.Println("Learning about waiting code.........")
	for i := 1; i <= 3; i++ {
		Queue.Add(1)
		go Task(i, &Queue)
	}
	Queue.Wait()
	fmt.Println("All Tasks Completed")
}
