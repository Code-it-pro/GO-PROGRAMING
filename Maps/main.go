package main

import "fmt"

func main() {
	fmt.Println("Starting Maps in Go")

	// Creating a map
	balance := make(map[string]int)

	// Adding key-value pairs to the map
	balance["Gagan"] = 10000000
	balance["Shona"] = 500000000
	balance["Mehar"] = balance["Gagan"] + balance["Shona"]
	balance["SomeOneElse"] = 4000

	// Accessing values from the map

	for name, value := range balance {
		fmt.Printf("%s's balance : %d\n", name, value)
	}

	fmt.Println("Deleting SomeOneElse From Map")

	delete(balance, "SomeOneElse")

	//Checking if SomeOneElse is present
	_, present := balance["SomeOneElse"]
	if !present {
		fmt.Println("SomeOneElse Removed")
	}

	for name, value := range balance {

		fmt.Printf("%s Value : %d\n", name, value)

	}

	fmt.Println(len(balance))

	//Different way to initialise

	name := map[string]int{
		"Annie":  23,
		"Doobie": 234,
	}

	for n, value := range name {
		fmt.Printf("%s : %d \n", n, value)
	}

}
