package main

import (
	"fmt"
	"time"
)

func main() {
	// This is a simple Go program that prints "Hello, World!" and waits for 10 seconds before printing "Welcome to Go programming!"
	// It demonstrates the use of variables, string concatenation, and the time package for delays
	MyLang := "Go"
	fmt.Println("Hello, World!")
	time.Sleep(10 * time.Second)
	fmt.Println("Welcome to " + MyLang + " programming!")

}
