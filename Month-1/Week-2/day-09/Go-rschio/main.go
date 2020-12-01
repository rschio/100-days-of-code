package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func Trapping(in []int) int {
	if len(in) <= 1 {
		return 0
	}
	j := 0
	total := 0
	for i := 0; i < len(in)-1; i = j {
		start := in[i]

		acc := 0
		maxEnd := 0
		for j = i + 1; j < len(in); j++ {
			end := in[j]
			maxEnd = max(end, maxEnd)
			if end >= start {
				total += acc
				break
			}
			acc += start - end
		}

		if maxEnd < start {
			acc = 0
			for j = i + 1; j < len(in); j++ {
				acc += maxEnd - in[j]
				if in[j] == maxEnd {
					total += acc
				}
			}
		}
	}
	return total
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	usage := fmt.Sprintf("Usage: %s <size>\n", os.Args[0])
	if len(os.Args) < 2 {
		fmt.Print(usage)
		return
	}
	n, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Print(usage)
		return
	}
	slice := make([]int, n)
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)
	for i := 0; i < n && sc.Scan(); i++ {
		slice[i], err = strconv.Atoi(sc.Text())
		if err != nil {
			log.Fatal("invalid input")
		}
	}
	fmt.Println(Trapping(slice))
}
