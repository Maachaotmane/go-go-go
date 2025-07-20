package main

import (
	"fmt"
)

func main() {
	// Defining a constant
	const myConst = "Hello, Go!" // Or const myConst string = "Hello, Go!"
	fmt.Println("Constant value:", myConst)

	// We can't change the value of a constant
	// myConst = "New Value" // This will cause a compile-time error

	// However, we can define a variable and change its value
	MyLang := "Go" // Using short variable declaration
	fmt.Println("Initial MyLang:", MyLang)
	MyLang = "Updated Go" // Changing the value of MyLang
	fmt.Println("Updated MyLang:", MyLang)
}
