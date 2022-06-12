package main

import "fmt"

var packageVar = "package" // scope entire main package

func main() {
	var functionVar = "function" // scope entire function

	if true {
		var conditionalVar = "conditional" // scope if conditional block
		fmt.Println(packageVar)
		fmt.Println(functionVar)
		fmt.Println(conditionalVar)
	}
	fmt.Println(packageVar)
	fmt.Println(functionVar)
	fmt.Println(conditionalVar)
}
