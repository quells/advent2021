package main

import "strings"

type grid struct {
	points []int
	w, h   int

	blinked []bool
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

func (g *grid) step() (blinked int) {
	g.blinked = make([]bool, len(g.points))

	for y := 0; y < g.h; y++ {
		offset := y * g.w
		for x := 0; x < g.w; x++ {
			idx := offset + x
			next := g.points[idx] + 1
			g.points[idx] = next
			if next > 9 && !g.blinked[idx] {
				g.blinked[idx] = true
				g.ripple(x, y, idx)
			}
		}
	}

	for idx, b := range g.blinked {
		if b {
			blinked++
			g.points[idx] = 0
		}
	}

	return
}

func (g *grid) ripple(x, y, idx int) {
	if x > 0 {
		wIdx := idx - 1
		w := g.points[wIdx] + 1
		g.points[wIdx] = w
		if w > 9 && !g.blinked[wIdx] {
			g.blinked[wIdx] = true
			g.ripple(x-1, y, wIdx)
		}
		if y > 0 {
			nwIdx := idx - g.w - 1
			nw := g.points[nwIdx] + 1
			g.points[nwIdx] = nw
			if nw > 9 && !g.blinked[nwIdx] {
				g.blinked[nwIdx] = true
				g.ripple(x-1, y-1, nwIdx)
			}
		}
		if y < g.h-1 {
			swIdx := idx + g.w - 1
			sw := g.points[swIdx] + 1
			g.points[swIdx] = sw
			if sw > 9 && !g.blinked[swIdx] {
				g.blinked[swIdx] = true
				g.ripple(x-1, y+1, swIdx)
			}
		}
	}
	if x < g.w-1 {
		eIdx := idx + 1
		e := g.points[eIdx] + 1
		g.points[eIdx] = e
		if e > 9 && !g.blinked[eIdx] {
			g.blinked[eIdx] = true
			g.ripple(x+1, y, eIdx)
		}
		if y > 0 {
			neIdx := idx - g.w + 1
			ne := g.points[neIdx] + 1
			g.points[neIdx] = ne
			if ne > 9 && !g.blinked[neIdx] {
				g.blinked[neIdx] = true
				g.ripple(x+1, y-1, neIdx)
			}
		}
		if y < g.h-1 {
			seIdx := idx + g.w + 1
			se := g.points[seIdx] + 1
			g.points[seIdx] = se
			if se > 9 && !g.blinked[seIdx] {
				g.blinked[seIdx] = true
				g.ripple(x+1, y+1, seIdx)
			}
		}
	}
	if y > 0 {
		nIdx := idx - g.w
		n := g.points[nIdx] + 1
		g.points[nIdx] = n
		if n > 9 && !g.blinked[nIdx] {
			g.blinked[nIdx] = true
			g.ripple(x, y-1, nIdx)
		}
	}
	if y < g.h-1 {
		sIdx := idx + g.w
		s := g.points[sIdx] + 1
		g.points[sIdx] = s
		if s > 9 && !g.blinked[sIdx] {
			g.blinked[sIdx] = true
			g.ripple(x, y+1, sIdx)
		}
	}
}

func (g grid) isSynched() bool {
	for _, b := range g.blinked {
		if !b {
			return false
		}
	}
	return true
}
