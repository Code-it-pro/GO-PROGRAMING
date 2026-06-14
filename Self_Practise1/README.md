# Go Practice Exercises

### 1. The Inventory Manager (Maps & Loops)
Create a program that manages a small grocery store inventory.
- Define a map where the key is the item name (string) and the value is the quantity (int).
- Add at least 5 items to the map.
- Use a `for range` loop to print all items that have a quantity less than 10 (low stock warning).
- Practice deleting an item from the map and then print the updated length of the inventory.

### 2. The Calculator (Switch Case & I/O)
Build a basic CLI calculator.
- Ask the user to input two numbers.
- Ask the user to input an operator (`+`, `-`, `*`, `/`).
- Use a `switch` statement on the operator to perform the calculation and print the result.
- Handle the case where the user enters an invalid operator using the `default` case.

### 3. Even-Odd Filter (For Loops & Slices)
Given a slice of integers: `numbers := []int{11, 22, 33, 44, 55, 66, 77, 88, 99}`
- Iterate through the slice using a `for range` loop.
- Use an `if` statement to check if a number is even or odd.
- Print "Even: X" or "Odd: X" for each number.

### 4. User Profile Builder (bufio & Formatting)
Expand on your `Reading/main.go` logic.
- Ask the user for their First Name, Last Name, and City.
- Use `bufio.NewReader` to handle spaces in city names (e.g., "New York").
- Use `fmt.Printf` to print a formatted "Mailing Label" like:
  ```
  Recipient: [LAST NAME], [FIRST NAME]
  Location: [CITY]
  ```