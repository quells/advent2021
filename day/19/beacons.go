package main

import (
	_ "embed"
	"fmt"
	"math"
	"strconv"
	"strings"
)

type Reading struct {
	n int
	x []int
	y []int
	z []int

	angles  Angles
	crosses CrossList
}

func (r *Reading) applyRotationMatrix(m [9]int) {
	prev := [3][]int{r.x, r.y, r.z}
	var next [3][]int
	for j := 0; j < 3; j++ {
		for i := 0; i < 3; i++ {
			mi := m[i+3*j]
			if mi == 1 {
				next[j] = prev[i]
			} else if mi == -1 {
				next[j] = make([]int, r.n)
				for idx, p := range prev[i] {
					next[j][idx] = -p
				}
			}
		}
	}
	r.x, r.y, r.z = next[0], next[1], next[2]

	// r.angles = nil
	// r.crosses = nil
	// r.calcCrosses()
	for angle, cross := range r.crosses {
		cross.applyRotationMatrix(m)
		r.crosses[angle] = cross
	}
}

func abs(x int) int {
	if x >= 0 {
		return x
	}
	return -x
}

func (r *Reading) alignTo(o Reading) {
	for angle := range r.Intersect(o) {
		cross := r.crosses[angle]
		cx := abs(cross.c1)
		cy := abs(cross.c2)
		cz := abs(cross.c3)
		if cx == cy || cx == cz || cy == cz {
			continue
		}

		oCross := o.crosses[angle]
		rot := findRotationMatrix(cross, oCross)
		r.applyRotationMatrix(rot)
		cross = r.crosses[angle]

		dx := oCross.x - cross.x
		for i := range r.x {
			r.x[i] += dx
		}
		dy := oCross.y - cross.y
		for i := range r.y {
			r.y[i] += dy
		}
		dz := oCross.z - cross.z
		for i := range r.z {
			r.z[i] += dz
		}

		return
	}

	panic("alignment not found")
}

type position struct {
	x, y, z int
}

func (r *Reading) merge(o Reading) {
	positions := make(map[position]struct{})
	for i := 0; i < r.n; i++ {
		positions[position{r.x[i], r.y[i], r.z[i]}] = struct{}{}
	}
	for i := 0; i < o.n; i++ {
		positions[position{o.x[i], o.y[i], o.z[i]}] = struct{}{}
	}

	r.n = len(positions)
	r.x = make([]int, r.n)
	r.y = make([]int, r.n)
	r.z = make([]int, r.n)
	i := 0
	for p := range positions {
		r.x[i] = p.x
		r.y[i] = p.y
		r.z[i] = p.z
		i++
	}

	r.angles = nil
	r.crosses = nil
	r.calcCrosses()
}

type CrossAngle struct {
	l  float64
	st float64
}

type Angles map[CrossAngle]int

type Cross struct {
	x, y, z    int
	a1, a2, a3 int
	b1, b2, b3 int
	c1, c2, c3 int
}

type CrossList map[CrossAngle]Cross

func findRotationMatrix(from, to Cross) (m [9]int) {
	f := []int{from.c1, from.c2, from.c3}
	t := []int{to.c1, to.c2, to.c3}

	for j := 0; j < 3; j++ {
		for i := 0; i < 3; i++ {
			idx := i + 3*j
			if f[i] == t[j] {
				m[idx] = 1
			} else if f[i] == -t[j] {
				m[idx] = -1
			}
		}
	}

	return
}

func applyRotationMatrix(m [9]int, x, y, z int) (nx, ny, nz int) {
	nx = m[0]*x + m[1]*y + m[2]*z
	ny = m[3]*x + m[4]*y + m[5]*z
	nz = m[6]*x + m[7]*y + m[8]*z
	return
}

func (c *Cross) applyRotationMatrix(m [9]int) {
	c.x, c.y, c.z = applyRotationMatrix(m, c.x, c.y, c.z)
	c.a1, c.a2, c.a3 = applyRotationMatrix(m, c.a1, c.a2, c.a3)
	c.b1, c.b2, c.b3 = applyRotationMatrix(m, c.b1, c.b2, c.b3)
	c.c1, c.c2, c.c3 = applyRotationMatrix(m, c.c1, c.c2, c.c3)
}

