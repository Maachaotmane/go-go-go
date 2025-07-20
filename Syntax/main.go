package main

import (
	"fmt"
	"time"
)

func main() {
	MyLang := "Go"
	fmt.Println("Hello, World!")
	time.Sleep(10 * time.Second)
	fmt.Println("Welcome to " + MyLang + " programming!")
}