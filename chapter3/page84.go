package main

import "fmt"

func main() {

	// 12.3456 --> 7 characters
	// %7.3f -->  (whitespace)12.346 --> 7 as minimum width (includes decimal point and decimal places)
	fmt.Printf("%%7.3f: %7.3f\n", 12.3456) // Rounded to three places
	fmt.Printf("%%7.2f: %7.2f\n", 12.3456) // Rounded to two places
	fmt.Printf("%%7.1f: %7.1f\n", 12.3456) // Rounded to one place
	fmt.Printf("%%.1f: %.1f\n", 12.3456)   // Rounded to one place, no padding
	fmt.Printf("%%.2f: %.2f\n", 12.3456)   // Rounded to two places, no padding

	fmt.Printf("%.2f\n", 1.260000000000)
	fmt.Printf("%.2f\n", 1.819999999999)
}
