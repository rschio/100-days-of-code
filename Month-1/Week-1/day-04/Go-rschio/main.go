package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func ProgressDays(results []int) int {
	if len(results) <= 1 {
		return 0
	}
	count := 0
	prev := results[0]
	for _, now := range results {
		if now > prev {
			count++
		}
		prev = now
	}
	return count
}

func main() {
	usage := fmt.Sprintf("Usage: %s <arr_size>\n", os.Args[0])
	if len(os.Args) < 2 {
		fmt.Print(usage)
		return
	}
	size, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Print(usage)
		return
	}
	slice := make([]int, size)
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)
	for i := 0; i < size && sc.Scan(); i++ {
		num := sc.Text()
		slice[i], err = strconv.Atoi(num)
		if err != nil {
			log.Fatalf("failed to read input: %v", err)
		}
	}
	res := ProgressDays(slice)
	fmt.Println(res)
}
