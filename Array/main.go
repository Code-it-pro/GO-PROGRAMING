package main

import "fmt"

func main() {
	fmt.Println("Learning array in Go language")

	var arr [5]int
	arr2 := [5]int{1, 2, 3, 4, 5}
	// arr3 := make([]string, 4)

	fmt.Println(arr)
	fmt.Println(arr2)

	fmt.Printf("arr length : %d", len(arr))
	fmt.Printf("\narr2 length : %d \n", len(arr2))

	arr2[2] = 52352

	// arr3[0] = "Gaming"
	// arr3 = append(arr3, "Gagan", "Shona")

	arr[3] = 53
	arr[0] = 352
	fmt.Println(arr)
	// fmt.Println(arr3)
	// fmt.Println(len(arr3))

}
