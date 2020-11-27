package main

import "testing"

func TestSockPairs(t *testing.T) {
	tests := []struct {
		name  string
		pairs string
		want  int
	}{
		{"one", "AA", 1},
		{"two", "ABABC", 2},
		{"four", "CABBACCC", 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := SockPairs(tt.pairs)
			if got != tt.want {
				t.Fatalf("got: %d, want: %d", got, tt.want)
			}
		})
	}
}
