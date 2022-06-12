package main

import (
	"fmt"
	"log"
	"strconv"
)

func main() {
	// only one variable in the short variable declaration has to be new

	a, err := strconv.ParseFloat("1.23", 64) // Declare a and err
	if err != nil {
		log.Fatal(err)
	}
	b, err := strconv.ParseFloat("4.56", 64) // Declare b and assign err
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(a, b)

	var c, d float64
	var er error

	c, er = strconv.ParseFloat("1.23", 64) // Declare a and err
	if er != nil {
		log.Fatal(er)
	}
	d, er = strconv.ParseFloat("4.56", 64) // Declare b and assign err
	if er != nil {
		log.Fatal(er)
	}
	fmt.Println(c, d)
}
