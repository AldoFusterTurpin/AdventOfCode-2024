package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
  inputFile := "inputs/sample.txt"
	// inputFile := "inputs/input.txt"
	b, err := os.ReadFile(inputFile)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fileContent := string(b)

  res := getNumberOfSafeReports(fileContent)
  fmt.Println("result is: ", res)
}

func convertContentToLists(fileContent string) ([][]int, error) {
	listOfLists := [][]int{}
  fileContent = strings.Trim(fileContent, " ")
	lines := strings.Split(fileContent, "\n")

	for _, line := range lines {
    var report []int
		// remove leading and trailing whitespaces
		line = strings.Trim(line, " ")
		if len(line) == 0 {
			continue
		}
		lineSplitted := strings.Fields(line)
    for _, v := range lineSplitted {
      valueInt, err := strconv.Atoi(v)
      if err != nil {
        return nil, err
      }
      report = append(report, valueInt)
    }
    listOfLists = append(listOfLists, report)
	}
  fmt.Println(listOfLists)
	return listOfLists, nil
}

func getNumberOfSafeReports(input string) int {
  convertContentToLists(input)  

  return 0 
}
