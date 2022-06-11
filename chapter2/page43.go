package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	fileInfo, err := os.Stat("my.txt") // When running in IDE change working directory to this file directory
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(fileInfo.Size())
}
