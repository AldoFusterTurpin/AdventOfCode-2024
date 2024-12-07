package main

import (
	"fmt"
	"math"
)

func getNumberOfSafeReportsPart2(input string, toleration int) (int, error) {
	safeReports := 0
	reports, err := convertToReports(input)
	if err != nil {
		return 0, err
	}
	for _, report := range reports {
		fmt.Printf("Analizing %v\n", report)
		if safeReportPart2(report, toleration) {
			fmt.Println("safe")
			safeReports++
		} else {
			fmt.Println("unsafe")
		}
		fmt.Println()
	}
	return safeReports, nil
}

func safeReportPart2(report []int, toleration int) bool {
	badIncreasingLevels := allIncreasingWithMaxToleration(report, toleration)
	fmt.Printf("badIncreasingLevels: %v\n", badIncreasingLevels)
	areAllIncreasingWithOneToleration := len(badIncreasingLevels) < 2

	badDecreasingLevels := allDecreasingWithMaxToleration(report, toleration)
	fmt.Printf("badDecreasingLevels: %v\n", badDecreasingLevels)
	areallDecreasingWithOneToleration := len(badDecreasingLevels) < 2

	badAdjecentLevels := adjecentWithMaxDiffAndToleration(report, 3, toleration)
	fmt.Printf("badAdjecentLevels: %v\n", badAdjecentLevels)

	// The levels are either all increasing or all decreasing.
	if !areAllIncreasingWithOneToleration || !areallDecreasingWithOneToleration {
		return false
	}
	// Any two adjacent levels differ by at least one and at most three with 1 toleration.
	if len(badAdjecentLevels) > 1 {
		return false
	}

	return allBadLevelsAreTheSame(badIncreasingLevels, badDecreasingLevels, badAdjecentLevels)
}

func allBadLevelsAreTheSame(badIncreasingLevels, badDecreasingLevels, badAdjecentLevels []int) bool {
	nBadIncreasingLevels := len(badIncreasingLevels)
	nBadDecreasingLevels := len(badDecreasingLevels)
	nBadAdjecentLevels := len(badAdjecentLevels)
	sameNumberOfWrongLevels := nBadIncreasingLevels == nBadDecreasingLevels && nBadDecreasingLevels == nBadAdjecentLevels

	if sameNumberOfWrongLevels {
		if nBadIncreasingLevels == 1 {
			return true
		}
		return badIncreasingLevels[0] == badDecreasingLevels[0] && badDecreasingLevels[0] == badAdjecentLevels[0]
	}
	return false
}

func adjecentWithMaxDiffAndToleration(report []int, maxDiff int, maxToleration int) (indexOfBadLevels []int) {
	for i := 1; i < len(report); i++ {
		diff := math.Abs(float64(report[i-1]) - float64(report[i]))
		if diff < 1 || diff > float64(maxDiff) {
			indexOfBadLevels = append(indexOfBadLevels, i)
		}
	}
	return indexOfBadLevels
}

func allIncreasingWithMaxToleration(report []int, maxToleration int) (indexOfBadLevels []int) {
	for i := 1; i < len(report); i++ {
		if report[i-1] >= report[i] {
			indexOfBadLevels = append(indexOfBadLevels, i)
		}
	}
	return indexOfBadLevels
}

func allDecreasingWithMaxToleration(report []int, maxToleration int) (indexOfBadLevels []int) {
	for i := 1; i < len(report); i++ {
		if report[i-1] <= report[i] {
			indexOfBadLevels = append(indexOfBadLevels, i)
		}
	}
	return indexOfBadLevels
}
