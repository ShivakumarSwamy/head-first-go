package main

import "fmt"

func main() {

	if true {
		fmt.Println("I'll be printed")
	}

	if false {
		fmt.Println("I won't")
	}

	grade := 100

	if grade == 100 {
		fmt.Println("Perfect!")
	} else if grade >= 60 {
		fmt.Println("You pass.")
	} else {
		fmt.Println("You fail!")
	}

}
