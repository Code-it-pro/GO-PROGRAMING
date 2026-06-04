package main

import "fmt"

func main() {
	fmt.Println("This is a switch case example in Go")

	//Switch case in Go

	// var day int
	// fmt.Print("Enter a day number (1-7): ")
	// fmt.Scan(&day)

	// switch day {
	// case 1:
	// 	fmt.Println("Monday")
	// case 2:
	// 	fmt.Println("Tuesday")
	// case 3:
	// 	fmt.Println("Wednesday")
	// case 4:
	// 	fmt.Println("Thursday")
	// case 5:
	// 	fmt.Println("Friday")
	// case 6:
	// 	fmt.Println("Saturday")
	// case 7:
	// 	fmt.Println("Sunday")
	// default:
	// 	fmt.Println("Invalid day")
	// }

	var temperature int
	fmt.Print("Enter the temperature: ")
	fmt.Scan(&temperature)

	switch {
	case temperature < 0:
		fmt.Println("It's freezing!")
		fmt.Println("Wear a coat!")
	case temperature >= 0 && temperature < 15:
		fmt.Println("It's cold!")
		fmt.Println("Wear a jacket!")
	case temperature >= 15 && temperature < 25:
		fmt.Println("It's warm!")
		fmt.Println("Wear a t-shirt!")
	case temperature >= 100:
		fmt.Println("It's boiling!")
		fmt.Println("Stay indoors!")
		fmt.Println("firefighters are on their way!")
	default:
		fmt.Println("It's hot!")
		fmt.Println("Wear light clothing!")
	}

}
