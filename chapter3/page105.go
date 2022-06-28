package main

import "fmt"

func main() {
	var myInt int = 4
	var myIntPointer *int = &myInt
	fmt.Println(myIntPointer)  // address
	fmt.Println(*myIntPointer) // value at address
	fmt.Println("--------------------")

	var myInt2 int = 4
	var myInt2Pointer *int = &myInt2
	*myInt2Pointer = 8          // change value at address
	fmt.Println(myInt2Pointer)  // address
	fmt.Println(&myInt2)        // address
	fmt.Println(*myInt2Pointer) // value at address
	fmt.Println(myInt2)         // value same as myInt2Pointer as line 13 changed value at address
}
