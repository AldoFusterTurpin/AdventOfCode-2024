package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	// inputFile := "inputs/sample.txt"
	inputFile := "inputs/input.txt"
	b, err := os.ReadFile(inputFile)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fileContent := string(b)

	res, err := getNumberOfSafeReports(fileContent)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("result is: ", res)
}

func convertToReports(fileContent string) ([][]int, error) {
	reports := [][]int{}
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
		reports = append(reports, report)
	}
	// fmt.Println(reports)
	return reports, nil
}

func getNumberOfSafeReports(input string) (int, error) {
	safeReports := 0
	reports, err := convertToReports(input)
	if err != nil {
		return 0, err
	}
	for _, report := range reports {
		if safeReport(report) {
			safeReports++
		}
	}
	return safeReports, nil
}

func safeReport(report []int) bool {
	return (allIncreasing(report) || allDecreasing(report)) && adjecentWithMaxDiff(report, 3)
}

// adjecentWithMaxDiff receives a report and returns true if the difference
// between each number is at least 1 and maxDiff at maximum .
func adjecentWithMaxDiff(report []int, maxDiff int) bool {
	for i := 1; i < len(report); i++ {
		diff := math.Abs(float64(report[i-1]) - float64(report[i]))
		if diff < 1 || diff > float64(maxDiff) {
			return false
		}
	}
  return true
}

func allIncreasing(report []int) bool {
	for i := 1; i < len(report); i++ {
		if report[i-1] >= report[i] {
			return false
		}
	}
	return true
}

func allDecreasing(report []int) bool {
	for i := 1; i < len(report); i++ {
		if report[i-1] <= report[i] {
			return false
		}
	}
	return true
}
