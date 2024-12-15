package main

import (
	"fmt"
	"strings"
)

/*
.M.S..
..A..M
.M.S.M
..A.AS
.M.S.M
......

Idea: iterate over all the matrix looking for the word A (as this is where both "MAS" join).
Then, look at the adjecent diagonal cells and if that corresponds to an X-MAS, we increment the counter.
*/
func part2(input string) (int, error) {
	matrix := strings.Split(input, "\n")
	nRows := len(matrix)
	nCols := len(matrix[0])
	fmt.Printf("Matrix has %v rows and %v columns\n", nRows, nCols)

	fmt.Println("\ncontents of matrix:")
	count := 0
	for i := 0; i < nRows; i++ {
		for j := 0; j < nCols; j++ {
			// fmt.Printf("matrix[%v][%v] is %v\n", i, j, string(matrix[i][j]))
			c := string(matrix[i][j])
			if c == "A" {
				if thereIsA_X_MASThere(matrix, i, j) {
					count++
				}
			}
		}
		// fmt.Println("__________________")
	}

	return count, nil
}

func thereIsA_X_MASThere(matrix []string, i, j int) bool {
	nRows := len(matrix)
	nCols := len(matrix[0])
	if i-1 < 0 || j-1 < 0 || i+1 >= nRows || j+1 >= nCols {
		return false
	}

	if string(matrix[i-1][j-1]) == "M" && string(matrix[i-1][j+1]) == "S" && string(matrix[i+1][j-1]) == "M" && string(matrix[i+1][j+1]) == "S" {
		return true
	}

	if string(matrix[i-1][j-1]) == "M" && string(matrix[i-1][j+1]) == "M" && string(matrix[i+1][j-1]) == "S" && string(matrix[i+1][j+1]) == "S" {
		return true
	}

	if string(matrix[i-1][j-1]) == "S" && string(matrix[i-1][j+1]) == "M" && string(matrix[i+1][j-1]) == "S" && string(matrix[i+1][j+1]) == "M" {
		return true
	}

	if string(matrix[i-1][j-1]) == "S" && string(matrix[i-1][j+1]) == "S" && string(matrix[i+1][j-1]) == "M" && string(matrix[i+1][j+1]) == "M" {
		return true
	}

	return false
}
