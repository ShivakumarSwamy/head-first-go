package main

import "fmt"

func main() {
	var amount, total float64
	amount = paintNeededTwo(4.2, 3.0)
	fmt.Printf("%0.2f litres needed\n", amount)
	total += amount
	amount = paintNeededTwo(5.2, 3.5)
	fmt.Printf("%0.2f litres needed\n", amount)
	total += amount
	fmt.Printf("Total: %0.2f litres\n", total)
}

func paintNeededTwo(width float64, height float64) float64 {
	area := width * height
	return area / 10.0
}
