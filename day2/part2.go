package main

func getNumberOfSafeReportsPart2(input string) (int, error) {
	safeReports := 0
	reports, err := convertToReports(input)
	if err != nil {
		return 0, err
	}

	for _, report := range reports {
		if safeReportPart2(report) {
			safeReports++
		}
	}
	return safeReports, nil
}

// getAllPermutationsOfReport returns a slice of copies of the original
// where the element at index i was deleted for i >= 0 ... i < len(original).
// Example:
// original: 1, 2, 3, 4
//
//	returns {
//			2, 3, 4 -> "1" was removed
//			1, 3, 4 -> "2" was removed
//			1, 2, 4 -> "3" was removed
//			1, 2, 3 -> "4" was deleted
//	}
func getAllPermutationsOfReport(original []int) [][]int {
	allPermutations := make([][]int, 0)
	for i := range original {
		sliceCopy := getCopyWithouthElementAtIndex(original, i)
		allPermutations = append(allPermutations, sliceCopy)
	}
	return allPermutations
}

func safeReportPart2(report []int) bool {
	allPermutations := getAllPermutationsOfReport(report)
	for _, v := range allPermutations {
		if safeReport(v) {
			return true
		}
	}
	return false
}

// getCopyWithouthElementAtIndex returns a deep copy of original but removing
// original[indexToExclude]. If indexToExclude is outbounds of original,
// it will return a exact copy from original.
func getCopyWithouthElementAtIndex(original []int, indexToExclude int) []int {
	if indexToExclude < 0 || indexToExclude >= len(original) {
		copySlice := make([]int, len(original))
		copy(copySlice, original)
		return copySlice
	}
	// we want a deep copy -> new slice initialized
	newReport := []int{}

	// add left part
	newReport = append(newReport, original[:indexToExclude]...)

	// add right part
	newReport = append(newReport, original[indexToExclude+1:]...)
	return newReport
}
