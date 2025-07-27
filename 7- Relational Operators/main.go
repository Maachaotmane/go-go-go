package main

import (
	"fmt"
	// "bufio"
	// "os"
	// "strings"
)

func main() {
	//  Relational Operators
	// Relational operators are used to compare two values and return a boolean result.
	// They are commonly used in conditional statements and loops.
	// The main relational operators in Go are:
	// == (equal to), != (not equal to), > (greater than), < (less than), >= (greater than or equal to), <= (less than or equal to)
	// Example of relational operators
	A := 10
	B := 20
	fmt.Println("A:", A, "B:", B)                          // Display values of A and B
	fmt.Println("A == B:", A == B)                         // Equal to
	fmt.Println("A != B:", A != B)                         // Not equal to
	fmt.Println("A > B:", A > B)                           // Greater than
	fmt.Println("A < B:", A < B)                           // Less than
	fmt.Println("A >= B:", A >= B)                         // Greater than or equal to
	fmt.Println("A <= B:", A <= B)                         // Less than or equal to
	fmt.Println("A == 10:", A == 10)                       // Check if A is equal to 10
	fmt.Println("B != 20:", B != 20)                       // Check if B is not equal to 20
	fmt.Println("A > 5:", A > 5)                           // Check if A is greater than 5
	fmt.Println("B < 30:", B < 30)                         // Check if B is less than 30
	fmt.Println("A >= 10:", A >= 10)                       // Check if A is greater than or equal to 10
	fmt.Println("B <= 20:", B <= 20)                       // Check if B is less
	fmt.Println("A == B && A < 15:", A == B && A < 15)     // Logical AND
	fmt.Println("A != B || B > 15:", A != B || B > 15)     // Logical OR
	fmt.Println("!(A < B):", !(A < B))                     // Logical NOT
	fmt.Println("A == 10 && B > 15:", A == 10 && B > 15)   // Check if A is equal to 10 and B is greater than 15
	fmt.Println("A != 10 || B < 25:", A != 10 || B < 25)   // Check if A is not equal to 10 or B is less than 25
	fmt.Println("!(A > 5):", !(A > 5))                     // Check if A is not greater than 5
	fmt.Println("A == 10 && B <= 20:", A == 10 && B <= 20) // Check if A is equal to 10 and B is less than or equal to 20
	fmt.Println("A != 10 || B >= 20:", A != 10 || B >= 20) // Check if A is not equal to 10 or B is greater than or equal to 20
}
