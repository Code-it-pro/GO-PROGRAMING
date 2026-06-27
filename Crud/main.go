package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

// Todos represents a Todo item from the JSONPlaceholder API
type Todos struct {
	UserID    int    `json:"userId"`    // User ID associated with the todo
	ID        int    `json:"id"`        // Unique identifier for the todo
	Title     string `json:"title"`     // Title/description of the todo
	Completed bool   `json:"completed"` // Whether the todo is completed
}

// GetResponseAPI performs a GET request to the given URL and decodes the JSON response
func GetResponseAPI(url string) {
	// Make HTTP GET request
	res, err := http.Get(url)
	if err != nil {
		fmt.Println("Error occured ", err)
		return
	}
	// Ensure response body is closed after function returns
	defer res.Body.Close()

	// Check if the HTTP status code is OK (200)
	if res.StatusCode != http.StatusOK {
		fmt.Println("Response error", res.Status)
	}
	fmt.Println(res.Status)

	// Decode JSON response into Todos struct
	var result Todos
	error3 := json.NewDecoder(res.Body).Decode(&result)
	if error3 != nil {
		fmt.Println("Error during Decode", error3)
		return
	}
	// Print the decoded todo data
	fmt.Println(result)
}

// PostRequestAPI sends a POST request with JSON data to the given URL
func PostRequestAPI(url, content string) {
	// Convert string content to io.Reader for HTTP body
	jsonReader := strings.NewReader(content)

	// Send POST request with JSON content type
	res, error5 := http.Post(url, "application/json", jsonReader)
	if error5 != nil {
		fmt.Println(error5)
		return
	}
	fmt.Println(res.Status)

	// Ensure response body is closed after function returns
	defer res.Body.Close()

	// Read the response body
	dataresponse, error7 := ioutil.ReadAll(res.Body)
	if error7 != nil {
		fmt.Println("Error 7: ", error7)
		return
	}
	// Print the response as string
	fmt.Println(string(dataresponse))
}

// UpdateRequest sends a PUT request to update an existing resource at the given URL
func UpdateRequest(url, content string) {
	// Convert string content to io.Reader for HTTP body
	jsonReader := strings.NewReader(content)

	// Create a new HTTP request with PUT method and JSON body
	req, error8 := http.NewRequest(http.MethodPut, url, jsonReader)
	if error8 != nil {
		fmt.Println("Error 8: ", error8)
		return
	}

	// Set the Content-Type header to application/json
	req.Header.Set("Content-type", "application/json")

	// Ensure request body is closed after function returns
	defer req.Body.Close()

	// Create an HTTP client and execute the request
	Client := http.Client{}
	res, error9 := Client.Do(req)

	if error9 != nil {
		fmt.Println("Error 9: ", error9)
		return
	}
	// Ensure response body is closed after function returns
	defer res.Body.Close()

	// Read the entire response body
	Updateresponse, error10 := ioutil.ReadAll(res.Body)
	if error10 != nil {
		fmt.Println("Error 10: ", error10)
		return
	}
	// Print the response as string
	fmt.Println(string(Updateresponse))
	// Print the HTTP status code
	fmt.Println(res.Status)
}

// DeleteServerEntry sends a DELETE request to remove a resource at the given URL
func DeleteServerEntry(url string) {
	// Create a new HTTP request with DELETE method and no body
	req, error11 := http.NewRequest(http.MethodDelete, url, nil)
	if error11 != nil {
		fmt.Println("Error 10: ", error11)
		return
	}

	// Create an HTTP client and execute the DELETE request
	Client := http.Client{}
	res, error12 := Client.Do(req)
	if error12 != nil {
		fmt.Println("Error 12: ", error12)
		return
	}
	// Ensure response body is closed after function returns
	defer res.Body.Close()

	// Print the HTTP status code
	fmt.Println(res.Status)

}

func main() {
	fmt.Println("Starting With CRUD")

	// Define the API endpoint URL for GET request
	link1 := "https://jsonplaceholder.typicode.com/todos/1"
	fmt.Println("\nGetting....")
	// Fetch and decode a single todo item
	GetResponseAPI(link1)

	// Create a new Todo struct with sample data
	fmt.Println("\nPosting....")
	link2 := "https://jsonplaceholder.typicode.com/todos"
	rough := Todos{
		UserID:    533,
		Title:     "Gagandeep Singh",
		Completed: true,
	}
	fmt.Println(rough)

	// Convert Todos struct to JSON format
	jsonData, error6 := json.Marshal(rough)
	if error6 != nil {
		fmt.Println("Error 6: ", error6)
		return
	}

	// Send POST request with the JSON data to the API
	PostRequestAPI(link2, string(jsonData))

	fmt.Println("\nUpdating....")

	// Define the API endpoint URL for PUT request to update todo with ID 2
	link3 := "https://jsonplaceholder.typicode.com/todos/2"

	// Create a new Todo struct with updated data
	rough1 := Todos{
		UserID:    5323523,
		Title:     "Gagandeep Singh Gaming Community",
		Completed: true,
	}

	// Convert Todos struct to JSON format for updating
	jsonData1, error6 := json.Marshal(rough1)
	if error6 != nil {
		fmt.Println("Error 6: ", error6)
		return
	}

	// Send PUT request with the updated JSON data to the API
	UpdateRequest(link3, string(jsonData1))

	// Delete an existing todo from the API
	fmt.Println("\nDeleting....")

	// Define the API endpoint URL for DELETE request to remove todo with ID 5
	link4 := "https://jsonplaceholder.typicode.com/todos/5"
	// Send DELETE request to remove the resource
	DeleteServerEntry(link4)

}

