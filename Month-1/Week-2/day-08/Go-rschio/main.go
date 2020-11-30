package main

import (
	"fmt"
	"os"
)

func Combinations(s string) []string {
	if len(s) == 0 {
		return []string{}
	}
	letters := [8]string{
		"abc", "def", "ghi", "jkl",
		"mno", "pqrs", "tuv", "wxyz",
	}
	total := 1
	for _, r := range s {
		total *= len(letters[r-'2'])
	}
	combs := make([]string, total)
	div := 1
	for _, r := range s {
		p := r - '2'
		l := len(letters[p])
		div *= l
		reps := total / div
		for i := 0; i < div; i += l {
			for j, letter := range letters[p] {
				start := (i * reps) + (j * reps)
				end := (i * reps) + ((j + 1) * reps)
				appendLetter(combs[start:end], letter)
			}
		}
	}

	return combs
}

func appendLetter(combs []string, l rune) {
	for i := range combs {
		combs[i] += string(l)
	}
}

func main() {
	usage := fmt.Sprintf("Usage: %s <digits>\n", os.Args[0])
	if len(os.Args) < 2 {
		fmt.Print(usage)
		return
	}
	digits := os.Args[1]
	for _, r := range digits {
		if !('2' <= r && r <= '9') {
			fmt.Println("invalid input")
			return
		}
	}
	fmt.Println(Combinations(digits))
}
