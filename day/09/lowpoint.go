package main

import (
	"strings"
)

type grid struct {
	points []int
	w, h   int
}

func parseInput(s string) (g grid) {
	lines := strings.Split(s, "\n")
	for _, line := range lines {
		g.w = len(line)
		g.h++

		linePoints := make([]int, len(line))
		for i := 0; i < len(line); i++ {
			c := line[i]
			linePoints[i] = int(c - '0')
		}

		g.points = append(g.points, linePoints...)
	}
	return
}

func (g grid) lowPoints() (lps []int) {
	for y := 0; y < g.h; y++ {
		offset := y * g.w
		for x := 0; x < g.w; x++ {
			idx := offset + x
			v := g.points[idx]

			if x > 0 {
				w := g.points[idx-1]
				if w <= v {
					continue
				}
			}
			if x < g.w-1 {
				e := g.points[idx+1]
				if e <= v {
					continue
				}
			}
			if y > 0 {
				n := g.points[idx-g.w]
				if n <= v {
					continue
				}
			}
			if y < g.h-1 {
				s := g.points[idx+g.w]
				if s <= v {
					continue
				}
			}
			lps = append(lps, v)
		}
	}
	return
}

// risk is the sum of each (low point + 1)
// which is equal to the sum of low points + count of low points
func risk(lps []int) (r int) {
	for _, lp := range lps {
		r += lp
	}
	r += len(lps)
	return
}

type topThree struct {
	a, b, c int
}

func (t *topThree) insert(x int) {
	if x >= t.a {
		t.a, t.b, t.c = x, t.a, t.b
	} else if x >= t.b {
		t.b, t.c = x, t.b
	} else if x >= t.c {
		t.c = x
	}
}

func (g *grid) basinSize(x, y int) (size int) {
	if x < 0 || g.w <= x || y < 0 || g.h <= y {
		return 0
	}
	idx := x + y*g.w

	v := g.points[idx]
	if v == 9 {
		return 0
	}

	g.points[idx] = 9
	size++

	size += g.basinSize(x, y-1)
	size += g.basinSize(x+1, y)
	size += g.basinSize(x, y+1)
	size += g.basinSize(x-1, y)
	return
}

func (g *grid) largestBasins() (tt topThree) {
	for y := 0; y < g.h; y++ {
		for x := 0; x < g.w; x++ {
			bs := g.basinSize(x, y)
			tt.insert(bs)
		}
	}
	return
}