func (r Reading) Intersect(o Reading) (u Angles) {
	u = make(Angles)
	for angle, aCount := range r.angles {
		bCount := o.angles[angle]
		if bCount == 0 {
			continue
		}

		if aCount < bCount {
			u[angle] = aCount
		} else {
			u[angle] = bCount
		}
	}
	return
}

func CountAngleMatches(u Angles) (n int) {
	for _, count := range u {
		n += count
	}
	return
}

func (r Reading) Overlaps(o Reading, threshold int) bool {
	// threshold = 60 for 12 beacons
	return CountAngleMatches(r.Intersect(o)) >= threshold
}

func (r *Reading) calcCrosses() {
	r.angles = make(Angles)
	r.crosses = make(CrossList)

	for idx := 0; idx < r.n; idx++ {
		xi := r.x[idx]
		yi := r.y[idx]
		zi := r.z[idx]

		dx := make([]int, r.n)
		dy := make([]int, r.n)
		dz := make([]int, r.n)
		for j := 0; j < r.n; j++ {
			dx[j] = r.x[j] - xi
			dy[j] = r.y[j] - yi
			dz[j] = r.z[j] - zi
		}

		dx = append(dx[:idx], dx[idx+1:]...)
		dy = append(dy[:idx], dy[idx+1:]...)
		dz = append(dz[:idx], dz[idx+1:]...)

		m := make([]float64, r.n-1)
		for i := 0; i < r.n-1; i++ {
			a1 := dx[i]
			a2 := dy[i]
			a3 := dz[i]
			m2 := a1*a1 + a2*a2 + a3*a3
			m[i] = math.Sqrt(float64(m2))
		}

		for i := 0; i < r.n-2; i++ {
			a1 := dx[i]
			a2 := dy[i]
			a3 := dz[i]
			a := m[i]
			if a > 1500 {
				continue
			}
			for j := i + 1; j < r.n-1; j++ {
				b1 := dx[j]
				b2 := dy[j]
				b3 := dz[j]
				b := m[j]
				if b > 1500 {
					continue
				}

				cx := a2*b3 - a3*b2
				cy := a3*b1 - a1*b3
				cz := a1*b2 - a2*b1
				c2 := cx*cx + cy*cy + cz*cz
				c := math.Sqrt(float64(c2))

				angle := CrossAngle{
					l:  c,
					st: c / (a * b),
				}
				r.angles[angle]++

				cross := Cross{
					xi, yi, zi,
					a1, a2, a3,
					b1, b2, b3,
					cx, cy, cz,
				}
				r.crosses[angle] = cross
			}
		}
	}
}

func Parse(input string) (readings []Reading) {
	var r Reading
	for _, line := range strings.Split(input, "\n") {
		if line == "" {
			r.calcCrosses()
			readings = append(readings, r)
			r = Reading{}
			continue
		}

		if line[1] == '-' {
			// 0v234...
			// --- scanner x ---
			// -123,456,-789
			continue
		}

		r.n++
		coords := strings.Split(line, ",")
		x, _ := strconv.Atoi(coords[0])
		y, _ := strconv.Atoi(coords[1])
		z, _ := strconv.Atoi(coords[2])
		r.x = append(r.x, x)
		r.y = append(r.y, y)
		r.z = append(r.z, z)
	}

	if r.n > 0 {
		r.calcCrosses()
		readings = append(readings, r)
	}

	return
}

func AlignAndCombine(readings []Reading) Reading {
	zero := readings[0]
	rest := readings[1:]

outer:
	for {
		if len(rest) == 0 {
			break
		}

		for i, other := range rest {
			if zero.Overlaps(other, 60) {
				other.alignTo(zero)
				zero.merge(other)

				rest = append(rest[:i], rest[i+1:]...)

				fmt.Println(zero.n, len(zero.angles))
				continue outer
			}
		}

		panic("no matches found")
	}

	return zero
}

//go:embed 19puzzle.txt
var puzzleInput string

func main() {
	readings := Parse(puzzleInput)
	combined := AlignAndCombine(readings)
	fmt.Println(combined.n)
}
