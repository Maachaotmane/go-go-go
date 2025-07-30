package main

import (
	"fmt"
)

func main() {
	// Fmt library in Go
	// The fmt package in Go provides formatted I/O functions similar to C's printf and scanf.
	// It is used for formatted output and input operations.
	// The main functions in the fmt package are:
	// Print, Printf, Println for output
	// Sprintf for formatted string output
	// Scanf for formatted input
	// Example of fmt package usage
	name := "Go"
	age := 10
	fmt.Println("Hello, World!") // Println automatically adds a newline
	fmt.Printf("Welcome to %s programming!\n", name) // Printf allows formatted output
	fmt.Printf("You are %d years old.\n", age) // Printf with integer formatting

	// Using Sprintf to create a formatted string
	formattedString := fmt.Sprintf("Language: %s, Age: %d", name, age)
	fmt.Println("Formatted String:", formattedString) // Print the formatted string

	// Using Scanf to read user input (uncomment to use)
	fmt.Print("Enter your name: ")
	var userName string
	fmt.Scanf("%s", &userName) // Read user input
	fmt.Println("Hello,", userName) // Greet the user
	// Note: Scanf is not commonly used for user input in Go, fmt.Scanln or bufio.Scanner is preferred for more complex input.
	// Example of formatted output with different data types
	floatValue := 3.14
	fmt.Printf("Float value: %.2f\n", floatValue) // Print float with 2 decimal places
	boolValue := true
	fmt.Printf("Boolean value: %t\n", boolValue) // Print boolean value
	intValue := 42
	fmt.Printf("Integer value: %d\n", intValue) // Print integer value
	// Using fmt package for error handling
	err := fmt.Errorf("an error occurred: %s", "something went wrong")
	if err != nil {
		fmt.Println("Error:", err) // Print error message
	}
}
