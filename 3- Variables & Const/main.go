package main

import (
	"fmt"
)

func main() {
	const myConst = "Hello, Go!"
	fmt.Println("Constant value:", myConst)
	MyLang := "Go"
	fmt.Println(MyLang)
	MyLang = "sds"
	fmt.Println("Updated MyLang:", MyLang)
}