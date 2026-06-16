package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Book struct {
	Title      string
	Author     string
	IsBorrowed bool
}

func BorrowBook(book []*Book, title string) error {

	for _, b := range book {
		if b.Title != title {
			continue
		}
		if b.IsBorrowed {
			return fmt.Errorf("book %s already borrowed", b.Title)
		}
		b.IsBorrowed = true
		return nil
	}
	return fmt.Errorf("Book %s not found", title)
}

func callme(lib []*Book) {
	fmt.Println("Enter Name of the Book:")
	reader := bufio.NewReader(os.Stdin)
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)

	fmt.Println(BorrowBook(lib, name))

	fmt.Printf("\nLibrary Books\n")
	for i := 0; i < 2; i++ {
		fmt.Println("Book:", lib[i].Title, "Borrowed Status:", lib[i].IsBorrowed)
	}
}

func main() {
	fmt.Println("Problem 4")

	library := []*Book{
		{Title: "The Go Programming Language",
			Author:     "Donovan & Kernighan",
			IsBorrowed: false},
		{Title: "Clean Code",
			Author:     "Robert C. Martin",
			IsBorrowed: false},
	}
	for i := 0; i < 2; i++ {
		fmt.Println(library[i].Title)
	}

	callme(library)
	callme(library)
	callme(library)
	// fmt.Println("Enter Name of the Book:")
	// reader := bufio.NewReader(os.Stdin)
	// name, _ := reader.ReadString('\n')
	// name = strings.TrimSpace(name)

	// fmt.Println(BorrowBook(library, name))

	// fmt.Printf("\nLibrary Books\n")
	// for i := 0; i < 2; i++ {
	// 	fmt.Println("Book:", library[i].Title, "Borrowed Status:", library[i].IsBorrowed)
	// }

}
