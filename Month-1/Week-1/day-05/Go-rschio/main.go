package main

import (
	"fmt"
	"os"
)

func SockPairs(s string) int {
	runes := make(map[rune]int)
	for _, r := range s {
		runes[r]++
	}
	total := 0
	for _, val := range runes {
		total += val / 2
	}
	return total
}

func main() {
	usage := fmt.Sprintf("Usage: %s <sockpairs>\n", os.Args[0])
	if len(os.Args) < 2 {
		fmt.Print(usage)
		return
	}
	pairs := SockPairs(os.Args[1])
	fmt.Println(pairs)
}
