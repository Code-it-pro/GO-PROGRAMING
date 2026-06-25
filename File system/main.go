package main

func main() {
	// // Print a startup message to the console.
	// fmt.Println("Starting with file management")

	// // Create (or truncate) the file at File system/example.txt.
	// file, error := os.Create("File system/example.txt")
	// if error != nil {
	// 	fmt.Println("encountered an error", error)
	// }

	// fmt.Println("File created successfully")
	// // Ensure the file is closed when main returns.
	// defer file.Close()

	// content := "Hello World!"
	// // Write the string to the file. n receives the number of bytes written.
	// n, error1 := io.WriteString(file, content)
	// fmt.Println(error1)
	// fmt.Println(n)

	// Reading files 2 Methods
	//Method 1

	// file, error := os.Open("File system/example.txt")
	// if error != nil {
	// 	fmt.Println("encountered an error", error)
	// }

	// buffer := make([]byte, 1024)

	// for {
	// 	n, error := file.Read(buffer)
	// 	if error == io.EOF {
	// 		break
	// 	}
	// 	if error != nil {
	// 		fmt.Println("Error Reading the file :", error)
	// 	}

	// 	fmt.Println(string(buffer[:n]))
	// }
	// defer file.Close()

	// Method 2

	// content, error := ioutil.ReadFile("File system/example.txt") // instead of ioutil you can also use io package
	// if error != nil {
	// 	fmt.Println("Error Reading file: ", error)
	// 	return
	// }
	// fmt.Println(string(content))

}
