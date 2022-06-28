package main

import (
	"fmt"
	"reflect"
)

func main() {

	var myInt int
	fmt.Println(reflect.TypeOf(myInt)) // *int

	var myFloat float64
	fmt.Println(reflect.TypeOf(myFloat)) // *float

	var myBool bool
	fmt.Println(reflect.TypeOf(myBool)) // *bool

	var myInt2 int
	var myInt2Pointer *int
	myInt2Pointer = &myInt2
	fmt.Println(myInt2Pointer)

	var myBool2 bool
	myBoolPointer := &myBool2
	fmt.Println(myBoolPointer)
}
