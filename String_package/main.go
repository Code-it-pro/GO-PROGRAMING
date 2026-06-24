package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("String package")

	// Example 1: Print a simple string
	name := "Gagandeep"
	fmt.Println(name)

	// Example 2: Use strings.Split() to split a string by delimiter
	var array = "Apple,Banana,Orange"
	fmt.Println(array)
	// Split the string by comma and store it in a slice
	splited := strings.Split(array, ",")
	fmt.Println(splited)
	// Print the type of the split result (slice of strings)
	fmt.Printf("Type of Spilt : %T\n", splited)

	// Example 3: Use strings.TrimSpace() to remove leading and trailing whitespace
	var str string = "          | Hello WOrld |        "
	fmt.Println(str)
	// TrimSpace removes whitespace from both ends of the string
	trimmed := strings.TrimSpace(str)
	fmt.Println("Trimmed :", trimmed)

	// Example 4: Use strings.Count() to count occurrences of a substring
	str2 := "One two three four two two two five"
	find := "two"
	// Count how many times "two" appears in the string
	count := strings.Count(str2, find)
	fmt.Printf("Count of %s : %d\n", find, count)

	// Example 5: Use strings.Join() to join multiple strings with a separator
	str_firstnam := "Gagandeep"
	str_lastname := "Singh"
	// Join first name and last name with a space separator
	full_name := strings.Join([]string{str_firstnam, str_lastname}, " ")
	fmt.Println("Full Name :", full_name)
}

// ================================================================================
// STRING PACKAGE - IMPORTANT FUNCTIONS DOCUMENTATION
// ================================================================================

// strings.Split()
// What it does: Splits a string into a slice of substrings based on a delimiter.
// Inputs: Takes a string and a delimiter string. Returns a slice of strings.
// Example:
// result := strings.Split("Apple,Banana,Orange", ",")
// // result = ["Apple", "Banana", "Orange"]

// strings.Join()
// What it does: Joins a slice of strings into a single string with a separator.
// Inputs: Takes a slice of strings and a separator string. Returns a single string.
// Example:
// result := strings.Join([]string{"Hello", "World"}, " ")
// // result = "Hello World"

// strings.TrimSpace()
// What it does: Removes leading and trailing whitespace from a string.
// Inputs: Takes a string. Returns a string with whitespace removed from both ends.
// Example:
// result := strings.TrimSpace("   Hello World   ")
// // result = "Hello World"

// strings.Count()
// What it does: Counts the number of occurrences of a substring in a string.
// Inputs: Takes a string and a substring to search for. Returns an integer count.
// Example:
// result := strings.Count("one two two three two", "two")
// // result = 3

// strings.Contains()
// What it does: Checks if a string contains a substring.
// Inputs: Takes a string and a substring. Returns a boolean (true/false).
// Example:
// result := strings.Contains("Hello World", "World")
// // result = true

// strings.ToUpper()
// What it does: Converts all characters in a string to uppercase.
// Inputs: Takes a string. Returns a string in uppercase.
// Example:
// result := strings.ToUpper("hello")
// // result = "HELLO"

// strings.ToLower()
// What it does: Converts all characters in a string to lowercase.
// Inputs: Takes a string. Returns a string in lowercase.
// Example:
// result := strings.ToLower("HELLO")
// // result = "hello"

// strings.Replace()
// What it does: Replaces occurrences of a substring with another substring.
// Inputs: Takes a string, old substring, new substring, and count (-1 for all). Returns the modified string.
// Example:
// result := strings.Replace("hello hello hello", "hello", "hi", 2)
// // result = "hi hi hello"

// strings.HasPrefix()
// What it does: Checks if a string starts with a specified prefix.
// Inputs: Takes a string and a prefix string. Returns a boolean.
// Example:
// result := strings.HasPrefix("Hello World", "Hello")
// // result = true

// strings.HasSuffix()
// What it does: Checks if a string ends with a specified suffix.
// Inputs: Takes a string and a suffix string. Returns a boolean.
// Example:
// result := strings.HasSuffix("Hello World", "World")
// // result = true

// strings.Index()
// What it does: Returns the index of the first occurrence of a substring.
// Inputs: Takes a string and a substring. Returns an integer index (-1 if not found).
// Example:
// result := strings.Index("Hello World", "World")
// // result = 6

// strings.Fields()
// What it does: Splits a string into a slice of words (separated by whitespace).
// Inputs: Takes a string. Returns a slice of strings.
// Example:
// result := strings.Fields("Hello   World   Go")
// // result = ["Hello", "World", "Go"]
