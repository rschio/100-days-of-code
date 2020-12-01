package main

import "testing"

func TestTrapping(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		want  int
	}{
		{"first", []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}, 6},
		{"second", []int{4, 2, 0, 3, 2, 5}, 9},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Trapping(tt.input)
			if got != tt.want {
				t.Fatalf("got: %d, want: %d", got, tt.want)
			}
		})
	}
}
