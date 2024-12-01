package main

func part2(fileContent string) (int, error) {
	lists, err := convertContentToLists(fileContent)
	if err != nil {
		return 0, err
	}
	secondList := lists[1]
	ocurrencesOfSecondList := getOcurrencesOfSecondList(secondList)
	firstList := lists[0]
	result := getSimilarityScore(firstList, ocurrencesOfSecondList)
	return int(result), nil
}

func getOcurrencesOfSecondList(list []int) map[int]int {
	m := make(map[int]int, 0)
	for _, v := range list {
		m[v]++
	}
	return m
}

func getSimilarityScore(firstList []int, ocurrencesOfSecondList map[int]int) int {
	similarityScore := 0
	for _, v := range firstList {
		similarityScore += v * ocurrencesOfSecondList[v]
	}
	return similarityScore
}
