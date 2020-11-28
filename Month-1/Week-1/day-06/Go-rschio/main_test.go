package main

import "testing"

func TestNextPrime(t *testing.T) {
	tests := []struct {
		name string
		num  int
		want int
	}{
		{"12", 12, 13},
		{"24", 24, 29},
		{"11", 11, 11},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NextPrime(tt.num)
			if tt.want != got {
				t.Fatalf("want: %d, got: %d", tt.want, got)
			}
		})
	}
}
