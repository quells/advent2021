package main

import "strings"

// A node on the path.
// Even values can be visited multiple times.
type node int

func (n node) canVisitMultipleTimes() bool {
	return n%2 == 0
}

const (
	start node = 1
	end   node = -1
)

type paths map[node][]node // node : connections

func (p *paths) insert(from, to node) {
	(*p)[from] = append((*p)[from], to)
	(*p)[to] = append((*p)[to], from)
}

func parsePaths(input string) paths {
	labels := map[string]node{
		"start": start,
		"end":   end,
	}
	nextLargeIdx := 2
	nextSmallIdx := 3
	readLabel := func(s string) node {
		n := labels[s]
		if n != 0 {
			return n
		}

		if s[0] < 'a' {
			// uppercase are large
			n = node(nextLargeIdx)
			nextLargeIdx += 2
		} else {
			n = node(nextSmallIdx)
			nextSmallIdx += 2
		}

		labels[s] = n
		return n
	}

	ps := make(paths)
	for _, line := range strings.Split(input, "\n") {
		ends := strings.Split(line, "-")
		l := readLabel(ends[0])
		r := readLabel(ends[1])
		ps.insert(l, r)
	}
	return ps
}

type toVisitList [][]node

func (tvl *toVisitList) pop() []node {
	n := len(*tvl) - 1
	tail := (*tvl)[n]
	*tvl = (*tvl)[:n]
	return tail
}

func (tvl *toVisitList) push(ns []node) {
	(*tvl) = append(*tvl, ns)
}

type visitedList []node

func (vl *visitedList) pop() node {
	n := len(*vl) - 1
	tail := (*vl)[n]
	*vl = (*vl)[:n]
	return tail
}

func (vl *visitedList) push(n node) {
	(*vl) = append(*vl, n)
}

func (vl visitedList) canVisit(n node) bool {
	for _, v := range vl {
		if v == n {
			return v.canVisitMultipleTimes()
		}
	}
	return true
}

func countPaths(ps paths) (count int) {
	var visited visitedList
	visited.push(start)

	var toVisit toVisitList
	toVisit.push(ps[start])

	for {
		if len(toVisit) == 0 {
			break
		}

		options := toVisit.pop()
		if len(options) == 0 {
			_ = visited.pop()
			continue
		}

		next := options[0]
		remaining := options[1:]
		toVisit.push(remaining)

		if next == end {
			count++
			continue
		}

		if visited.canVisit(next) {
			toVisit.push(ps[next])
			visited.push(next)
		}
	}
	return
}
