package main

import (
	"fmt"
)

func main() {
	// Boolean variable True and False 1 and 0
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