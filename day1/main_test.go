package main

import (
	_ "embed"
	"strconv"
	"testing"
)

//go:embed inputs/input.txt
var simple_input string

func Test_getResult(t *testing.T) {
	tests := []struct {
		fileContent string
		want        int
	}{
		{
			fileContent: simple_input,
			want:        2264607,
		},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			if got, err := getResult(tt.fileContent); got != tt.want {
				if err != nil {
					t.Errorf("expected no error, but got %v", err)
				}
				t.Errorf("getResult() = %v, want %v", got, tt.want)
			}
		})
	}
}
