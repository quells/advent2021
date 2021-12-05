package main

import (
	"regexp"
	"strconv"
	"strings"
)

func sign(x int) int {
	if x == 0 {
		return 0
	}
	if x < 0 {
		return -1
	}
	return 1
}

type line [2][2]int

func (l line) isOrthogonal() bool {
	if l[0][0] == l[1][0] {
		return true
	}
	if l[0][1] == l[1][1] {
		return true
	}
	return false
}

func (l line) points() (ps [][2]int) {
	dx := sign(l[1][0] - l[0][0])
	dy := sign(l[1][1] - l[0][1])

	if dx == 0 {
		// vertical
		var m, M int
		if l[0][1] < l[1][1] {
			m = l[0][1]
			M = l[1][1]
		} else {
			m = l[1][1]
			M = l[0][1]
		}
		x := l[0][0]
		for y := m; y <= M; y++ {
			ps = append(ps, [2]int{x, y})
		}
		return
	}

	if dy == 0 {
		// horizontal
		var m, M int
		if l[0][0] < l[1][0] {
			m = l[0][0]
			M = l[1][0]
		} else {
			m = l[1][0]
			M = l[0][0]
		}
		y := l[0][1]
		for x := m; x <= M; x++ {
			ps = append(ps, [2]int{x, y})
		}
		return
	}

	// diagonal
	x := l[0][0]
	y := l[0][1]
	for {
		ps = append(ps, [2]int{x, y})
		if x == l[1][0] || y == l[1][1] {
			return
		}
		x += dx
		y += dy
	}
}

var linePattern = regexp.MustCompile(`^(\d+),(\d+) -> (\d+),(\d+)$`)

func parseLine(s string) (coords line) {
	matches := linePattern.FindAllStringSubmatch(s, 1)
	m := matches[0]
	a, _ := strconv.Atoi(m[1])
	b, _ := strconv.Atoi(m[2])
	c, _ := strconv.Atoi(m[3])
	d, _ := strconv.Atoi(m[4])
	return line{{a, b}, {c, d}}
}

func parseInput(s string) (parsed []line) {
	split := strings.Split(s, "\n")
	parsed = make([]line, len(split))
	for i, si := range split {
		parsed[i] = parseLine(si)
	}
	return
}

type ventMap map[[2]int]int // point : count

func (vm *ventMap) draw(l line) {
	for _, p := range l.points() {
		oldCount := (*vm)[p]
		(*vm)[p] = oldCount + 1
	}
}

func (vm ventMap) countOverlaps() (overlaps int) {
	for _, c := range vm {
		if c > 1 {
			overlaps++
		}
	}
	return
}
