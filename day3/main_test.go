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

func Test_getResultGood(t *testing.T) {
	tests := []struct {
		input   string
		want    int
		wantErr bool
	}{
		{
			input: sample,
			want:  161,
		},
		{
			input: input,
			want:  187833789,
		},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			got, err := getResult(tt.input)
			if err != nil {
				t.Fatalf("expected no error but got %v", err)
			}

			if got != tt.want {
				t.Errorf("getResult() = %v, want %v", got, tt.want)
			}
		})
	}
}
