package main

import (
	"bufio"
	"fmt"
	"github.com/ErikPelli/TL-parser-JSON/parser"
	"os"
)

func main() {
	// Get file name as input
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Insert .tl schema file name: ")
	scanner.Scan()
	fileName := scanner.Text()

	// Open file and close at end
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Parse TL to JSON and save it
	err = parser.Parse(file)
	if err != nil {
		panic(err)
	}
}
