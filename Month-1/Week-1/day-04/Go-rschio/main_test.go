package main

import "testing"

func TestProgressDays(t *testing.T) {
	tests := []struct {
		name  string
		slice []int
		want  int
	}{
		{"two", []int{3, 4, 1, 2}, 2},
		{"three", []int{10, 11, 12, 9, 10}, 3},
		{"one", []int{6, 5, 4, 3, 2, 9}, 1},
		{"zero", []int{9, 9}, 0},
		{"not enough", []int{9}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ProgressDays(tt.slice)
			if got != tt.want {
				t.Fatalf("got: %d, want: %d", got, tt.want)
			}
		})
	}
}
