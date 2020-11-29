package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func Merge(nums1, nums2 []int, m, n int) []int {
	copy(nums1[n:], nums1[:m])
	res := nums1
	nums1 = res[n:]
	var j, k int
	for i := range res {
		if j >= m {
			copy(res[i:], nums2[k:])
			break
		}
		if k >= n {
			copy(res[i:], nums1[j:])
			break
		}
		if nums1[j] <= nums2[k] {
			res[i] = nums1[j]
			j++
			continue
		}
		res[i] = nums2[k]
		k++
	}
	return res
}

func main() {
	usage := fmt.Sprintf("Usage: %s <n> <m>\n", os.Args[0])
	if len(os.Args) < 3 {
		fmt.Print(usage)
		return
	}
	m, err := strconv.Atoi(os.Args[1])
	if err != nil || m < 0 {
		fmt.Print(usage)
		return
	}
	n, err := strconv.Atoi(os.Args[2])
	if err != nil || n < 0 {
		fmt.Print(usage)
		return
	}
	nums1 := make([]int, m+n)
	nums2 := make([]int, n)
	slice := nums1
	p := 0
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)
	for i := 0; i < m+n && sc.Scan(); i++ {
		p = i
		if i >= m {
			slice = nums2
			p = i - m
		}
		slice[p], err = strconv.Atoi(sc.Text())
		if err != nil {
			log.Fatal(err)
		}
	}
	fmt.Println(Merge(nums1, nums2, m, n))
}
