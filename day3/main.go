package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
)

// Check regex.md for explanations about everything regex related.
var regexStr = `mul\((\d{1,3}),(\d{1,3})\)`

var r = regexp.MustCompile(regexStr)

func main() {
	// inputFile := "inputs/sample.txt"
	inputFile := "inputs/input.txt"
	b, err := os.ReadFile(inputFile)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fileContent := string(b)
	res, err := getResult(fileContent)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("result is: ", res)
}

func getResult(input string) (int, error) {
	total := 0
	s := r.FindAllStringSubmatch(input, -1)
	for _, match := range s {
		// match is like
		// [ "mul(2,4)", "2", "4" ]
		leftNumber, err := strconv.Atoi(match[1])
		if err != nil {
			return 0, err
		}

		rightNumber, err := strconv.Atoi(match[2])
		if err != nil {
			return 0, err
		}
		total += leftNumber * rightNumber
	}
	return total, nil
}
