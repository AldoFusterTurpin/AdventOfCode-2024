package main

import (
	_ "embed"
	"strconv"
	"testing"
)

//go:embed inputs/sample.txt
var sample string

//go:embed inputs/input.txt
var input string

func Test_getResult(t *testing.T) {
	tests := []struct {
		fileContent string
		want        int
	}{
		{
			fileContent: sample,
			want:        2,
		},
    {
			fileContent: input,
			want:        246,
		},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			if got, err := getNumberOfSafeReports(tt.fileContent); got != tt.want {
				if err != nil {
					t.Errorf("expected no error, but got %v", err)
				}
				t.Errorf("getResult() = %v, want %v", got, tt.want)
			}
		})
	}
}
