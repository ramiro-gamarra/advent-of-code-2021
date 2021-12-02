package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strconv"
)

func main() {
	bs, err := os.ReadFile("testdata/input")
	if err != nil {
		panic(err)
	}

	input := parseInput(bs)
	incs := countThreeSumIncreases(input)
	fmt.Println(incs)
}

func countThreeSumIncreases(in []int) int {
	if len(in) < 4 {
		return 0
	}

	increases := 0
	l := in[0] + in[1] + in[2]

	for i := 3; i < len(in); i++ {
		r := in[i] + in[i-1] + in[i-2]
		if l < r {
			increases++
		}
		l = r
	}

	return increases
}

func parseInput(bs []byte) []int {
	var ss []string

	sc := bufio.NewScanner(bytes.NewBuffer(bs))
	for sc.Scan() {
		ss = append(ss, sc.Text())
	}

	var out []int
	for _, s := range ss {
		n, err := strconv.Atoi(s)
		if err != nil {
			continue
		}
		out = append(out, n)
	}

	return out
}
