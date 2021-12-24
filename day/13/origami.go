package main

import (
	"bytes"
	"strconv"
	"strings"
)

type fold struct {
	x int
	y int
}

func parse(input string) (points [][2]int, folds []fold) {
	lines := strings.Split(input, "\n")
	var seenBlank bool
	for _, line := range lines {
		if line == "" {
			seenBlank = true
			continue
		}

		if !seenBlank {
			// positions
			pos := strings.Split(line, ",")
			x, _ := strconv.Atoi(pos[0])
			y, _ := strconv.Atoi(pos[1])
			points = append(points, [2]int{x, y})
		} else {
			// fold instructions
			var f fold
			fa := strings.Split(line[11:], "=")
			if fa[0] == "x" {
				f.x, _ = strconv.Atoi(fa[1])
			} else {
				f.y, _ = strconv.Atoi(fa[1])
			}
			folds = append(folds, f)
		}
	}

	return
}

func apply(points [][2]int, f fold) [][2]int {
	applied := make([][2]int, len(points))
	if f.y != 0 {
		// fold up
		for i, p := range points {
			if p[1] > f.y {
				p[1] = 2*f.y - p[1]
			}
			applied[i] = p
		}
	} else {
		// fold left
		for i, p := range points {
			if p[0] > f.x {
				p[0] = 2*f.x - p[0]
			}
			applied[i] = p
		}
	}
	return applied
}

func countVisible(points [][2]int) int {
	visible := make(map[int]struct{})
	for _, p := range points {
		h := (p[1] << 16) | p[0]
		visible[h] = struct{}{}
	}
	return len(visible)
}

func draw(points [][2]int) string {
	var w, h int
	for _, p := range points {
		if p[0] >= w {
			w = p[0] + 1
		}
		if p[1] >= h {
			h = p[1] + 1
		}
	}
	if w > 100 || h > 100 {
		return "too big"
	}

	rows := make([][]byte, h)
	{
		row := bytes.Repeat([]byte{'.'}, w)
		for i := range rows {
			rows[i] = make([]byte, w)
			copy(rows[i], row)
		}
	}

	for _, p := range points {
		rows[p[1]][p[0]] = '#'
	}

	var drawn []byte
	for _, row := range rows {
		drawn = append(drawn, '\n')
		drawn = append(drawn, row...)
	}

	return string(drawn)
}
