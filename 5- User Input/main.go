package main

import (
	"fmt"
	// "bufio"
	// "os"
	// "strings"
)

func main() {
	// User Input in Go
	// Go provides several ways to read user input from the console.
	// The most common way is to use the `fmt.Scanln` function, which reads input until a newline character is encountered.
	// Here is an example of reading different types of user input:
	var name string
	fmt.Print("Enter your name: ")
	fmt.Scanln(&name)           // Read user input
	fmt.Println("Hello,", name) // Greet the user

	// Read an integer value
	var age int
	fmt.Print("Enter your age: ")
	fmt.Scanln(&age)                         // Read user input
	fmt.Println("You are", age, "years old") // Display age

	// Read a float value
	var height float64
	fmt.Print("Enter your height in meters: ")
	fmt.Scanln(&height)                                // Read user input
	fmt.Printf("Your height is %.2f meters\n", height) // Display height with formatting

	// Read a boolean value
	var isStudent bool
	fmt.Print("Are you a student? (true/false): ")
	fmt.Scanln(&isStudent)                    // Read user input
	fmt.Println("Student status:", isStudent) // Display student status

	// // Read a string slice (array)
	// var hobbies []string
	// fmt.Print("Enter your hobbies (separated by space): ")
	// var input string
	// // Use bufio.Scanner to read the whole line
	// scanner := bufio.NewScanner(os.Stdin)
	// if scanner.Scan() {
	// 	input = scanner.Text()
	// 	hobbies = append(hobbies, strings.Fields(input)...)
	// }
	// fmt.Println("Your hobbies are:", hobbies) // Display hobbies
	// Note: The above code for reading a slice is commented out as it requires additional imports and handling.(Next step would be to handle slices properly)

	// Print a summary of all inputs
	fmt.Println("Summary:")
	fmt.Printf("Name: %s, Age: %d, Height: %.2f, Student: %t\n", name, age, height, isStudent)
	fmt.Println("Thank you for your input!")
	// Note: In a real application, you might want to handle errors and validate inputs.
}
