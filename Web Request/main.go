package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// main demonstrates making HTTP GET requests and reading the response body
func main() {
	fmt.Println("Starting with Web Requests")

	// Define the URL to fetch
	url := "https://jsonplaceholder.typicode.com/todos/1"

	// Make an HTTP GET request to the specified URL
	res, error := http.Get(url)
	if error != nil {
		fmt.Println("Error getting response: ", error)
	}

	// Ensure the response body is closed after we're done reading
	defer res.Body.Close()

	// Read the entire response body into a byte slice
	data, error := ioutil.ReadAll(res.Body)

	// Print the data type and the actual response content
	fmt.Printf("Response type: %T", data)
	fmt.Println("Response :", string(data))

}
