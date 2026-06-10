package main

import (
	"fmt"
)

func main() {

	fmt.Println("Problem 1 Solutions")

	grocery := map[string]int{
		"Oats":       3,
		"snacks":     20,
		"Milk":       2,
		"Bread":      12,
		"Cheese":     3,
		"Eggs":       2,
		"Biscuts":    25,
		"Jelly":      24,
		"Sour Candy": 45,
		"Chocolates": 20,
	}

	LowStock := []string{}

	for item, quantity := range grocery {
		if quantity < 10 {
			fmt.Printf("%s: %d\n", item, quantity)
		} else {
			LowStock = append(LowStock, item)
		}
	}

	fmt.Println("\nLow in Stock Items : ")

	for i := 0; i < len(LowStock); i++ {
		fmt.Printf("%d. %s\n", i, LowStock[i])
	}

	fmt.Println("\nName an Item to delete from the List above")

	var delItem string

	for i := 0; i < 1; {
		fmt.Println("Enter an Item Name to delete Item\nEnter 'Add' to add an item\nEnter 'exit' to exit and print the Final List")
		fmt.Scan(&delItem)
		if _, exists := grocery[delItem]; exists {
			delete(grocery, delItem)
			fmt.Printf("%s deleted from the list\n", delItem)
		} else if delItem == "exit" {
			i = 1
			fmt.Println("Final List: ")
			for item, value := range grocery {
				fmt.Printf("%s : %d\n", item, value)
			}
		} else if delItem == "Add" {
			var newItem string
			var quantity int
			fmt.Println("Enter the Item name")
			fmt.Scan(&newItem)
			fmt.Println("Enter its quantity")
			fmt.Scan(&quantity)
			grocery[newItem] = quantity
			fmt.Printf("%s added to the list\n", newItem)
		} else {
			fmt.Println("Error: Item not Found")

		}
	}
}
