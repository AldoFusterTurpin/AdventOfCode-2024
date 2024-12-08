package main

import (
	_ "embed"
	"strconv"
	"testing"
)

//go:embed inputs/sample_part2.txt
var samplePart2 string

func Test_getResulPart2(t *testing.T) {
	tests := []struct {
		input   string
		want    int
		wantErr bool
	}{
		{
			input: samplePart2,
			want:  48,
		},
		{
			input: input,
			want:  94455185,
		},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			got, err := getResultPart2(tt.input)
			if err != nil {
				t.Fatalf("expected no error but got %v", err)
			}

			if got != tt.want {
				t.Errorf("getResult() = %v, want %v", got, tt.want)
			}
		})
	}
}
