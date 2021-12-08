package main

import (
	"math"
	"sort"
	"strconv"
	"strings"
)

func parseInput(s string) (parsed []int) {
	csv := strings.Split(s, ",")

	parsed = make([]int, len(csv))
	for i, si := range csv {
		parsed[i], _ = strconv.Atoi(si)
	}

	sort.Ints(parsed)
	return
}

func median(sorted []int) int {
	n := len(sorted)
	if n%2 == 0 {
		return sorted[n/2-1]
	}
	return sorted[n/2]
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func linearFuel(vs []int, target int) (f int) {
	for _, v := range vs {
		f += abs(v - target)
	}
	return
}

func mean(vs []int) int {
	var s int
	for _, v := range vs {
		s += v
	}
	mu := float64(s) / float64(len(vs))
	return int(math.Round(mu))
}

func sumFuel(vs []int, target int) (f int) {
	for _, v := range vs {
		d := abs(v - target)
		f += d * (d + 1) / 2
	}
	return
}
