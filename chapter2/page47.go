package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

func main() {

	s := "\t formely surrounded by space \n"
	fmt.Println(strings.TrimSpace(s))

	input := "60"
	grade, err := strconv.ParseFloat(input, 64)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(grade)
}
