package main

import "fmt"

func main() {
	fmt.Println("Starting Struct or Structure")

	// type person struct {
	// 	name   string
	// 	age    int
	// 	mobile int
	// }

	// var Some1 person
	// Some1.name = "Gagandeep Singh"
	// Some1.age = 22
	// Some1.mobile = 35253235912

	// fmt.Println(Some1)

	// type address struct {
	// 	city        string
	// 	Area        string
	// 	Housenumber int
	// 	Landmark    string
	// }

	// type Person struct {
	// 	Name    string
	// 	age     int
	// 	number  int
	// 	address address
	// }

	// Human := Person{
	// 	Name: "Gagandeep Singh",
	// 	age:  22,
	// 	address: address{
	// 		city: "Chandigarh",
	// 		Area: "Sector 17",
	// 	},
	// }

	// fmt.Println(Human)

	type person struct {
		Name   string
		age    int
		Gender string
	}

	type address struct {
		country  string
		state    string
		city     string
		area     string
		landmark string
		houseno  int
		pincode  int
	}

	type contact struct {
		phoneno     int
		email       string
		alternateno int
	}

	type Alldetails struct {
		Person  person
		Address address
		Contact contact
	}

	Gagandeep := Alldetails{
		Person: person{
			Name:   "Gagandeep Singh",
			age:    22,
			Gender: "Male",
		},
		Address: address{
			country:  "Texas",
			state:    "Los Santos",
			city:     "Ranches",
			area:     "Skinwalker Ranch",
			landmark: "Scary Gas Station",
			houseno:  666,
		},
		Contact: contact{
			phoneno:     3673734745,
			email:       "afwafwf@adawfa.com",
			alternateno: 24626262624,
		},
	}

	fmt.Printf("All Details of %s\n", Gagandeep.Person.Name)
	fmt.Printf("Age: %d\n", Gagandeep.Person.age)
	fmt.Printf("Gender: %s\n", Gagandeep.Person.Gender)
	fmt.Printf("Country: %s\n", Gagandeep.Address.country)
	fmt.Printf("State: %s\n", Gagandeep.Address.state)
	fmt.Printf("City: %s\n", Gagandeep.Address.city)
	fmt.Printf("Area: %s\n", Gagandeep.Address.area)
	fmt.Printf("Landmark: %s\n", Gagandeep.Address.landmark)
	fmt.Printf("Houseno: %d\n", Gagandeep.Address.houseno)
	fmt.Printf("Pincode: %d\n", Gagandeep.Address.pincode)

	fmt.Printf("Phoneno: %d\n", Gagandeep.Contact.phoneno)
	fmt.Printf("Email: %s\n", Gagandeep.Contact.email)
	fmt.Printf("Alternateno: %d\n", Gagandeep.Contact.alternateno)

}
