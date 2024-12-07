package main

import (
	"reflect"
	"strconv"
	"testing"
)

func Test_getCopyWithouthElementAtIndex(t *testing.T) {
	tests := []struct {
		original       []int
		indexToExclude int
		want           []int
	}{
		{
			original:       []int{1, 2, 3, 4},
			indexToExclude: 0,
			want:           []int{2, 3, 4},
		},
		{
			original:       []int{6, 9, 0, 7},
			indexToExclude: 1,
			want:           []int{6, 0, 7},
		},
		{

			original:       []int{6, 9, 0, 7},
			indexToExclude: 2,
			want:           []int{6, 9, 7},
		},
		{
			original:       []int{1, 2, 3, 4},
			indexToExclude: 3,
			want:           []int{1, 2, 3},
		},
		{
			original:       []int{1, 2, 3, 4},
			indexToExclude: 999,
			want:           []int{1, 2, 3, 4},
		},
		{
			original:       []int{1, 2, 3, 4},
			indexToExclude: -1,
			want:           []int{1, 2, 3, 4},
		},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			if got := getCopyWithouthElementAtIndex(tt.original, tt.indexToExclude); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getCopyWithouthElementAtIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getAllPermutationsOfReport(t *testing.T) {
	tests := []struct {
		original []int
		want     [][]int
	}{
		{
			original: []int{1, 2, 3, 4},
			want: [][]int{
				{2, 3, 4},
				{1, 3, 4},
				{1, 2, 4},
				{1, 2, 3},
			},
		},
		{
			original: []int{1, 2},
			want: [][]int{
				{2},
				{1},
			},
		},
		{
			original: []int{4, 10, 20},
			want: [][]int{
				{10, 20},
				{4, 20},
				{4, 10},
			},
		},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			if got := getAllPermutationsOfReport(tt.original); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getAllPermutationsOfReport() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getNumberOfSafeReportsPart2(t *testing.T) {
	tests := []struct {
		input string
		want  int
	}{
		{
			input: sample,
			want:  4,
		},
		{
			input: input,
			want:  318,
		},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			got, err := getNumberOfSafeReportsPart2(tt.input)
			if err != nil {
				t.Fatalf("expected no error, but got %v", err)
			}

			if got != tt.want {
				t.Errorf("getNumberOfSafeReportsPart2() = %v, want %v", got, tt.want)
			}
		})
	}
}
