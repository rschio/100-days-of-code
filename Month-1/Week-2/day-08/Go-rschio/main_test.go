package main

import (
	"sort"
	"testing"
)

func TestCombinations(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  []string
	}{
		{"first", "23",
			[]string{
				"ad", "ae", "af",
				"bd", "be", "bf",
				"cd", "ce", "cf",
			},
		},
		{"27", "27",
			[]string{
				"ap", "aq", "ar", "as",
				"bp", "bq", "br", "bs",
				"cp", "cq", "cr", "cs",
			},
		},
		{"237", "237",
			[]string{
				"adp", "adq", "adr", "ads", "aep", "aeq",
				"aer", "aes", "afp", "afq", "afr", "afs",
				"bdp", "bdq", "bdr", "bds", "bep", "beq",
				"ber", "bes", "bfp", "bfq", "bfr", "bfs",
				"cdp", "cdq", "cdr", "cds", "cep", "ceq",
				"cer", "ces", "cfp", "cfq", "cfr", "cfs",
			},
		},
		{"second", "", []string{}},
		{"third", "2", []string{"a", "b", "c"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Combinations(tt.input)
			sort.Strings(got)
			sort.Strings(tt.want)
			if !equal(got, tt.want) {
				t.Fatalf("want: %v\ngot: %v", tt.want, got)
			}
		})
	}
}

func equal(a, b []string) bool {
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
