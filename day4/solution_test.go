package main

import (
	_ "embed"
	"reflect"
	"strconv"
	"testing"
)

//go:embed inputs/sample.txt
var sample string

//go:embed inputs/input.txt
var input string

//go:embed inputs/hola.txt
var hola string

//go:embed inputs/hola_2.txt
var hola_2 string

//go:embed inputs/hola_3.txt
var hola_3 string

//go:embed inputs/generic.txt
var generic string

// func Test_getResult(t *testing.T) {
// 	tests := []struct {
// 		fileContent string
// 		want        int
// 	}{
// 		{
// 			fileContent: sample,
// 			want:        2,
// 		},
// 		/* {
// 			fileContent: input,
// 			want:        0,
// 		}, */
// 	}
// 	for i, tt := range tests {
// 		t.Run(strconv.Itoa(i), func(t *testing.T) {
// 			if got, err := getResult(tt.fileContent); got != tt.want {
// 				if err != nil {
// 					t.Errorf("expected no error, but got %v", err)
// 				}
// 				t.Errorf("getResult() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }

func Test_reverseString(t *testing.T) {
	tests := []struct {
		s    string
		want string
	}{
		{
			s:    "hola",
			want: "aloh",
		},
		{
			s:    "",
			want: "",
		},
		{
			s:    "reverse",
			want: "esrever",
		},
		{
			s:    "MMMSXXMASM",
			want: "MSAMXXSMMM",
		},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			got, err := reverseString(tt.s)
			if err != nil {
				t.Errorf("expected no error, but got %v", err)
			}
			if got != tt.want {
				t.Errorf("reverseString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getReversedVertically(t *testing.T) {
	tests := []struct {
		input string
		want  []string
	}{
		{
			input: hola,
			want: []string{
				"...h",
				"...o",
				"...l",
				"...a",
			},
		},
		{
			input: sample,
			want: []string{
				"MSAMXXSMMM",
				"ASMSMXMASM",
				"MMAAMXSXMA",
				"XMSMSAMASM",
				"MMAXMASAMX",
				"AMAXXMMAXX",
				"SSXSASMSMS",
				"AAASAMAXAS",
				"MMMMXMMMAM",
				"XSAMXAXMXM",
			},
		},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			got, err := getReversedVertically(tt.input)
			if err != nil {
				t.Errorf("expected no error, but got %v", err)
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getReversedVertically() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getReversedHorizontally(t *testing.T) {
	tests := []struct {
		input string
		want  []string
	}{
		{
			input: hola_2,
			want: []string{
				"hola",
				"hola",
				"....",
				"....",
			},
		},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			got, err := getReversedHorizontally(tt.input)
			if err != nil {
				t.Errorf("expected no error, but got %v", err)
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getReversedHorizontally() = %v, want %v", got, tt.want)
			}
		})
	}
}

/* func Test_getDiagonalAtPoint(t *testing.T) {
	tests := []struct {
		input  string
		startI int
		startJ int
		want   string
	}{
		{
			input: hola_3,
			want:  "hola",
		},
		{
			input: `....
....
..L.
...A`,
			startI: 2,
			startJ: 2,
			want:   "LA",
		},
		// 		{
		// 			input: `....
		// ....
		// R...
		// .O..`,
		// 			startI: 2,
		// 			startJ: 0,
		// 			want:   "RO",
		// 		},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			got, err := getDiagonalAtPoint(tt.input, tt.startI, tt.startJ)
			if err != nil {
				t.Errorf("expected no error, but got %v", err)
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getDiagonalAtPoint() = %v, want %v", got, tt.want)
			}
		})
	}
} */

func Test_getAllDiagonals(t *testing.T) {
	tests := []struct {
		input string
		want  []string
	}{
		{
			input: generic,
			want:  []string{"1111", "222", "33", "0", "444", "55", "0"},
		},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			got, err := getAllDiagonals(tt.input)
			if err != nil {
				t.Errorf("expected no error, but got %v", err)
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getAllDiagonals() = %v, want %v", got, tt.want)
			}
		})
	}
}
