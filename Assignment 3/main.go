package main

import (
	"io"
	"log"
	"os"
)

func main() {
	input := os.Args[1]
	inputFile, err := os.Open(input)
	if err != nil {
		log.Fatal(err)
	}
	io.Copy(os.Stdout, inputFile)
}
