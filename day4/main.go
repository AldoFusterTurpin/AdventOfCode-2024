package main

import (
	"fmt"
	"os"
)

func main() {
	// inputFile := "inputs/generic.txt"
	// inputFile := "inputs/hola_3.txt"
	// inputFile := "inputs/sample.txt"
	inputFile := "inputs/input.txt"
	// inputFile := "inputs/sample_part_2.txt"
	b, err := os.ReadFile(inputFile)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fileContent := string(b)
	res, err := part2(fileContent)
	// res, err := getResult(fileContent)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("result is: ", res)
}
