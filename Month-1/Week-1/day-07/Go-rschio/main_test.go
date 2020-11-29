package main

import "testing"

func TestMerge(t *testing.T) {
	tests := []struct {
		name  string
		nums1 []int
		nums2 []int
		m, n  int
		want  []int
	}{
		{"first",
			[]int{1, 2, 3, 0, 0, 0},
			[]int{2, 5, 6},
			3, 3,
			[]int{1, 2, 2, 3, 5, 6},
		},
		{"second",
			[]int{2, 5, 6, 0, 0, 0},
			[]int{1, 2, 3},
			3, 3,
			[]int{1, 2, 2, 3, 5, 6},
		},
		{"zero",
			[]int{},
			[]int{},
			0, 0,
			[]int{},
		},
		{"zero1",
			[]int{0, 0, 0},
			[]int{1, 2, 3},
			0, 3,
			[]int{1, 2, 3},
		},
		{"zero2",
			[]int{2, 5, 6},
			[]int{},
			3, 0,
			[]int{2, 5, 6},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Merge(tt.nums1, tt.nums2, tt.m, tt.n)
			if !equal(tt.want, got) {
				t.Fatalf("want: %v\ngot: %v", tt.want, got)
			}
		})
	}
}

func equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
