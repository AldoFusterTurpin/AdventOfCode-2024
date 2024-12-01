package main

import (
	_ "embed"
	"strconv"
	"testing"
)

func Test_part2(t *testing.T) {
	tests := []struct {
		fileContent string
		want        int
	}{
		{
			fileContent: simple_input,
			want:        19457120,
		},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			got, err := part2(tt.fileContent)
			if err != nil {
				t.Errorf("expected no error, but got %v", err)
			}

			if got != tt.want {
				t.Errorf("part2() = %v, want %v", got, tt.want)
			}
		})
	}
}
