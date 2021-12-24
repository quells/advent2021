package main

import (
	"bytes"
	"log"
	"strings"
)

type cave struct {
	risks   []int
	w, h    int
	visited []bool
}

func (c cave) scale(n int) (larger cave) {
	larger.w = c.w * 5
	larger.h = c.h * 5
	larger.risks = make([]int, larger.w*larger.h)
	larger.visited = make([]bool, larger.w*larger.h)

	for j := 0; j < n; j++ {
		for i := 0; i < n; i++ {
			offset := i + j - 1
			for y := 0; y < c.h; y++ {
				for x := 0; x < c.w; x++ {
					oldIdx := c.idx(x, y)
					newIdx := larger.idx(x+c.w*i, y+c.h*j)
					newRisk := ((c.risks[oldIdx] + offset) % 9) + 1
					larger.risks[newIdx] = newRisk
				}
			}
		}
	}

	return
}

func (c cave) idx(x, y int) int {
	return x + y*c.w
}

func (c *cave) up(riskSoFar, depth, x, y int) (hn HeapNode) {
	if y == 0 {
		return
	}

	idx := x + (y-1)*c.w
	if c.visited[idx] {
		return
	}
	c.visited[idx] = true
	risk := c.risks[idx]

	hn.riskSoFar = riskSoFar + risk
	hn.x = x
	hn.y = y - 1
	hn.depth = depth + 1
	return
}

func (c *cave) down(riskSoFar, depth, x, y int) (hn HeapNode) {
	if y >= c.h-1 {
		return
	}

	idx := x + (y+1)*c.w
	if c.visited[idx] {
		return
	}
	c.visited[idx] = true
	risk := c.risks[idx]

	hn.riskSoFar = riskSoFar + risk
	hn.x = x
	hn.y = y + 1
	hn.depth = depth + 1
	return
}

func (c *cave) left(riskSoFar, depth, x, y int) (hn HeapNode) {
	if x == 0 {
		return
	}

	idx := x - 1 + y*c.w
	if c.visited[idx] {
		return
	}
	c.visited[idx] = true
	risk := c.risks[idx]

	hn.riskSoFar = riskSoFar + risk
	hn.x = x - 1
	hn.y = y
	hn.depth = depth + 1
	return
}

func (c *cave) right(riskSoFar, depth, x, y int) (hn HeapNode) {
	if x >= c.w-1 {
		return
	}

	idx := x + 1 + y*c.w
	if c.visited[idx] {
		return
	}
	c.visited[idx] = true
	risk := c.risks[idx]

	hn.riskSoFar = riskSoFar + risk
	hn.x = x + 1
	hn.y = y
	hn.depth = depth + 1
	return
}

func parse(input string) (c cave) {
	lines := strings.Split(input, "\n")

	c.w = len(lines[0])
	c.h = len(lines)
	c.risks = make([]int, c.w*c.h)
	c.visited = make([]bool, c.w*c.h)

	idx := 0
	for _, line := range lines {
		for _, char := range line {
			c.risks[idx] = int(char - '0')
			idx++
		}
	}

	return
}

func draw(c cave, h *Heap) string {
	board := bytes.Repeat([]byte{'.'}, c.w*c.h)
	for i, v := range c.visited {
		if v {
			board[i] = '_'
		}
	}

	for _, n := range h.nodes {
		i := c.idx(n.x, n.y)
		board[i] = '#'
	}

	var drawn []byte
	for y := 0; y < c.h; y++ {
		drawn = append(drawn, '\n')
		drawn = append(drawn, board[y*c.w:(y+1)*c.w]...)
	}

	return string(drawn)
}

func safestRoute(c cave) (risk int) {
	c.visited[0] = true

	h := new(Heap)
	x, y := 0, 0
	depth := 0

	rounds := 0
	for {
		if x == c.w-1 && y == c.h-1 {
			break
		}

		{
			hn := c.right(risk, depth, x, y)
			if hn.riskSoFar > 0 {
				h.Insert(hn)
			}
		}
		{
			hn := c.down(risk, depth, x, y)
			if hn.riskSoFar > 0 {
				h.Insert(hn)
			}
		}
		{
			hn := c.left(risk, depth, x, y)
			if hn.riskSoFar > 0 {
				h.Insert(hn)
			}
		}
		{
			hn := c.up(risk, depth, x, y)
			if hn.riskSoFar > 0 {
				h.Insert(hn)
			}
		}

		// log.Println(draw(c, h))

		pos := h.Pop()
		x = pos.x
		y = pos.y
		risk = pos.riskSoFar
		depth = pos.depth
		rounds++
	}
	log.Println(rounds, len(c.risks))

	return
}
