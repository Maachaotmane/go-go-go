package main

import (
	"fmt"
	"strconv" // Import strconv for string conversion
)

func main() {
	// Strings Concatenation in Go
	// String concatenation is the process of joining two or more strings together.
	// In Go, you can concatenate strings using the `+` operator or the `fmt.Sprintf` function.
	// Here are some examples of string concatenation:
	str1 := "Hello"
	str2 := "World"
	// Using the + operator
	concatenated := str1 + " " + str2
	fmt.Println("Concatenated using + operator:", concatenated)

	// Using fmt.Sprintf
	concatenatedWithSprintf := fmt.Sprintf("%s %s", str1, str2)
	fmt.Println("Concatenated using fmt.Sprintf:", concatenatedWithSprintf)

	// You can also concatenate strings with variables and constants
	const greeting = "Greetings"
	name := "Alice"
	concatenatedWithVar := greeting + ", " + name + "!"
	fmt.Println("Concatenated with variable and constant:", concatenatedWithVar)

	// Concatenating multiple strings
	multipleStrings := "Go" + " " + "is" + " " + "fun"
	fmt.Println("Concatenated multiple strings:", multipleStrings)

	// Concatenating strings with numbers (converted to string)
	num := 42
	concatenatedWithNumber := str1 + " " + fmt.Sprintf("%d", num) // or
	// using strconv.Itoa to convert int to string
	concatenatedWithNumber2 := str1 + " " + strconv.Itoa(num)
	fmt.Println("Concatenated with number:", concatenatedWithNumber)
	fmt.Println("Concatenated with number using strconv:", concatenatedWithNumber2)

	// Concatenating strings with special characters
	specialChar := "!"
	concatenatedWithSpecialChar := str1 + " " + str2 + specialChar
	fmt.Println("Concatenated with special character:", concatenatedWithSpecialChar)

	// Concatenating strings with a slice
	strSlice := []string{"Go", "is", "awesome"}
	concatenatedSlice := ""
	for _, str := range strSlice {
		concatenatedSlice += str + " "
	}
	fmt.Println("Concatenated slice:", concatenatedSlice)

}
