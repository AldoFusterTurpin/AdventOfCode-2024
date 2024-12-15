package main

import (
	"reflect"
	"strconv"
	"testing"
)

func Test_part2(t *testing.T) {
	tests := []struct {
		input string
		want  int
	}{
		{
			input: sample,
			want:  9,
		},
		{
			input: input,
			want:  1969,
		},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			got, err := part2(tt.input)
			if err != nil {
				t.Errorf("expected no error, but got %v", err)
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getResult() = %v, want %v", got, tt.want)
			}
		})
	}
}
