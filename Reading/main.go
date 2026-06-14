package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	// name := "Something"
	// age := 20
	// isStudent := true
	// height := 5.9
	// // Constants

	// const Pi = 3.14159

	// fmt.Println("Hello world")

	// fmt.Printf("type of name is %T \n", name)
	// fmt.Printf("type of age is %T \n", age)
	// fmt.Printf("type of isStudent is %T \n", isStudent)
	// fmt.Printf("type of height is %T \n", height)
	// fmt.Printf("Pi is %.2f \n", Pi)

	// var name string
	// var age int
	// var isStudent bool

	// fmt.Println("Enter name")
	// fmt.Scan(&name)
	// fmt.Println("Enter age")
	// fmt.Scan(&age)
	// fmt.Println("Are you a Student (true or false)")
	// fmt.Scan(&isStudent)

	// fmt.Println("Scaning......")
	// fmt.Println("Name : ", name)
	// fmt.Println("age : ", age)
	// fmt.Println("isStudent : ", isStudent)

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter name : ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)

	fmt.Println("Enter Age : ")
	age, _ := reader.ReadString('\n')
	age = strings.TrimSpace(age)

	fmt.Println("Enter isStudent : ")
	isStudent, _ := reader.ReadString('\n')
	isStudent = strings.TrimSpace(isStudent)

	fmt.Println("Name : ", name)
	fmt.Println("Age : ", age)
	fmt.Println("isStudent : ", isStudent)

}
