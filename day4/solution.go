package main

import (
	"bytes"
	"strings"
)

func getResult(input string) (int, error) {
	matrix := strings.Split(input, "\n")
	nRows := len(matrix)
	/* 	nCols := len(matrix[0])
	   	fmt.Printf("Matrix has %v rows and %v columns\n", nRows, nCols)

	   	fmt.Println("\ncontents of matrix:")
	   	for i := 0; i < nRows; i++ {
	   		for j := 0; j < nCols; j++ {
	   			fmt.Printf("matrix[%v][%v] is %v\n", i, j, string(matrix[i][j]))
	   		}
	   		fmt.Println("__________________")
	   	} */

	rowsReversed, err := getRowsReversed(input)
	if err != nil {
		return 0, err
	}

	convertedRowsToColumns, err := convertRowsToColumns(input)
	if err != nil {
		return 0, err
	}

	total := 0
	targetString := "XMAS"

	// rows and cols
	for i := 0; i < nRows; i++ {
		row := matrix[i]
		total += strings.Count(row, targetString)

		column := convertedRowsToColumns[i]
		total += strings.Count(column, targetString)

		columnReversed, err := reverseString(column)
		if err != nil {
			return 0, err
		}
		total += strings.Count(columnReversed, targetString)

		rowReversed := rowsReversed[i]
		total += strings.Count(rowReversed, targetString)
	}

	// diagonals
	originalDiagonals, err := getAllDiagonalsFromString(input)
	if err != nil {
		return 0, err
	}
	diagonalsMirrored, err := getAllDiagonalsFromSliceOfRows(rowsReversed)
	if err != nil {
		return 0, err
	}

	for i := 0; i < len(originalDiagonals); i++ {
		originalDiagonal := originalDiagonals[i]
		total += strings.Count(originalDiagonal, targetString)

		diagonalReversed, err := reverseString(originalDiagonal)
		if err != nil {
			return 0, err
		}
		total += strings.Count(diagonalReversed, targetString)

		diagonalMirrored := diagonalsMirrored[i]
		total += strings.Count(diagonalMirrored, targetString)

		reversedDiagonalMirrored, err := reverseString(diagonalMirrored)
		if err != nil {
			return 0, err
		}
		total += strings.Count(reversedDiagonalMirrored, targetString)
	}

	return total, nil
}

// getAllDiagonalsFromString traverses the input string returning all the diagonals of the matrix
// by modifying i and j indexes accordingly. Check diagonals.jpg for the explanation (I have
// included the main diagonal in what is called "right diagonals" in "diagonals.jpg").
func getAllDiagonalsFromString(matrixStr string) ([]string, error) {
	matrixStr = strings.TrimSpace(matrixStr)
	matrix := strings.Split(matrixStr, "\n")

	// perform main and right diagonals
	diagonals := getRightDiagonals(matrix)

	// perform left diagonals
	leftDiagonals := getLeftDiagonals(matrix)
	diagonals = append(diagonals, leftDiagonals...)
	return diagonals, nil
}

// getAllDiagonalsFromSliceOfRows is the same as getAllDiagonalsFromString but
// expecting as an input a slice of strings (Matrix splitted)
// instead of the original Matrix in string format, for convenience.
func getAllDiagonalsFromSliceOfRows(matrix []string) ([]string, error) {
	// perform main and right diagonals
	diagonals := getRightDiagonals(matrix)

	// perform left diagonals
	leftDiagonals := getLeftDiagonals(matrix)
	diagonals = append(diagonals, leftDiagonals...)
	return diagonals, nil
}

func getLeftDiagonals(matrix []string) []string {
	n := len(matrix)
	var diagonals []string
	for currentRow := 1; currentRow < n; currentRow++ {
		diagonalBuffer := bytes.Buffer{}
		i := currentRow
		j := 0

		for j < n-currentRow {
			element := matrix[i][j]
			// elementStr := string(matrix[i][j])
			// fmt.Println(elementStr)

			diagonalBuffer.WriteByte(element)

			i++
			j++
		}
		diagonal := diagonalBuffer.String()
		diagonals = append(diagonals, diagonal)
	}
	return diagonals
}

