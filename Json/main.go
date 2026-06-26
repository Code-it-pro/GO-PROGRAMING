package main

import (
	"encoding/json"
	"fmt"
)

// person struct represents a person with name, age, and gaming status
// struct tags map Go field names to JSON field names
type person struct {
	Name    string `json:"name"`     // Maps to "name" in JSON
	Age     int    `json:"age"`      // Maps to "age" in JSON
	IsGamer bool   `json:"is_Gamer"` // Maps to "is_Gamer" in JSON
}

// main demonstrates JSON marshaling (Go to JSON) and unmarshaling (JSON to Go)
func main() {
	fmt.Println("Starting with JSON")

	// Create a person struct instance with sample data
	Student := person{
		Name:    "Gagandeep Singh",
		Age:     22,
		IsGamer: true,
	}
	fmt.Println(Student)

	// Marshal: Convert Go struct to JSON byte slice
	JsonData, error := json.Marshal(Student)
	if error != nil {
		fmt.Println(error)
		return
	}
	fmt.Printf("Type of Student: %T\n", JsonData) // Show type is []byte
	fmt.Println("Student json: ", JsonData)       // Display JSON bytes

	// Unmarshal: Convert JSON byte slice back to Go struct
	var NewStudent person
	error1 := json.Unmarshal(JsonData, &NewStudent) // Pass pointer to struct
	if error1 != nil {
		fmt.Println(error1)
		return
	}
	fmt.Printf("Type of JsonData: %T\n", NewStudent) // Show type is person struct
	fmt.Println("New Student : ", NewStudent)        // Display decoded struct
}

/*
EXPLANATION:
============

JSON is a lightweight data format for data exchange. Go's encoding/json package
provides tools to convert between Go structs and JSON:

1. STRUCT DEFINITION:
   - A struct defines the shape of data with fields
   - Struct tags (json:"fieldname") control how fields map to JSON
   - Field names are case-sensitive; unexported fields (lowercase) are ignored

2. MARSHALING (Go → JSON):
   - json.Marshal(Student) converts a Go struct to JSON bytes
   - Returns: []byte containing JSON representation
   - The struct tags determine the JSON field names
   - Example output: {"name":"Gagandeep Singh","age":22,"is_Gamer":true}

3. UNMARSHALING (JSON → Go):
   - json.Unmarshal(JsonData, &NewStudent) converts JSON bytes back to struct
   - Must pass a POINTER (&NewStudent) to the target struct
   - Matches JSON fields to struct fields using struct tags
   - Populates the struct with values from JSON

4. STRUCT TAGS EXPLAINED:
   - `json:"name"` → Maps field Name to JSON key "name"
   - `json:"age"` → Maps field Age to JSON key "age"
   - `json:"is_Gamer"` → Maps field IsGamer to JSON key "is_Gamer"
   - Tags allow flexible mapping between Go naming (PascalCase) and JSON naming (snake_case)

5. ERROR HANDLING:
   - Marshal/Unmarshal return an error if conversion fails
   - Always check for errors before using the result

6. COMMON USE CASES:
   - API requests/responses (client-server communication)
   - File storage and retrieval
   - Configuration files
   - Data serialization for transmission

7. KEY DIFFERENCES:
   - Marshal: struct → []byte (serialization)
   - Unmarshal: []byte → struct (deserialization)
   - Pointers required for Unmarshal to modify struct
   - Round-trip: Go struct → JSON → Go struct (data preserved with struct tags)
*/
