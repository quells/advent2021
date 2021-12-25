package main

import "math"

type target struct {
	x [2]int
	y [2]int
}

func (t target) vxRange() (min, max int) {
	// n*(n+1)/2 = c
	c := float64(t.x[0])
	x := (math.Sqrt(8*c+1) - 1) / 2
	min = int(math.Ceil(x))

	// 1 step
	max = t.x[1]

	return
}

func (t target) vyRange() (min, max int) {
	// n*(n+1)/2
	n := 1 - t.y[0]
	max = n * (n + 1) / 2

	// 1 step
	min = t.y[0]

	return
}

type trajectory struct {
	x, y   int
	vx, vy int
}

func (t *trajectory) step() {
	t.x += t.vx
	t.y += t.vy

	if t.vx > 0 {
		t.vx--
	}
	t.vy--
}

func (t trajectory) inTarget(tar target) bool {
	if t.x < tar.x[0] || tar.x[1] < t.x {
		return false
	}
	if t.y < tar.y[0] || tar.y[1] < t.y {
		return false
	}
	return true
}

func fire(vx, vy int, tar target) (lands bool) {
	t := trajectory{vx: vx, vy: vy}
	for {
		if tar.x[1] < t.x {
			return false
		}
		if t.y < tar.y[0] {
			return false
		}
		if t.inTarget(tar) {
			return true
		}

		t.step()
	}
}

func countValidTrajectories(tar target) (count int) {
	xMin, xMax := tar.vxRange()
	yMin, yMax := tar.vyRange()

	for vx := xMin; vx <= xMax; vx++ {
		for vy := yMin; vy <= yMax; vy++ {
			if fire(vx, vy, tar) {
				count++
			}
		}
	}
	return
}
