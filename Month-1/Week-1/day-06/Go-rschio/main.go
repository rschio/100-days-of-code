package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

func NextPrime(n int) int {
	if n < 2 {
		return 2
	}
	max := int(math.Sqrt(float64(n)))
	for val := 2; val <= max; val++ {
		if n%val == 0 {
			n++
			max = int(math.Sqrt(float64(n)))
			val = 1
		}
	}
	return n
}

func main() {
	usage := fmt.Sprintf("Usage: %s <number>\n", os.Args[0])
	if len(os.Args) < 2 {
		fmt.Print(usage)
		return
	}
	n, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Print(usage)
		return
	}
	next := NextPrime(n)
	fmt.Println(next)
}
