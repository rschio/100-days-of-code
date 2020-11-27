package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func BarbecueSkewers(skewers []string) (veg, nveg int) {
	for _, s := range skewers {
		pos := strings.Index(s, "x")
		if pos == -1 {
			veg++
			continue
		}
		nveg++
	}
	return
}

func main() {
	in, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatalf("failed to read input: %v", err)
	}
	p := 0
	if p = bytes.Index(in, []byte("[")); p == -1 {
		log.Fatal("invalid input")
	}
	in = in[p+1:]
	if p = bytes.Index(in, []byte("]")); p == -1 {
		log.Fatal("invalid input")
	}
	in = in[:p]
	in = bytes.Replace(in, []byte(`"`), []byte(""), -1)
	skewers := strings.Split(string(in), ",")
	veg, nveg := BarbecueSkewers(skewers)
	fmt.Printf("[%d, %d]\n", veg, nveg)
}
