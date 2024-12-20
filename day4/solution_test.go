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

//go:embed inputs/generic.txt
var generic string

//go:embed inputs/generic_2.txt
var generic2 string

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

func Test_getRowsReversed(t *testing.T) {
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
			got, err := getRowsReversed(tt.input)
			if err != nil {
				t.Errorf("expected no error, but got %v", err)
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getReversedVertically() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_convertRowsToColumns(t *testing.T) {
	tests := []struct {
		input string
		want  []string
	}{
		{
			input: hola_2,
			want: []string{
				"hola",
				"hola",
				"AAAA",
				"BBBB",
			},
		},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			got, err := convertRowsToColumns(tt.input)
			if err != nil {
				t.Errorf("expected no error, but got %v", err)
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getReversedHorizontally() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getAllDiagonals(t *testing.T) {
	tests := []struct {
		input string
		want  []string
	}{
		{
			input: generic,
			want:  []string{"1111", "222", "33", "0", "444", "55", "0"},
		},
		{
			input: generic2,
			want:  []string{"11111", "2222", "333", "00", "9", "4444", "555", "00", "9"},
		},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			got, err := getAllDiagonalsFromString(tt.input)
			if err != nil {
				t.Errorf("expected no error, but got %v", err)
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getAllDiagonals() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getResult(t *testing.T) {
	tests := []struct {
		input string
		want  int
	}{
		{
			input: sample,
			want:  18,
		},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			got, err := getResult(tt.input)
			if err != nil {
				t.Errorf("expected no error, but got %v", err)
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getResult() = %v, want %v", got, tt.want)
			}
		})
	}
}
