package main

func main() {
	paintNeededThree(1.0, 3.0)
}

func paintNeededThree(width float64, height float64) float64 {
	area := width * height
	return area / 10.0
}
