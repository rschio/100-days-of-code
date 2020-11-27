package main

import "testing"

func TestBarbecueSkewers(t *testing.T) {
	tests := []struct {
		name    string
		skewers []string
		want    [2]int
	}{
		{"1veg", []string{
			"--xo--x--ox--",
			"--xx--x--xx--",
			"--oo--o--oo--",
			"--xx--x--ox--",
			"--xx--x--ox--",
		},
			[2]int{1, 4},
		},
		{"2veg", []string{
			"--oooo-ooo--",
			"--xx--x--xx--",
			"--o---o--oo--",
			"--xx--x--ox--",
			"--xx--x--ox--",
		},
			[2]int{2, 3},
		},
		{"3veg", []string{
			"--oooo-ooo--",
			"--xxxxxxxx--",
			"--o---",
			"-o-----o---x--",
			"--o---o-----",
		},
			[2]int{3, 2},
		},
		{"0veg", []string{
			"--ooxoo-ooo--",
			"--xxxxxxxx--",
			"--ox----",
			"-o---x--o---x--",
			"--o--x-o-----",
		},
			[2]int{0, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			veg, nveg := BarbecueSkewers(tt.skewers)
			if tt.want[0] != veg || tt.want[1] != nveg {
				t.Fatalf("got: [%d, %d], want: %v", veg, nveg, tt.want)
			}
		})
	}
}
