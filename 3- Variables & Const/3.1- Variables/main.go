package main

import (
	"fmt"
)

func main() {
	// Variable Declaration and Initialization
	// In Go, variables can be declared using the `var` keyword or using shorthand notation
	// Variables can be of different types, such as string, int, float, bool, etc.
	// Here are some examples of variable declaration and initialization:
	// Using var keyword
	var myVar string = "Hello, Go!"
	fmt.Println("Variable value:", myVar)

	var myvar2 string
	myvar2 = "Welcome to Go programming."
	fmt.Println("Variable myvar2:", myvar2)

	// concatenation
	myVar = myVar + " " + myvar2 // or myVar += " " + myvar2 or myVar = fmt.Sprintf("%s %s", myVar, myvar2)
	fmt.Println("Concatenated Variable:", myVar)

	myVar3 := "This is a shorthand declaration."
	fmt.Println("Variable myVar3:", myVar3)

	ahmed, ali := "Ahmed", "Ali" // := shorthand for variable declaration and assignment
	fmt.Println("Variables ahmed and ali:", ahmed, ali)

	var Mere, Pere = "Mere", "Pere"
	fmt.Println("Variables Mere and Pere:", Mere, Pere)

	var a, b, c = 1, 2, 3
	fmt.Println("Variables a, b, c:", a, b, c)
	// or
	var d, e, f int
	d, e, f = 1, 2, 3
	fmt.Println("Variables a, b, c after assignment:", d, e, f)
	//or
	myNumber := 100
	fmt.Println("Variable myNumber:", myNumber)

	var myFloat float64 = 3.14 // Using float64 for precision or float32 for less precision
	fmt.Println("Variable myFloat:", myFloat)

	var myBool bool = true // or false or myBool := true
	fmt.Println("Variable myBool:", myBool)

	var mySlice []string = []string{"Go", "is", "fun"} // or var mySlice = []string{"Go", "is", "fun"} or mySlice := []string{"Go", "is", "fun"}
	fmt.Println("Variable mySlice:", mySlice)
}
