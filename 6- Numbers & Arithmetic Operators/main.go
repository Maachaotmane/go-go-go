package main

import (
	"fmt"
	// "bufio"
	// "os"
	// "strings"
)

func main() {
	//  Numbers & Arithmetic Operators in Go
	// Arithmetic operators are used to perform mathematical operations on numeric values.
	// The main arithmetic operators in Go are:
	// + (addition), - (subtraction), * (multiplication), / (division), % (modulus)
	// Example of arithmetic operators
	A := 10
	B := 3
	fmt.Println("A:", A, "B:", B)
	fmt.Println("Addition:", A+B)       // Addition
	var A1 = A + B                      // Increment A by B or A += B
	fmt.Println("Incremented A1:", A1)  // Display incremented value of A
	fmt.Println("Subtraction:", A-B)    // Subtraction or A -= B
	fmt.Println("Multiplication:", A*B) // Multiplication
	fmt.Println("Division:", A/B)       // Division (integer division)
	fmt.Println("Modulus:", A%B)        // Modulus (remainder)
}