/*
================================================================================
                    CRUD OPERATIONS ALGORITHMS
================================================================================

1. GET METHOD (Read Operation)
   Algorithm: GetResponseAPI(url string)
   ─────────────────────────────────────────────────────────────────────────
   INPUT: url (API endpoint URL)
   OUTPUT: Prints the fetched and decoded Todo data

   STEPS:
   1. Send HTTP GET request to the provided URL
   2. Check if request completed without errors
   3. Verify HTTP status code is 200 (OK)
   4. Create a Todos struct to store decoded data
   5. Use JSON decoder to parse response body into Todos struct
   6. Handle any decoding errors
   7. Print the fetched Todo data to console
   8. Close response body (via defer)

   COMPLEXITY: O(n) - where n is the size of response body

   EXAMPLE:
   GetResponseAPI("https://jsonplaceholder.typicode.com/todos/1")
   // Returns: {UserID: 1, ID: 1, Title: "delectus aut autem", Completed: false}


2. POST METHOD (Create Operation)
   Algorithm: PostRequestAPI(url, content string)
   ─────────────────────────────────────────────────────────────────────────
   INPUT: url (API endpoint), content (JSON string of Todo data)
   OUTPUT: Prints the API response

   STEPS:
   1. Convert JSON string content to io.Reader using strings.NewReader
   2. Send HTTP POST request with content-type "application/json"
   3. Handle request errors
   4. Check HTTP response status
   5. Read entire response body into byte slice
   6. Handle any read errors
   7. Convert response bytes to string and print
   8. Close response body (via defer)

   COMPLEXITY: O(n) - where n is the size of request/response body

   EXAMPLE:
   data := Todos{UserID: 533, Title: "New Todo", Completed: true}
   jsonStr := JSON.Marshal(data)
   PostRequestAPI("https://jsonplaceholder.typicode.com/todos", jsonStr)
   // Creates a new resource and returns the created data


3. PUT METHOD (Update Operation)
   Algorithm: UpdateRequest(url, content string)
   ─────────────────────────────────────────────────────────────────────────
   INPUT: url (API endpoint with resource ID), content (Updated JSON data)
   OUTPUT: Prints the API response and status

   STEPS:
   1. Convert JSON string content to io.Reader using strings.NewReader
   2. Create a new HTTP request with PUT method
   3. Handle request creation errors
   4. Set "Content-Type" header to "application/json"
   5. Create an HTTP client instance
   6. Execute the request using Client.Do()
   7. Handle execution errors
   8. Read entire response body into byte slice
   9. Handle read errors
   10. Print response as string and HTTP status code
   11. Close response body (via defer)

   COMPLEXITY: O(n) - where n is the size of request/response body

   EXAMPLE:
   data := Todos{UserID: 533, Title: "Updated Todo", Completed: false}
   jsonStr := JSON.Marshal(data)
   UpdateRequest("https://jsonplaceholder.typicode.com/todos/2", jsonStr)
   // Updates resource with ID 2 and returns updated data


4. DELETE METHOD (Delete Operation)
   Algorithm: DeleteServerEntry(url string)
   ─────────────────────────────────────────────────────────────────────────
   INPUT: url (API endpoint with resource ID to delete)
   OUTPUT: Prints the HTTP status code

   STEPS:
   1. Create a new HTTP request with DELETE method
   2. Set the body to nil (DELETE requests typically have no body)
   3. Handle request creation errors
   4. Create an HTTP client instance
   5. Execute the request using Client.Do()
   6. Handle execution errors
   7. Print the HTTP status code (204 No Content indicates successful deletion)
   8. Close response body (via defer)

   COMPLEXITY: O(1) - constant time operation, no body to process

   EXAMPLE:
   DeleteServerEntry("https://jsonplaceholder.typicode.com/todos/5")
   // Deletes resource with ID 5 from the server


================================================================================
                         PROGRAM WORKFLOW
================================================================================

main() execution order:
   1. Display "Starting With CRUD" message
   2. Call GetResponseAPI() - Fetch existing todo from API (READ)
   3. Create Todos object with sample data
   4. Marshal Todos object to JSON string
   5. Call PostRequestAPI() - Send new todo to API (CREATE)
   6. Create new Todos object with different data
   7. Marshal new Todos object to JSON string
   8. Call UpdateRequest() - Update existing todo in API (UPDATE)
   9. Call DeleteServerEntry() - Remove todo from API (DELETE)


================================================================================
                      HTTP STATUS CODES REFERENCE
================================================================================

- 200 OK: Request successful
- 201 Created: Resource created successfully (POST)
- 400 Bad Request: Invalid request format
- 404 Not Found: Resource not found
- 500 Server Error: Server encountered error

================================================================================
*/
