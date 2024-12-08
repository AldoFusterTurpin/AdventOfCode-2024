package main

import (
	"regexp"
	"strings"
)

var regexStrPart2 = `do\(\)(.*?)don't\(\)`
var rPart2 = regexp.MustCompile(regexStrPart2)

func getResultPart2(input string) (int, error) {
	input = strings.ReplaceAll(input, "\n", "")

	total := 0
	input = "do()" + input + "don't()"

	matches := rPart2.FindAllStringSubmatch(input, -1)
	for _, match := range matches {
		stringToAnalyze := match[1]
		subtotal, err := getResult(stringToAnalyze)
		if err != nil {
			return 0, err
		}
		total += subtotal
	}
	return total, nil
}
