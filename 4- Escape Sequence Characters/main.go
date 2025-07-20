package main

import (
	"fmt"
)

func main() {
	// Escape sequences characters in Go
	fmt.Println("Hello, World!") // Normal string
	fmt.Println("Hello,\nWorld!") // New line
	fmt.Println("Hello, \bWorld!") // Backspace
	fmt.Println("Hello, \rWorld!") // Carriage return
	fmt.Println("Hello,\tWorld!") // Horizontal tab
	fmt.Println("Hello, \"World\"!") // Double quotes
	fmt.Println("Hello, 'World'!") // Single quotes
	fmt.Println("Hello, \\World!") // Backslash
	fmt.Println("Hello, \vWorld!") // Vertical tab
	fmt.Println("Hello, \aWorld!") // Alert (bell)
	fmt.Println("Hello, \fWorld!") // Form feed
	fmt.Println("Hello, \x1B[31mWorld!\x1B[0m") // Colored text (red)

}