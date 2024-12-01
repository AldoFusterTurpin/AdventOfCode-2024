package main

import (
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	inputFile := "inputs/input.txt"
	b, err := os.ReadFile(inputFile)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fileContent := string(b)
	fmt.Println(getResult(fileContent))
}

func getResult(fileContent string) (int, error) {
	lists, err := convertContentToLists(fileContent)
	if err != nil {
		return 0, err
	}
	sortLists(lists)
	result := getTotalDistanceBetweenLists(lists)
	return int(result), nil
}

func convertContentToLists(fileContent string) ([][]int, error) {
	fileContent = strings.Trim(fileContent, " ")
	lines := strings.Split(fileContent, "\n")
	var leftList []int
	var rightList []int

	for _, line := range lines {
		// remove leading and trailing whitespaces
		line = strings.Trim(line, " ")
		if len(line) == 0 {
			continue
		}
		lineSplitted := strings.Fields(line)
		leftNumber, rightNumber := lineSplitted[0], lineSplitted[1]
		leftNumberInt, err := strconv.Atoi(leftNumber)
		if err != nil {
			return nil, err
		}
		leftList = append(leftList, leftNumberInt)

		rightNumberInt, err := strconv.Atoi(rightNumber)
		if err != nil {
			return nil, err
		}
		rightList = append(rightList, rightNumberInt)
	}
	listOfLists := [][]int{leftList, rightList}
	return listOfLists, nil
}

func getTotalDistanceBetweenLists(lists [][]int) float64 {
	result := float64(0)

	// we assume all lists have the same lenght
	lengthOfLists := len(lists[0])
	for i := 0; i < lengthOfLists; i++ {
		leftList := lists[0]
		left := leftList[i]

		rightList := lists[1]
		right := rightList[i]

		result += math.Abs(float64(left) - float64(right))
	}
	return result
}

func sortLists(lists [][]int) {
	for _, list := range lists {
		slices.Sort(list)
	}
}
