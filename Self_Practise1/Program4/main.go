package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("User Profile Builder")

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter First Name")
	Firstname, _ := reader.ReadString('\n')
	Firstname = strings.TrimSpace(Firstname)

	fmt.Println("Enter Last Name")
	Lastname, _ := reader.ReadString('\n')
	Lastname = strings.TrimSpace(Lastname)

	fmt.Println("Enter city")
	city, _ := reader.ReadString('\n')
	city = strings.TrimSpace(city)

	fmt.Printf("\nRecipient : %s, %s \nLocation: %s", Lastname, Firstname, city)
}
