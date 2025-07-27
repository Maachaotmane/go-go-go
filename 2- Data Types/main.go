package main

import (
	"fmt"
)

func main() {
	// Data Types in Go
	// Go has several built-in data types, including:
	// 1. Boolean: Represents true or false values.
	// 2. Integer: Represents whole numbers, both positive and negative.
	// 3. Float: Represents decimal numbers.
	// 4. String: Represents a sequence of characters.
	// 5. Slice: Represents a dynamically-sized array.
	// Each data type has its own characteristics and use cases.
	// Here are examples of each data type:
	isTrue := true
	isFalse := false
	myBool := 0
	fmt.Println("Boolean True:", isTrue)
	fmt.Printf("Type %T\n", isTrue)
	fmt.Println("Boolean False:", isFalse)
	fmt.Printf("%T\n", isFalse)
	fmt.Println("Integer representation of True:", myBool)
	fmt.Printf("%T\n", myBool)

	// Integer variable
	myInt := 42 * 3
	fmt.Printf("%v, %T\n", myInt, myInt)

	// Float variable
	myFloat := 3.14
	fmt.Printf("%v, %T\n", myFloat, myFloat)

	// String variable
	myString := "Hello, Go!"
	fmt.Printf("%v, %T\n", myString, myString)
	// string with concat
	myString = myString + " Welcome to Go programming."
	fmt.Printf("%v, %T\n", myString, myString)

	// []string variable or slice
	mySlice := []string{"Go", "is", "fun"}
	fmt.Printf("%v, %T\n", mySlice, mySlice)
}
