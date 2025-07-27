package main

import (
	"fmt"
)

func main() {
	// Variables and Constants in Go
	// Variables are used to store data that can change during the execution of a program.
	// Constants are immutable values that cannot be changed after they are defined.
	// They are declared using the `const` keyword.
	// Variables can be declared using the `var` keyword or using shorthand notation.
	// Here is an example of declaring a variable and a constant:
	const myConst = "Hello, Go!"
	fmt.Println("Constant value:", myConst)
	MyLang := "Go"
	fmt.Println(MyLang)
	MyLang = "sds"
	fmt.Println("Updated MyLang:", MyLang)
}
