package main

import "testing"

func TestComb(t *testing.T) {
	tests := []struct {
		n, r int
		want int
	}{
		{0, 1, 0},
		{0, 3, 0},
		{1, 0, 1},
		{1, 1, 1},
		{5, 5, 1},
		{5, 3, 10},
		{5, 2, 10},
	}

	for _, tt := range tests {
		got, want := Comb(tt.n, tt.r), tt.want
		if got != want {
			t.Errorf("Comb(%d, %d) got %d want %d", tt.n, tt.r, got, want)
		}
	}
}
