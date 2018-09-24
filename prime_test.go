package main

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestPrimeFactors(t *testing.T) {
	tests := []struct {
		input int
		want  []int
	}{
		{0, nil},
		{1, nil},
		{2, []int{2}},
		{4, []int{2, 2}},
		{60, []int{2, 2, 3, 5}},
		{100000, []int{2, 2, 2, 2, 2, 5, 5, 5, 5, 5}},
	}

	for _, tt := range tests {
		got, want := PrimeFactors(tt.input), tt.want
		if diff := cmp.Diff(got, want); diff != "" {
			t.Errorf("given %d, want %v got %v", tt.input, want, got)
		}
	}
}