func getRightDiagonals(matrix []string) []string {
	n := len(matrix)
	var diagonals []string
	for currentColumn := 0; currentColumn < n; currentColumn++ {
		diagonalBuffer := bytes.Buffer{}
		i := 0
		j := currentColumn

		for i < n-currentColumn {
			element := matrix[i][j]
			// elementStr := string(matrix[i][j])
			// fmt.Println(elementStr)

			diagonalBuffer.WriteByte(element)

			i++
			j++
		}
		diagonal := diagonalBuffer.String()
		diagonals = append(diagonals, diagonal)
	}
	return diagonals
}

/* func print(s []string) {
	for i := 0; i < len(s); i++ {
		fmt.Println(s[i])
	}
} */

// getRowsReversed returns a copy of input but reversed vertically.
func getRowsReversed(input string) ([]string, error) {
	input = strings.TrimSpace(input)
	matrix := strings.Split(input, "\n")
	nRows := len(matrix)

	reversedVettically := make([]string, nRows)
	for i := 0; i < nRows; i++ {
		reversedString, err := reverseString(matrix[i])
		if err != nil {
			return nil, err
		}
		reversedVettically[i] = reversedString
	}
	return reversedVettically, nil
}

// convertRowsToColumns returns a copy of input but reversed horizontally.
func convertRowsToColumns(input string) ([]string, error) {
	var result []string
	matrix := strings.Split(input, "\n")
	nRows := len(matrix)
	nCols := len(matrix[0])
	for i := 0; i < nRows; i++ {
		currentColumn := bytes.Buffer{}
		for j := 0; j < nCols; j++ {
			element := matrix[j][i]
			err := currentColumn.WriteByte(element)
			if err != nil {
				return nil, err
			}
		}
		result = append(result, currentColumn.String())
	}
	return result, nil
}

func reverseString(s string) (string, error) {
	w := bytes.Buffer{}
	n := len(s)
	for i := n - 1; i >= 0; i-- {
		err := w.WriteByte(s[i])
		if err != nil {
			return "", err
		}
	}
	return w.String(), nil
}

// unused
// type conditionCheckFn func(i, j int) bool

// UNUSDED but fun to play with while I was solving this...
// getDiagonal returns the diagonal of input starting at pos [startI, startJ].
// conditionCheck is used to decide if the function should return the element at pos [i,j]
// for all i..j. Depending in which diagonal we are interested, the function will be different.
/* func getDiagonal(input string, conditionCheck conditionCheckFn, startI, startJ int) (string, error) {
	result := bytes.Buffer{}
	matrix := strings.Split(input, "\n")
	nRows := len(matrix)
	nCols := len(matrix[0])
	for i := startI; i < nRows; i++ {
		for j := startJ; j < nCols; j++ {
			if conditionCheck(i, j) {
				element := matrix[i][j]
				// s := string(element)
				// fmt.Println(element)
				err := result.WriteByte(element)
				if err != nil {
					return "", err
				}
			}
		}
	}
	return result.String(), nil
} */

/*
To play with
func getDiagonals_Old() {
	f := func(i, j int) bool { return i == j }
	d, err := getDiagonal(input, f, 0, 0)
	if err != nil {
		return 0, err
	}
	fmt.Println(d)

	f = func(i, j int) bool { return i == j-1 }
	d, err = getDiagonal(input, f, 0, 1)
	if err != nil {
		return 0, err
	}
	fmt.Println(d)

	f = func(i, j int) bool { return i == j-2 }
	d, err = getDiagonal(input, f, 0, 2)
	if err != nil {
		return 0, err
	}
	fmt.Println(d)

	f = func(i, j int) bool { return i == j+1 }
	d, err = getDiagonal(input, f, 1, 0)
	if err != nil {
		return 0, err
	}
	fmt.Println(d)

	f = func(i, j int) bool { return i == j+2 }
	d, err = getDiagonal(input, f, 2, 0)
	if err != nil {
		return 0, err
	}
	fmt.Println(d)
}
*/
